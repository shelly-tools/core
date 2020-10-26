package models

// User represents a User for the UI or the API
type User struct {
	IDUser        int    `storm:"id,unique"` //unique member ID
	Username      string `storm:"index"`     // username, will be used for login
	Surname       string // the users surname
	IDMemberGroup int    // the ID of the primary membergroup the user belongs to
	Name          string // the users last name
	EMail         string // the users E-Mail address
	Password      string // the password
}
