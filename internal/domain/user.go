package domain

type User struct {
	Id       int64
	Email    string
	Username string
	Password []byte
}
