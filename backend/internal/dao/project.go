package dao

import "database/sql"

type Project struct {
	ID		uint32				`json:"id" db:"id"`
	Name 	string				`json:"name" db:"name"`
	AccessKey string			`json:"access_key" db:"access_key"`
	UpdateTime	sql.NullTime	`json:"updated_at" db:"update_time"`
}