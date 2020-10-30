package models

// DeviceGroup is used to group devices in order to control them simultanously
type Device struct {
	IDDevice   int    `storm:"id,unique,increment" json:"idDevice"` // primary key for a Device group
	DeviceName string `storm:"index" json:"deviceName"`             // Name of the Group
	DeviceIP   string
}
