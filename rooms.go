package main

import "net/http"

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

func RoomList(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "RoomList", "")
}

func RoomAdd(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "RoomAdd", "")
}
func RoomDelete(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "RoomDelete", "")
}
func RoomEdit(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "RoomEdit", "")
}
