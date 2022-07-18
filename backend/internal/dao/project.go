package dao

import "database/sql"

type Project struct {
	ID		uint32				`db:"id"`
	Name 	string				`db:"name"`
	UpdateTime	sql.NullTime	`db:"update_time"`
}