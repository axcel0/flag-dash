package dao

import "database/sql"

type Flag struct {
	ID			uint32			`db:"id"`
	ProjectID	uint32			`db:"project_id"`
	Name		string			`db:"name"`
	Active	bool				`db:"active"`
	UpdateTime	sql.NullTime	`db:"update_time"`
}

type FlagContext struct {
	ID			uint32			`db:"id"`
	FlagID		uint32  		`db:"flag_id"`
	Name		string			`db:"name"`
	Condition	string  		`db:"condition"`
	Value		string			`db:"value"`
	UpdateTime	sql.NullTime	`db:"update_time"`
}