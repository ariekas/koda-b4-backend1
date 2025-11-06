package models

type User struct {
	Id       int    `form:"id"`
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
	ProfilePic string `form:"profile"`
	Token string 
}

var Users = []User{}