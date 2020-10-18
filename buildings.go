package main

// Building such as home or garage
type Building struct {
	IDBuilding      int    `storm:"id,unique"`
	BuildingName    string `storm:"index"`
	BuildingOrder   int
	BuildingPicture string
}

// Buildings is a collection of buildings
type Buildings struct {
	Building []Building
}
