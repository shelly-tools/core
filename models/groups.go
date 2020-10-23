package models

// Groups adawda
type Groups struct {
	IDGroup   int    `storm:"id,unique,increment"`
	Groupname string `storm:"index"`
}
