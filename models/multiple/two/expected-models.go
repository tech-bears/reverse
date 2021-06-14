package models

type A struct {
	Id int `xorm:"integer"`
}

type B struct {
	Id int `xorm:"INTEGER"`
}

type C struct {
	Id int `xorm:"INTEGER"`
}

type D struct {
	Id int `xorm:"INTEGER"`
}
