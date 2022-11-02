package language

import "github.com/gofiber/fiber/v2"

var Phrases = map[string]string{
	// Index Page
	"ru.home": "Главная",
	"en.home": "Home",
	"cn.home": "主页",

	"ru.profile": "Профиль",
	"en.profile": "Profile",
	"cn.profile": "个人资料",

	"ru.signup": "Создать аккаунт",
	"en.signup": "Sign Up",
	"cn.signup": "注册",

	"ru.signin": "Войти",
	"en.signin": "Sign In",
	"cn.signin": "登录",

	"ru.description": "Вы получите доступ к лоадеру с Nemesis на один месяц.",
	"en.description": "You will get access to the loader with Nemesis for one month.",
	"cn.description": "您将获得同时使用Nemesis的机会",

	"ru.description3": "Вы получите доступ к лоадеру с Nemesis на три месяца.",
	"en.description3": "You will get access to the loader with Nemesis for three months.",
	"cn.description3": "您将获得同时使用Nemesis的机会",

	"ru.buy": "Купить",
	"en.buy": "Buy",
	"cn.buy": "购买",

	"ru.month": "1 месяц",
	"en.month": "1 month",
	"cn.month": "1个月",

	"ru.3month": "3 месяца",
	"en.3month": "3 months",
	"cn.3month": "3个月",

	"ru.weloveyou": "Мы вас любим",
	"en.weloveyou": "We love you",
	"cn.weloveyou": "We love you",

	"ru.language": "Язык",
	"en.language": "Language",
	"cn.language": "语言",

	// Sign Pages
	"ru.authorization": "Авторизация",
	"en.authorization": "Login to your account",
	"cn.authorization": "登录到您的账户",

	"ru.registration": "Создание аккаунта",
	"en.registration": "Create new account",
	"cn.registration": "创建新账户",

	"ru.username": "Имя пользователя",
	"en.username": "Username",
	"cn.username": "帐号",

	"ru.password": "Пароль",
	"en.password": "Password",
	"cn.password": "密码",

	"ru.dont_have_account": "Еще нет аккаунта?",
	"en.dont_have_account": "Don't have account yet?",
	"cn.dont_have_account": "还没有账户吗？",

	"ru.have_account": "Уже есть аккаунт?",
	"en.have_account": "Already have account?",
	"cn.have_account": "已经有了账户？",

	"ru.user_doesnt_exist": "Пользователь с таким никнеймом не найден.",
	"en.user_doesnt_exist": "No user with this nickname was found.",
	"cn.user_doesnt_exist": "没有发现有此昵称的用户。",

	"ru.wrong_fields": "Проверьте заполненость всех полей.",
	"en.wrong_fields": "Check that all the fields have been filled in.",
	"cn.wrong_fields": "检查所有字段是否已经填写。",

	"ru.wrong_password": "Указан неверный пароль, попробуйте снова.",
	"en.wrong_password": "Please check the password you have entered.",
	"cn.wrong_password": "请检查您所输入的密码。",

	"ru.user_exists": "Пользователь с таким никнеймом уже существует.",
	"en.user_exists": "A user with this nickname or email already exists.",
	"cn.user_exists": "已经有一个拥有此昵称或电子邮件的用户存在。",

	"ru.wrong_email": "Проверьте правильность указанного email.",
	"en.wrong_email": "Please check your Email for correctness.",
	"cn.wrong_email": "请检查您的电子邮件是否正确。",

	"ru.wrong_captcha": "Пройдите тест на робота еще раз.",
	"en.wrong_captcha": "Pass the robot test again.",
	"cn.wrong_captcha": "再次通过机器人测试。",

	// Profile
	"ru.activate": "Активировать",
	"en.activate": "Activate",
	"cn.activate": "激活",

	"ru.promo": "Промокод",
	"en.promo": "Promocode",
	"cn.promo": "激活秘钥兑换",

	"ru.wrong_code": "Проверьте правильность промокода.",
	"en.wrong_code": "Please check the promo code provided.",
	"cn.wrong_code": "请检查提供的促销代码。",

	"ru.code_already_activated": "Промокод уже активирован.",
	"en.code_already_activated": "The promo code has already been activated.",
	"cn.code_already_activated": "该优惠代码已经被激活。",

	"ru.done": "Промокод успешно активирован.",
	"en.done": "The promo code has been successfully activated.",
	"cn.done": "优惠代码已成功激活。",

	"ru.hwidattached": "Привязан",
	"en.hwidattached": "Attached",
	"cn.hwidattached": "附带。",

	"ru.hwidunattached": "Не привязан",
	"en.hwidunattached": "Unattached",
	"cn.hwidunattached": "暂未绑定",

	"ru.downloadloader": "Скачать лоадер",
	"en.downloadloader": "Download Loader",
	"cn.downloadloader": "下载加载器。",

	"ru.subenddate": "Дата окончания подписки",
	"en.subenddate": "End date of subscription",
	"cn.subenddate": "订阅的结束日期。",

	"ru.subactive": "С подпиской",
	"en.subactive": "With subscription",
	"cn.subactive": "有订阅",

	"ru.subinactive": "Без подписки",
	"en.subinactive": "Without subscription",
	"cn.subinactive": "没有订阅",
}

func GetPrefix(c *fiber.Ctx) string {
	cookie := c.Cookies("lang", "ru")
	if cookie == "en" {
		return "en."
	}
	if cookie == "cn" {
		return "cn."
	}

	return "ru."
}
