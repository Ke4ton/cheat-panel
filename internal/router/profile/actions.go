package profile

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gofiber/fiber/v2"
	"onefluid/internal/database"
	"onefluid/internal/models"
	"onefluid/pgk/cookiesession"
	"onefluid/pgk/hcaptcha"
	"onefluid/pgk/logger"
	"onefluid/pgk/memcache"
	"onefluid/pgk/validation"
	"time"
)

func SetLang(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{Name: "lang", Value: c.Query("lang"), Expires: time.Now().Add(999 * time.Hour)})
	return c.Redirect("/")
}

func SignUp(c *fiber.Ctx) error {
	if hcaptcha.CheckHCaptcha(c.FormValue("h-captcha-response")) == false {
		return c.Redirect("/auth/sign-up?err=wrong_captcha")
	}

	login, password := c.FormValue("login"), c.FormValue("password")
	if len(login) < 3 || len(password) < 5 || len(login) > 32 || len(password) > 32 {
		return c.Redirect("/auth/sign-up?err=wrong_fields")
	}

	email := c.FormValue("email")
	if validation.IsValidEmail(email) == false || len(email) < 4 || len(email) > 48 {
		return c.Redirect("/auth/sign-up?err=wrong_email")
	}

	var userExist int64
	database.DB.Model(&models.UserModel{}).Where("`username` = ? or `email` = ?", login, email).Count(&userExist)
	if userExist != 0 {
		return c.Redirect("/auth/sign-up?err=user_exists")
	}

	passwordHash := md5.Sum([]byte(password))
	hashString := hex.EncodeToString(passwordHash[:])

	var newUser = &models.UserModel{
		Username:     login,
		Password:     hashString,
		Email:        email,
		Subscription: time.Now(),
		Referrer:     c.Cookies("ref"),
	}
	database.DB.Create(&newUser)

	c.Cookie(&fiber.Cookie{Name: "USR", Value: login, Expires: time.Now().Add(48 * time.Hour)})
	c.Cookie(&fiber.Cookie{Name: "Authorized", Value: "true", Expires: time.Now().Add(48 * time.Hour)})

	go memcache.UserCaching()
	logger.Database(login, "Created account", logger.Registrations, c)

	return c.Redirect("/profile")
}

func SignIn(c *fiber.Ctx) error {
	login, password := c.FormValue("login"), c.FormValue("password")
	if len(login) < 3 || len(password) < 5 || len(login) > 32 || len(password) > 32 {
		return c.Redirect("/auth/sign-in?err=wrong_fields")
	}

	var userExist int64
	database.DB.Model(&models.UserModel{}).Where("`username` = ?", login).Count(&userExist)
	if userExist != 1 {
		return c.Redirect("/auth/sign-in?err=user_doesnt_exist")
	}

	var userFound models.UserModel
	database.DB.Find(&userFound, "`username` = ?", login)

	passwordHash := md5.Sum([]byte(password))
	hashString := hex.EncodeToString(passwordHash[:])

	if hashString != userFound.Password {
		return c.Redirect("/auth/sign-in?err=wrong_password")
	}

	c.Cookie(&fiber.Cookie{Name: "USR", Value: login, Expires: time.Now().Add(48 * time.Hour)})
	c.Cookie(&fiber.Cookie{Name: "Authorized", Value: "true", Expires: time.Now().Add(48 * time.Hour)})

	go memcache.UserCaching()
	logger.Database(login, "Loggined to account", logger.Authorizations, c)

	return c.Redirect("/profile")
}

func LogOut(c *fiber.Ctx) error {
	c.ClearCookie()
	return c.Redirect("/")
}

func ActivateKey(c *fiber.Ctx) error {
	var promoCode models.PromoCode
	database.DB.Find(&promoCode, "`code` = ?", c.Params("key"))

	if promoCode.Code == "" {
		return c.Redirect("/profile?err=wrong_code")
	}

	var bucket models.PromoBucket
	database.DB.Find(&bucket, "id = ?", promoCode.BucketID)

	if promoCode.Activations == promoCode.MaxActivations {
		return c.Redirect("/profile?err=code_already_activated")
	}

	currUser := cookiesession.GetUser(c)

	if currUser.Subscription.Unix() > time.Now().Unix() {
		return c.Redirect("/profile?err=already_with_subscription")
	}

	var activates int64
	database.DB.Find(&models.PromoActivations{}, "`code` = ? and `activator` = ?", c.Params("key"), currUser.Username).Count(&activates)
	if activates > 0 {
		return c.Redirect("/profile?err=code_already_activated")
	}

	if currUser.Subscription.Unix() > time.Now().Unix() {
		database.DB.Model(&currUser).Updates(models.UserModel{
			Subscription: currUser.Subscription.Add(time.Duration(promoCode.Time) * time.Hour),
		})
	} else {
		database.DB.Model(&currUser).Updates(models.UserModel{
			Subscription: time.Now().Add(time.Duration(promoCode.Time) * time.Hour),
		})
	}

	database.DB.Model(&promoCode).Updates(map[string]interface{}{
		"Activations": promoCode.Activations + 1,
	})

	if bucket.Type == 1 {
		database.DB.Model(&promoCode).Updates(map[string]interface{}{
			"Time": promoCode.Time / 2,
		})
	}

	database.DB.Create(&models.PromoActivations{
		CodeID:    promoCode.ID,
		Activator: currUser.Username,
	})

	go memcache.UserCaching()
	logger.Database(currUser.Username, "Activated key: "+c.Params("key"), logger.Other, c)

	return c.Redirect("/profile?err=done")
}

func ActivateKeyForm(c *fiber.Ctx) error {
	var promoCode models.PromoCode
	database.DB.Find(&promoCode, "`code` = ?", c.FormValue("key"))

	if promoCode.Code == "" {
		return c.Redirect("/profile?err=wrong_code")
	}

	var bucket models.PromoBucket
	database.DB.Find(&bucket, "id = ?", promoCode.BucketID)

	if promoCode.Activations == promoCode.MaxActivations {
		return c.Redirect("/profile?err=code_already_activated")
	}

	currUser := cookiesession.GetUser(c)

	if currUser.Subscription.Unix() > time.Now().Unix() {
		return c.Redirect("/profile?err=already_with_subscription")
	}

	var activates int64
	database.DB.Find(&models.PromoActivations{}, "`code` = ? and `activator` = ?", c.Params("key"), currUser.Username).Count(&activates)
	if activates > 0 {
		return c.Redirect("/profile?err=code_already_activated")
	}

	if currUser.Subscription.Unix() > time.Now().Unix() {
		database.DB.Model(&currUser).Updates(models.UserModel{
			Subscription: currUser.Subscription.Add(time.Duration(promoCode.Time) * time.Hour),
		})
	} else {
		database.DB.Model(&currUser).Updates(models.UserModel{
			Subscription: time.Now().Add(time.Duration(promoCode.Time) * time.Hour),
		})
	}

	database.DB.Model(&promoCode).Updates(map[string]interface{}{
		"Activations": promoCode.Activations + 1,
	})

	if bucket.Type == 1 {
		database.DB.Model(&promoCode).Updates(map[string]interface{}{
			"Time": promoCode.Time / 2,
		})
	}

	database.DB.Create(&models.PromoActivations{
		CodeID:    promoCode.ID,
		Activator: currUser.Username,
	})

	go memcache.UserCaching()
	logger.Database(currUser.Username, "Activated key: "+c.Params("key"), logger.Other, c)

	return c.Redirect("/profile?err=done")
}
