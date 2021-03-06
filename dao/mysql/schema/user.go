package schema

import "time"

/*
   This code is generated by gendry
   ginger-gen schema -d ginger_db -t user_oauth2
*/

const UserTableName = "user"

// User is a mapping object for user table in mysql
type User struct {
	ID       uint      `ddb:"id" json:"id"`
	Name     string    `ddb:"name" json:"name"`
	Age      byte      `ddb:"age" json:"age"`
	Avatar   string    `ddb:"avatar" json:"avatar"`
	Gender   int8      `ddb:"gender" json:"gender"`
	Email    string    `ddb:"email" json:"email"`
	Phone    string    `ddb:"phone" json:"phone"`
	Password string    `ddb:"password" json:"password"`
	Salt     string    `ddb:"salt" json:"salt"`
	Status   int8      `ddb:"status" json:"status"`
	UpdateAt time.Time `ddb:"update_at" json:"update_at"`
	CreateAt time.Time `ddb:"create_at" json:"create_at"`
}

func (*User)TableName() string  {
	return UserTableName
}