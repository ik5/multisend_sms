package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	multisendsms "github.com/ik5/multisend_sms"
)

var (
	userName  = os.Getenv("MULTISEND_USERNAME")
	password  = os.Getenv("MULTISEND_PASSWORD")
	from      = os.Getenv("MULTISEND_FROM")
	recipient = os.Getenv("MULTISEND_RECIPIENT")
	messageID = os.Getenv("MULTISEND_MESSAGE_ID")
)

func main() {
	client := http.Client{
		Timeout: time.Second * 15,
	}

	rest := multisendsms.RESTSendSMS{
		UserName:          userName,
		Password:          password,
		From:              from,
		Recipient:         recipient,
		Message:           "Hello world",
		CustomerMessageID: messageID,
	}

	resp, err := rest.SendSMS(http.MethodGet, &client, nil)
	fmt.Println(resp, err)
}
