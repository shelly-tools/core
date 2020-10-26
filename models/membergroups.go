package models

// MemberGroup is used to group users in order to allows access to features / devices
type MemberGroup struct {
	IDMemberGroup   int    `storm:"id,unique,increment"` // primary key for a member group
	MemberGroupName string `storm:"index"`               // Name of the Group
}
