package slackhandler 

import (
	"net/http"
	"log"
)

func PostSlack(jsonStr string) {
	//jsonStr := "{}"
	url := "https://maker.ifttt.com/trigger/huawei_alert/with/key/c9GxSBX5gGyKITjQTGsuwH"
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		logger.Errorf("Invalid http request")
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Errorf("Fail to send notification")
	}
	defer resp.Body.Close()
}
