package dao

import (
	"database/sql"
)

type User struct {
	ID			uint32			`db:"id"`
	Email		string 			`db:"email"`
	Password 	string 			`db:"password"`
	LastLogin	sql.NullTime	`db:"last_login"`
	FirstName 	string			`db:"first_name"`
	LastName 	string			`db:"last_name"`
	RoleName	string			`db:"role_name"`
	RoleLevel 	uint32			`db:"role_level"`
}

type UserProfile struct {
	UserID		uint32 	`db:"user_id"`
	FirstName 	string	`db:"first_name"`
	LastName 	string	`db:"last_name"`
	PhoneNumber string	`db:"phone_number"`
}

type RoleType struct{
	Name	string
	Level 	uint32
}

type UserRole struct {
	UserID		uint32
	RoleID		uint32
}