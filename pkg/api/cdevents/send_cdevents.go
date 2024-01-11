package cdevents

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	sdk "github.com/cdevents/sdk-go/pkg/api"
	"gopkg.in/yaml.v3"
)

type Config struct {
	MessageBrokerURL string `yaml:"message_broker"`
}

func SendCDEvent(cdEvent sdk.CDEventReader) {

	req, err := createNewCDEventPostRequest(cdEvent)
	if err != nil {
		fmt.Println("Error creating new HTTP Post request with CDEvent :", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("POST request response status: %s\n", resp.Status)
}

func createNewCDEventPostRequest(cdEvent sdk.CDEventReader) (*http.Request, error) {
	var req *http.Request
	jsonBody, errJson := json.Marshal(cdEvent)
	if errJson != nil {
		fmt.Println("Error marshaling cdEvent to JSON:", errJson)
		return req, errJson
	}

	config, errCfg := loadConfig("config.yaml")
	if errCfg != nil {
		fmt.Println("Error loading configuration:", errCfg)
		return req, errCfg
	}
	fmt.Printf("Sending CDEvent %s to configured MessageBroker URL %s\n", jsonBody, config.MessageBrokerURL)

	// Create a new POST request with the MessageBrokerURL
	req, errReq := http.NewRequest("POST", config.MessageBrokerURL, bytes.NewBuffer([]byte(jsonBody)))
	if errReq != nil {
		fmt.Println("Error creating request:", errReq)
		return req, errReq
	}
	event, err := sdk.AsCloudEvent(cdEvent)

	// Set CloudEvent headers for the request
	req.Header.Set("Ce-Id", event.ID())
	req.Header.Set("Ce-Specversion", event.SpecVersion())
	req.Header.Set("Ce-Type", event.Type())
	req.Header.Set("Ce-Source", event.Source())
	req.Header.Set("Content-Type", "application/json")

	return req, err

}
func loadConfig(fileName string) (Config, error) {
	var config Config

	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error Reading configuration:", err)
		return config, err
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("Error Unmarshal configuration:", err)
		return config, err
	}

	return config, nil
}
