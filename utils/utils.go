package utils

import (
	"encoding/json"
	"golang-redis/domains"
)

// Converts from []byte to a json object according to the User struct.
func ToJson(val []byte) domains.User {
	user := domains.User{}
	err := json.Unmarshal(val, &user)
	if err != nil {
		panic(err)
	}
	return user
}
