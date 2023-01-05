package service

import (
	"errors"

	"github.com/supermaxio/nflplayoffbracket/database"
	"github.com/supermaxio/nflplayoffbracket/types"
)

func GetBracket(username string) (bracketToReturn types.Bracket, err error) {
	bracketToReturn, err = database.FindBracket(username)
	if err != nil {
		err = errors.New("unable to find bracket")
		return
	}

	return
}

func CreateBracket(bracket types.Bracket) (bracketToReturn types.Bracket, err error) {
	bracketToReturn, err = database.CreateBracket(bracket)
	if err != nil {
		err = errors.New("unable to create bracket")
		return
	}

	return
}

func UpdateBracket(bracket types.Bracket) (bracketToReturn types.Bracket, err error) {
	bracketToReturn, err = database.UpdateBracket(bracket)
	if err != nil {
		err = errors.New("unable to Update bracket")
		return
	}

	return
}

func DeleteBracket(username string) (err error) {
	results, err := database.DeleteBracket(username)
	if results != 1 || err != nil {
		err = errors.New("unable to delete bracket")
		return
	}

	return
}
