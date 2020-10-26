package models

// Building such as home or garage
type Building struct {
	IDBuilding      int    `storm:"id,increment"` //The ID of the building
	BuildingName    string `storm:"index"`        //the Name of the building
	BuildingOrder   int    // the order number of the building, 1 is first in the list..
	BuildingPicture string // a presentational picture for the building
}
