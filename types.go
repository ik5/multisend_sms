package multisendsms

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/ik5/gostrutils"
)

// MessageType is the type of message to tell the API to use
type MessageType uint

// DLRType holds information regarding the DLR
type DLRType string

// SchedulerDateTime holds date and time format to schedule sms
type SchedulerDateTime struct {
	dateTime *time.Time
	valid    bool
}

// Bool is a boolean that understand also numeric values
type Bool bool

// Recipients holds a list of recipients to send them SMS
type Recipients []string

// Type Of messages to use
const (
	MessageTypeUnknown MessageType = iota
	MessageTypeTTS
	MessageTypeSMS
	MessageTypeSMSAndTTS
)

// The Types of given DLR
const (
	DLRTypeSuccess             DLRType = "888"
	DLRTypeNumberWithoutDevice DLRType = "000"
	DLRTypeDeviceMemoryFull1   DLRType = "003"
	DLRTypeFilteredMessage     DLRType = "009"
	DLRTypeDeviceNotSupported  DLRType = "012"
	DLRTypeServiceNotSupported DLRType = "021"
	DLRTypeDeviceMemoryFull2   DLRType = "032"
	DLRTypeMessageExpired      DLRType = "041"
)

// NewSchedulerFromTime initialize a new SchedulerDateTime from
// time.Time
func NewSchedulerFromTime(t time.Time) *SchedulerDateTime {
	s := SchedulerDateTime{
		valid:    true,
		dateTime: &t,
	}
	if t.Equal(time.Time{}) {
		s.valid = false
	}

	return &s
}

// NewSchedulerDateTime initialize a new SchedulerDateTime.
func NewSchedulerDateTime(str string) (*SchedulerDateTime, error) {
	if str == "" {
		return &SchedulerDateTime{nil, false}, nil
	}
	dateTime, err := time.Parse(DefaultDateTimeFormat, str)
	if err != nil {
		return nil, err
	}

	return &SchedulerDateTime{&dateTime, true}, nil
}

// Add a new number for the recipients
func (r *Recipients) Add(number string) {
	*r = append(*r, number)
}

// RemoveByIndex removes a number from Recipients based on an index
// if index is too small or too big, an error is raised
func (r *Recipients) RemoveByIndex(idx int) error {
	if idx < 0 {
		return fmt.Errorf("%d is out of bound", idx)
	}

	if idx >= len(*r) {
		return fmt.Errorf("%d is out of bound", idx)
	}

	*r = append((*r)[:idx], (*r)[idx+1:]...)
	return nil
}

// IsNumberExists searches for a number if exists on the Recipients list
// if found returns true and the index
func (r Recipients) IsNumberExists(number string) (bool, int) {
	index := gostrutils.GetStringIndexInSlice(r, number)
	return index > -1, index
}

func (mt MessageType) String() string {
	switch mt {
	case MessageTypeTTS:
		return "tts"
	case MessageTypeSMS:
		return "sms"
	case MessageTypeSMSAndTTS:
		return "sms+tts"
	}

	return "unknown"
}

func (dt SchedulerDateTime) String() string {
	if !dt.valid || dt.dateTime == nil {
		return ""
	}
	return dt.dateTime.Format(DefaultDateTimeFormat)
}

func (r Recipients) String() string {
	return strings.Join(r, ",")
}

// Int returns int from Bool
func (b Bool) Int() int {
	if b {
		return 1
	}
	return 0
}

// Uint returns int from Bool
func (b Bool) Uint() uint {
	if b {
		return 1
	}
	return 0
}

// Error returns an error from DLR
func (dlr DLRType) Error() string {
	return dlrDescEng[dlr]
}

// SetDateTime update SchedulerDateTime to a new date time
func (dt *SchedulerDateTime) SetDateTime(dateTime time.Time) {
	if dateTime.Equal(time.Time{}) {
		dt.valid = false
		dt.dateTime = nil
		return
	}

	dt.valid = true
	dt.dateTime = &dateTime
}

// IsValid returns true if the date is valid
func (dt SchedulerDateTime) IsValid() bool {
	return dt.valid
}

// Value implements the database interface of Value
func (mt MessageType) Value() (driver.Value, error) {
	return uint(mt), nil
}

// Value implements the database interface of Value
func (dlr DLRType) Value() (driver.Value, error) {
	return string(dlr), nil
}

// Value implements the database interface of Value
func (dt SchedulerDateTime) Value() (driver.Value, error) {
	if !dt.valid || dt.dateTime == nil {
		return nil, nil
	}

	return dt.dateTime.Format(DefaultDateTimeFormat), nil
}

