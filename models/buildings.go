package models

// Building such as home or garage
type Building struct {
	ID          int               `storm:"id,increment" json:"id"` // the ID of the building
	Name        string            `storm:"index" json:"name"`      // the Name of the building
	Order       int               `json:"order"`                   // the order number of the building, 1 is first in the list..
	PicturePath string            `json:"picturePath"`             // a presentational picture for the building
	PictureData string            `json:"pictureData"`
	MetaData    map[string]string `json:"metaData"` // some metadata for later functions or grouping
}
