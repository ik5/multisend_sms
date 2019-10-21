package multisendsms

import (
	"net/url"
	"reflect"
	"testing"
	"time"
)

const (
	userName    = "1"
	password    = "2"
	from        = "me"
	recipient   = "9721234567"
	message     = "Hello world"
	messageType = MessageTypeSMS
)

var (
	currentTime      = time.Now()
	scheduleDateTime = NewSchedulerFromTime(currentTime)
)

func TestValidToURL(t *testing.T) {
	r := RESTSendSMS{
		UserName:         userName,
		Password:         password,
		From:             from,
		Recipient:        recipient,
		Message:          message,
		MessageType:      messageType,
		ScheduleDateTime: *scheduleDateTime,
		International:    false,
	}
	expected := url.Values{
		"user":          []string{userName},
		"password":      []string{password},
		"from":          []string{from},
		"recipient":     []string{recipient},
		"message":       []string{message},
		"message_type":  []string{messageType.String()},
		"international": []string{"0"},
	}

	gen := r.ToURL()

	// reflect.DeepEqual does not properly check if the maps are the same,
	// and always fails, so a loop is used instead
	for key, value := range gen {
		value2, found := expected[key]
		if !found {
			t.Errorf("%s was not found at the expected values", key)
		}

		if !reflect.DeepEqual(value, value2) {
			t.Errorf("%s is '%+v' on gen, but '%+v' on expected", key, value, value2)
		}
	}
}
