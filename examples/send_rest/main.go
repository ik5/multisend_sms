package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	multisendsms "github.com/ik5/multisend_sms"
)

var (
	method    = os.Getenv("MULTISEND_METHOD")
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
		Recipient:         multisendsms.Recipients{recipient},
		Message:           time.Now().String(),
		CustomerMessageID: messageID,
	}

	resp, err := rest.SendSMS(method, &client, nil)
	fmt.Println(resp, err)
}
