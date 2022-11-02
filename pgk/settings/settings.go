package settings

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type projectSettings struct {
	DSN        string `json:"dsn"`
	Debug      bool   `json:"debug"`
	HCaptcha   string `json:"chaptcha"`
	Telegram   string `json:"telegram"`
	TelegramID string `json:"telegram_id"`
	Technical  bool   `json:"technical"`
}

var Config = projectSettings{
	DSN:        "log:pass@/dbname?parseTime=true",
	Debug:      false,
	HCaptcha:   "hcaptcha_token",
	Telegram:   "telegram_bot_token",
	TelegramID: "telegram_chat_id",
	Technical:  false,
}

func Parse() {
	if configExists("config.json") == true {
		jsonFile, _ := os.Open("config.json")
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &Config)
		return
	} else {
		file, _ := os.Create("config.json")
		defer file.Close()
		json_example, _ := json.Marshal(Config)
		file.WriteString(string(json_example))
	}
}

func Export() {
	file, _ := os.Create("config.json")
	defer file.Close()
	json_example, _ := json.Marshal(Config)
	file.WriteString(string(json_example))
}

func configExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
