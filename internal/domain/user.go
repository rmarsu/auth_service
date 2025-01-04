package domain

import "time"

type User struct {
	Id        int64
	Email     string
	Username  string
	Password  []byte
	CreatedAt time.Time
}
