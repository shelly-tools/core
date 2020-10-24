package models

// Building such as home or garage
type Building struct {
	IDBuilding      int    `storm:"id,increment"`
	BuildingName    string `storm:"index"`
	BuildingOrder   int
	BuildingPicture string
}
