package models

type A struct {
	Id int `xorm:"'Id' integer"`
}

type B struct {
	Id int `xorm:"'Id' INTEGER"`
}

type C struct {
	Id int `xorm:"'Id' INTEGER"`
}

type D struct {
	Id int `xorm:"'Id' INTEGER"`
}
