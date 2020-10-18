package main

// Room inside a building
type Room struct {
	IDRoom          int    `storm:"id,unique"`
	RoomName        string `storm:"index"`
	IDBuilding      int
	RoomOrder       int
	RoomPicture     string
	IDDefaultSensor int
}

// Rooms is a Collection of Rooms
type Rooms struct {
	Room []Room
}
