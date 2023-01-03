package service

import (
	"errors"

	"github.com/supermaxio/nflplayoffbracket/database"
	"github.com/supermaxio/nflplayoffbracket/types"
)

func GetBracket(username string) (bracketToReturn types.Bracket, err error) {
	bracketToReturn = database.FindBracket(username)

	if bracketToReturn.SuperBowlChampion == "" {
		err = errors.New("unable to find bracket")
	}

	return
}

func CreateBracket(bracket types.Bracket) (bracketToReturn types.Bracket, err error) {
	bracketToReturn = database.CreateBracket(bracket)

	if bracketToReturn.SuperBowlChampion == "" {
		err = errors.New("unable to create bracket")

	}

	return
}

func UpdateBracket(bracket types.Bracket) (bracketToReturn types.Bracket, err error) {
	bracketToReturn = database.UpdateBracket(bracket)

	if bracketToReturn.SuperBowlChampion == "" {
		err = errors.New("unable to Update bracket")

	}

	return
}

func DeleteBracket(username string) (err error) {
	results := database.DeleteBracket(username)

	if results != 1 {
		err = errors.New("unable to delete bracket")
	}

	return
}
