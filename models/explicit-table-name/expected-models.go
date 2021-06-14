package models

type A struct {
	Id int `xorm:"integer"`
}

func (m *A) TableName() string {
	return "a"
}

type B struct {
	Id int `xorm:"INTEGER"`
}

func (m *B) TableName() string {
	return "b"
}

type C struct {
	Id int `xorm:"INTEGER"`
}

func (m *C) TableName() string {
	return "c"
}

type D struct {
	Id int `xorm:"INTEGER"`
}

func (m *D) TableName() string {
	return "d"
}
