package handlers

import (
	"encoding/json"
	"er-api-consumer/models"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getRequestUrl() string {
	ER_API_CONSUMER_URL := os.Getenv("ER_API_CONSUMER_URL")
	ER_API_CONSUMER_KEY := os.Getenv("ER_API_CONSUMER_KEY")
	ER_API_CONSUMER_CNC := os.Getenv("ER_API_CONSUMER_CNC")

	if ER_API_CONSUMER_KEY == "" {
		log.Fatal("ER_API_CONSUMER_KEY is not set")
	}
	return fmt.Sprintf("%s/%s/latest/%s", ER_API_CONSUMER_URL, ER_API_CONSUMER_KEY, ER_API_CONSUMER_CNC)
}

func GetData() (models.Response, error) {
	resp, err := http.Get(getRequestUrl())
	if err != nil {
		fmt.Println("No response from request")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var result models.Response
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Cannot unmarshal JSON")
	}

	if result.Result == "success" {
		return result, nil
	} else {
		return result, err
	}
}
