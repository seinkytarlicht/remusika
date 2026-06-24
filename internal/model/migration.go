package model

type Migration struct {
	Id         int64 `db:"Id"`
	LastUpdate int64 `db:"LastUpdate"`
}
