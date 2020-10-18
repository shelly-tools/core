package main

import "time"

type Device struct {
	IDDevice   string `storm:"id,unique"`
	DeviceType string `storm:"index"` //SHSW-1, SHSW-PM, SHDM-1 ...
	DeviceName string
	IP         string `storm:"unique"`
	IDRoom     int    //ID Room the device is located in
	User       string //device user and password in order to access the REST-API
	Password   string

	CreatedAt time.Time `storm:"index"`
}

type Devices struct {
	Device []Device
}
