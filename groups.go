package main

type Groups struct {
	IDGroup   int    `storm:"id,unique"`
	Groupname string `storm:"index"`
}
