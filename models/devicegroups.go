package models

// DeviceGroup is used to group devices in order to control them simultanously
type DeviceGroup struct {
	IDDeviceGroup   int    `storm:"id,unique,increment"` // primary key for a Device group
	DeviceGroupName string `storm:"index"`               // Name of the Group
}
