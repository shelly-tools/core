package models

// User represents a User for the UI or the API
type User struct {
	IDUser   int    `storm:"id,unique"`
	Username string `storm:"index"`
	Surname  string
	IDGroup  int
	Name     string
	EMail    string
	Password string
}
