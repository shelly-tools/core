package models

// Building such as home or garage
type Building struct {
	IDBuilding      int    `storm:"id,increment" json:"idBuilding"` //The ID of the building
	BuildingName    string `storm:"index" json:"buildingName"`      //the Name of the building
	BuildingOrder   int    `json:"buildingOrder"`                   // the order number of the building, 1 is first in the list..
	BuildingPicture string `json:"buildingPicture"`                 // a presentational picture for the building
}
