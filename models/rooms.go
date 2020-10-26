package models

// Room inside a building
type Room struct {
	IDRoom          int    `storm:"id,unique,increment"` // the uniqe ID of the room
	RoomName        string `storm:"index"`               //the Rooms name
	IDBuilding      int    // the ID of the building the room was placed in
	RoomOrder       int    // the order number, 1 is shown first in the list of rooms
	RoomPicture     string // a picture from the room
	IDDefaultSensor int    // the id of a temperature (and humidity) sensor, temperature wiil be shown in room index
}
