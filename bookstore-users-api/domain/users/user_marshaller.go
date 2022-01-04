package users

import "encoding/json"

type PublicUser struct {
	Id          int64  `json:"id"`
	Status      string `json:"status"`
	DateCreated string `json:"date_created"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	DateCreated string `json:"date_created"`
}

func (user *User) Marshall(isPublic bool) interface{} {
	userJson, _ := json.Marshal(user)

	if isPublic {
		var publicUser PublicUser
		json.Unmarshal(userJson, &publicUser)
		return publicUser
	}

	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))

	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}

	return result
}
