package models

import "time"

type User struct{
	ID uint32
	Nickname string
	Email string
	Password string
	Created_At time.Time
	Updated_At time.Time
}