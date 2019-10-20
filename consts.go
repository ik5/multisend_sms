package multisendsms

// Default settings
const (
	DefaultHTTPAddress    = "https://api.multisend.co.il/MultiSendAPI/"
	DefaultSendSMSPage    = "sendsms"
	DefaultDateTimeFormat = "2006-02-01+15:04:05"
)

var dlrDescEng = map[DLRType]string{
	DLRTypeNumberWithoutDevice: "The number is not linked to a device",
	DLRTypeDeviceMemoryFull1:   "Device memory is full",
	DLRTypeFilteredMessage:     "Receiver filtered message",
	DLRTypeDeviceNotSupported:  "Device does not support SMS",
	DLRTypeServiceNotSupported: "Service does not support SMS",
	DLRTypeDeviceMemoryFull2:   "Device is full",
	DLRTypeMessageExpired:      "Message expired",
	DLRTypeSuccess:             "Delivered",
}