// Value implements the database interface of Value
func (b Bool) Value() (driver.Value, error) {
	return bool(b), nil
}

// Value implements the database interface of Value
func (r Recipients) Value() (driver.Value, error) {
	return r.String(), nil
}

// Scan implements the database interface for Scan
func (mt *MessageType) Scan(src interface{}) error {
	if src == nil {
		return errors.New("src cannot be nil")
	}

	switch src.(type) {
	case string:
		str := reflect.ValueOf(src).String()
		i, err := strconv.ParseUint(str, 10, 16)
		if err != nil {
			return err
		}
		*mt = MessageType(i)
	case uint, uint8, uint16, uint32, uint64:
		val := reflect.ValueOf(src).Uint()
		*mt = MessageType(val)
	case int, int8, int16, int32, int64:
		val := reflect.ValueOf(src).Int()
		if val < 0 {
			return fmt.Errorf("Negative values %d not allowed", val)
		}
		*mt = MessageType(val)
	case float32, float64:
		*mt = MessageType(uint(reflect.ValueOf(src).Float()))
	default:
		return fmt.Errorf("Invalid type of src: %T", src)
	}

	return nil
}

// Scan implements the database interface for Scan
func (dlr *DLRType) Scan(src interface{}) error {
	if src == nil {
		return errors.New("src cannot be nil")
	}

	switch src.(type) {
	case string:
		str := reflect.ValueOf(src).String()
		*dlr = DLRType(str)
	case uint, uint8, uint16, uint32, uint64:
		val := reflect.ValueOf(src).Uint()
		*dlr = DLRType(fmt.Sprintf("%0.3d", val))
	case int, int8, int16, int32, int64:
		val := reflect.ValueOf(src).Int()
		if val < 0 {
			return fmt.Errorf("Negative values %d not allowed", val)
		}
		*dlr = DLRType(fmt.Sprintf("%0.3d", val))
	case float32, float64:
		val := uint(reflect.ValueOf(src).Float())
		*dlr = DLRType(fmt.Sprintf("%0.3d", val))
	default:
		return fmt.Errorf("Invalid type of src: %T", src)
	}
	return nil
}

// Scan implements the database interface for Scan
func (dt *SchedulerDateTime) Scan(src interface{}) error {
	if src == nil {
		dt.dateTime, dt.valid = nil, false
		return nil
	}
	var err error
	dt.valid = true
	switch src.(type) {
	case *time.Time:
		dt.dateTime = src.(*time.Time)
		return nil
	case *string:
		val := src.(*string)
		*dt.dateTime, err = time.Parse(DefaultDateTimeFormat, *val)
		if err != nil {
			dt.dateTime = nil
			dt.valid = false
			return err
		}

	case *[]byte:
		if src == nil {
			dt.valid = false
			dt.dateTime = nil
			return nil
		}

		val := src.(*[]byte)
		*dt.dateTime, err = time.Parse(DefaultDateTimeFormat, string(*val))
		if err != nil {
			dt.dateTime = nil
			dt.valid = false
			return err
		}
		return nil
	}
	return nil
}

// Scan implements the database interface for Scan
func (b *Bool) Scan(src interface{}) error {
	if src == nil {
		return errors.New("src cannot be nil")
	}

	switch src.(type) {
	case string:
		str := reflect.ValueOf(src).String()
		boolean, err := strconv.ParseBool(str)
		if err != nil {
			return err
		}
		*b = Bool(boolean)
	case int, int8, int16, int32, int64:
		*b = Bool(reflect.ValueOf(src).Int() > 0)
	case float32, float64:
		*b = Bool(int(reflect.ValueOf(src).Float()) > 0)
	case bool:
		*b = Bool(reflect.ValueOf(src).Bool())
	default:
		return fmt.Errorf("Invalid type of src: %T", src)
	}

	return nil
}

// Scan implements the database interface for Scan
func (r *Recipients) Scan(src interface{}) error {
	if src == nil {
		return errors.New("src cannot be nil")
	}

	switch src.(type) {
	case string:
		val := reflect.ValueOf(src).String()
		*r = Recipients(strings.Split(val, ","))
	case []byte:
		val := string(reflect.ValueOf(src).Bytes())
		*r = Recipients(strings.Split(val, ","))
	default:
		return fmt.Errorf("Invalid type of src: %T", src)
	}

	return nil
}
