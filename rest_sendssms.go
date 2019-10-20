package multisendsms

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// RESTSendSMS holds the fields to send SMS using GET and POST requests
type RESTSendSMS struct {
	UserName                   string            `url:"user"`
	Password                   string            `url:"password"`
	From                       string            `url:"from"`
	Recipient                  string            `url:"recipient"`
	Message                    string            `url:"message"`
	MessageType                MessageType       `url:"message_type,omitempty"`
	ScheduleDateTime           SchedulerDateTime `url:"scheduledatetime,omitempty"`
	International              Bool              `url:"international,omitempty"`
	DeliveryNotificationURL    string            `url:"deliverynotification_url,omitempty"`
	CustomerMessageID          string            `url:"customermessageid,omitempty"`
	DeliveryNotificationMethod string            `url:"deliverynotificationmethod,omitempty"`
	SendID                     string            `url:"sendID,omitempty"`
}

// ToURL converts the struct to url.Values
func (r *RESTSendSMS) ToURL() url.Values {
	var result url.Values
	val := reflect.ValueOf(r).Elem()
	fieldsCount := val.NumField()

	for i := 0; i < fieldsCount; i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		name := tag.Get("url")
		empty := ""
		if name == "" || name == "-" {
			continue
		}
		elems := strings.Split(name, ",")
		switch len(elems) {
		case 0:
			continue
		case 2:
			empty = elems[1]
			fallthrough
		case 1:
			name = elems[0]
		}

		omitEmpty := empty == "omitempty"
		if name == "" || name == "-" {
			continue
		}
		if omitEmpty {
			continue
		}
		switch valueField.Interface().(type) {
		case string:
			val := valueField.String()
			if val == "" {
				continue
			}
			result.Add(typeField.Name, val)
		case Bool:
			val := Bool(valueField.Bool())
			result.Add(typeField.Name, strconv.Itoa(val.Int()))
		case bool:
			val := valueField.Bool()
			if val {
				result.Add(typeField.Name, "1")
				continue
			}
			result.Add(typeField.Name, "0")
		case MessageType:
			val := valueField.Interface().(MessageType)
			result.Add(typeField.Name, val.String())
		case SchedulerDateTime:
			val := valueField.Interface().(SchedulerDateTime)
			if !val.IsValid() {
				continue
			}
			result.Add(typeField.Name, val.String())
		}
	}

	return result
}
