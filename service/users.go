package service

import (
	"errors"

	"github.com/supermaxio/nflplayoffbracket/database"
	"github.com/supermaxio/nflplayoffbracket/types"
)

func GetAllUsers() ([]types.UserResponse, error) {
	usersDB, err := database.GetUsers()
	if err != nil {
		err = errors.New("unable to find user")
		return []types.UserResponse{}, err
	}

	var usersToReturn []types.UserResponse
	for _, userDB := range usersDB {
		usersToReturn = append(usersToReturn, userDB.Response())
	}
	return usersToReturn, nil
}

func GetUser(username string) (types.UserResponse, error) {
	userDB, err := database.FindUser(username)
	if err != nil {
		err = errors.New("unable to find user")
		return types.UserResponse{}, err
	}

	return userDB.Response(), nil
}

func CreateUser(user types.User) (types.UserResponse, error) {
	userDB, err := database.CreateUser(user)
	if err != nil {
		err = errors.New("unable to create user")
		return types.UserResponse{}, err
	}

	return userDB.Response(), nil
}

func UpdateUser(username string, userUpdate types.UserUpdate) (types.UserResponse, error) {
	//TODO add validation

	userDB, err := database.UpdateUser(username, userUpdate)
	if err != nil {
		err = errors.New("unable to Update user")
		return types.UserResponse{}, err
	}

	return userDB.Response(), nil
}

func DeleteUser(username string) (err error) {
	results, err := database.DeleteUser(username)
	if results != 1 || err != nil {
		err = errors.New("unable to delete user")
		return
	}

	return
}
