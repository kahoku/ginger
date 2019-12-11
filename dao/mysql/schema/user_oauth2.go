package schema

import "time"

/*
   This code is generated for gendry
   ginger-cli schema -d ginger_db -t user_oauth2
*/

const UserOauth2TableName = "user_oauth2"


// UserOauth2 is a mapping object for user_oauth2 table in mysql
type UserOauth2 struct {
	ID          int       `ddb:"id"`
	Platform    int8      `ddb:"platform"`
	AccessToken string    `ddb:"access_token"`
	OpenID      string    `ddb:"open_id"`
	UnionID     string    `ddb:"union_id"`
	NickName    string    `ddb:"nick_name"`
	Gender      int8      `ddb:"gender"`
	AvatarURL   string    `ddb:"avatar_url"`
	CreateAt    time.Time `ddb:"create_at"`
	UpdateAt    time.Time `ddb:"update_at"`
}

func (*UserOauth2) TableName() string {
	return UserOauth2TableName
}
