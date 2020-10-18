package main

type User struct {
	IDUser   int    `storm:"id,unique"`
	Username string `storm:"index"`
	Surname  string
	IDGroup  int
	Name     string
	EMail    string
	Password string
}
type Users struct {
	User []User
}
