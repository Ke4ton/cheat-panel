package hcaptcha

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"onefluid/pgk/settings"
)

func CheckHCaptcha(response string) bool {
	req, _ := http.NewRequest("POST", "https://hcaptcha.com/siteverify", nil)
	q := req.URL.Query()

	q.Add("secret", settings.Config.HCaptcha)
	q.Add("response", response)

	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	var hCaptchaResponse map[string]interface{}

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &hCaptchaResponse)

	return hCaptchaResponse["success"].(bool)
}
