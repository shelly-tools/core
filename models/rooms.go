package models

// Room inside a building
type Room struct {
	IDRoom          int    `storm:"id,unique,increment"`
	RoomName        string `storm:"index"`
	IDBuilding      int
	RoomOrder       int
	RoomPicture     string
	IDDefaultSensor int
}
