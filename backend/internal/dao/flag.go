package dao

import "database/sql"

type Flag struct {
	ID			uint32			`db:"id" json:"id"`
	ProjectID	uint32			`db:"project_id" json:"project_id"`
	Name		string			`db:"name" json:"name"`
	Active	bool				`db:"active" json:"active"`
	UpdateTime	sql.NullTime	`db:"update_time" json:"updated_at"`
}

type FlagContext struct {
	ID			uint32			`db:"id" json:"id"`
	FlagID		uint32  		`db:"flag_id" json:"flag_id"`
	Name		string			`db:"name" json:"name"`
	Condition	string  		`db:"condition" json:"condition"`
	Value		string			`db:"value" json:"value"`
	UpdateTime	sql.NullTime	`db:"update_time" json:"updated_at"`
}