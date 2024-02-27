package utils

import (
	"Middleware-test/internal/models"
	"encoding/json"
	"log"
	"os"
)

var jsonFile = Path + "data/users.json"

// retrieveUsers
// retrieves all models.User present in jsonFile and stores them in a slice of models.User.
// It returns the slice of models.User and an error.
func retrieveUsers() ([]models.User, error) {
	var users []models.User

	data, err := os.ReadFile(jsonFile)

	if len(data) == 0 {
		return nil, nil
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// changeUsers
// overwrites jsonFile with `users` in json format.
func changeUsers(users []models.User) {
	data, errJSON := json.Marshal(users)
	if errJSON != nil {
		log.Fatal("log: CreateUser()\t JSON Marshall error!\n", errJSON)
	}
	errWrite := os.WriteFile(jsonFile, data, 0666)
	if errWrite != nil {
		log.Fatal("log: CreateUser()\t WriteFile error!\n", errWrite)
	}
}

// GetIdNewUser
// returns first unused id in jsonFile.
func GetIdNewUser() int {
	users, err := retrieveUsers()
	if err != nil {
		log.Fatal("log: retrieveUsers() error!\n", err)
	}
	var id int
	var idFound bool
	for id = 1; !idFound; id++ {
		idFound = true
		for _, user := range users {
			if user.Id == id {
				idFound = false
			}
		}
	}
	id--
	return id
}

// CreateUser
// adds the models.User `newUser` to jsonFile.
func CreateUser(newUser models.User) {
	users, err := retrieveUsers()
	if err != nil {
		log.Fatal("log: retrieveUsers() error!\n", err)
	}
	users = append(users, newUser)
	changeUsers(users)
}

// removeUser
// remove the models.User which models.User.Id is sent in argument from jsonFile.
func removeUser(id int) {
	users, err := retrieveUsers()
	if err != nil {
		log.Fatal("log: retrieveUsers() error!\n", err)
	}
	for i, user := range users {
		if user.Id == id {
			users = append(users[:i], users[i+1:]...)
		}
	}
	changeUsers(users)
}

// SelectUser
// returns the models.User which models.User.Username matches the `username` argument.
func SelectUser(username string) (models.User, bool) {
	var user models.User
	users, err := retrieveUsers()
	if err != nil {
		log.Fatal("log: retrieveUsers() error!\n", err)
	}
	var ok bool
	for _, singleUser := range users {
		if singleUser.Username == username {
			ok = true
			user = singleUser
		}
	}
	return user, ok
}

// updateUser
// modifies the models.User in jsonFile that matches
// `updatedUser`'s Id with `updatedUser`'s content.
func updateUser(updatedUser models.User) {
	users, err := retrieveUsers()
	if err != nil {
		log.Fatal("log: retrieveUsers() error!\n", err)
	}
	for i, user := range users {
		if user.Id == updatedUser.Id {
			users[i] = updatedUser
		}
	}
	changeUsers(users)
}
