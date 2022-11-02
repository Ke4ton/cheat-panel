package admin

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"onefluid/internal/database"
	"onefluid/internal/models"
	"onefluid/pgk/cookiesession"
	"onefluid/pgk/generator"
	"strconv"
)

func CreateBucket(c *fiber.Ctx) error {
	bucketType, _ := strconv.Atoi(c.FormValue("type"))

	bucket := models.PromoBucket{
		BucketName:   c.FormValue("name"),
		BucketAuthor: cookiesession.GetUser(c).Username,
		Type:         bucketType,
	}

	database.DB.Create(&bucket)

	return c.Redirect("/admin/promo")
}

func GenerateRandomPromo(c *fiber.Ctx) error {
	bucketId, _ := strconv.Atoi(c.FormValue("bucket"))
	amount, _ := strconv.Atoi(c.FormValue("amount"))
	activations, _ := strconv.Atoi(c.FormValue("activations"))
	hours, _ := strconv.Atoi(c.FormValue("hours"))

	for i := 0; i < amount; i++ {
		database.DB.Create(&models.PromoCode{
			BucketID:       uint(bucketId),
			Code:           generator.RandStringRunes(12),
			Creator:        cookiesession.GetUser(c).Username,
			Activations:    0,
			MaxActivations: activations,
			Time:           hours,
		})
	}

	return c.Redirect("/admin/promo/bucket/" + strconv.Itoa(bucketId))
}

func GeneratePromo(c *fiber.Ctx) error {
	bucketId, _ := strconv.Atoi(c.FormValue("bucket"))
	activations, _ := strconv.Atoi(c.FormValue("activations"))
	hours, _ := strconv.Atoi(c.FormValue("hours"))

	database.DB.Create(&models.PromoCode{
		BucketID:       uint(bucketId),
		Code:           c.FormValue("code"),
		Creator:        cookiesession.GetUser(c).Username,
		Activations:    0,
		MaxActivations: activations,
		Time:           hours,
	})

	return c.Redirect("/admin/promo/bucket/" + strconv.Itoa(bucketId))
}

func ExportBucket(c *fiber.Ctx) error {
	var buckets models.PromoBucket
	database.DB.Find(&buckets, "id = ?", c.Params("id"))

	var keys []models.PromoCode
	database.DB.Find(&keys, "`bucket_id` = ?", buckets.ID)

	var str string
	for _, k := range keys {
		if k.MaxActivations != k.Activations {
			str += fmt.Sprintf("Hours: %d | Promo: %s | Link: https://cheat.company/activate/%s\n", k.Time, k.Code, k.Code)
		}
	}

	return c.SendString(str)
}

func DeleteKey(c *fiber.Ctx) error {
	database.DB.Delete(&models.PromoCode{}, "`code` = ?", c.Params("code"))
	return c.Redirect("/admin/promo/bucket/" + c.Params("bucket"))
}
