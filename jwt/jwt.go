package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

const JwtSecret = "adsfhgwougf;aksdjhfhewfwhefuyqewfihf"

type UserInfo struct {
	ClientID string   `json:"client_id"`
	Roles    []string `json:"roles"`
	Scope    string   `json:"scope"`
	UserID   string   `json:"user_id"`
}

func main() {
	user := &UserInfo{
		ClientID: "4b81c5c0c9b944f64beff7777eac2b9bbd0fdcaa",
		Roles:    []string{"operatorAttache"},
		Scope:    "operatorAttache",
		UserID:   "liaopujian",
	}
	fmt.Println(userJwtEncode(user))
}

func userJwtEncode(user *UserInfo) (string, error) {
	jwtData := make(map[string]interface{})
	jwtData["client_id"] = user.ClientID
	jwtData["roles"] = user.Roles
	jwtData["scope"] = user.Scope
	jwtData["user_id"] = user.UserID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(jwtData))
	return token.SignedString([]byte(JwtSecret))
}
