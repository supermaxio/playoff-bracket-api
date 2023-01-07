package util

import (
	"strings"

	"github.com/supermaxio/nflplayoffbracket/types"
)

// BearerAuthHeader validates incoming `r.Header.Get("Authorization")` header
// and returns token otherwise an empty string.
func BearerAuthHeader(authHeader string) string {
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, "Bearer")
	if len(parts) != 2 {
		return ""
	}

	token := strings.TrimSpace(parts[1])
	if len(token) < 1 {
		return ""
	}

	return token
}

func ValidateBracket(bracket types.Bracket, standings []types.ConferenceStandingResponse) (errList []string) {
	if bracket.FinalScoreSum == 0 {
		errList = append(errList, "missing final score")
	}

	if bracket.AfcWildCard1Winner == "" ||
		bracket.AfcWildCard2Winner == "" ||
		bracket.AfcWildCard3Winner == "" ||
		bracket.AfcDivisionalRound1Winner == "" ||
		bracket.AfcDivisionalRound2Winner == "" ||
		bracket.AfcConferenceChampion == "" ||
		bracket.NfcWildCard1Winner == "" ||
		bracket.NfcWildCard2Winner == "" ||
		bracket.NfcWildCard3Winner == "" ||
		bracket.NfcDivisionalRound1Winner == "" ||
		bracket.NfcDivisionalRound2Winner == "" ||
		bracket.NfcConferenceChampion == "" ||
		bracket.SuperBowlChampion == "" {
		errList = append(errList, "missing entry")
	}

	if bracket.SuperBowlChampion != bracket.AfcConferenceChampion && bracket.SuperBowlChampion != bracket.NfcConferenceChampion {
		errList = append(errList, "incorrect super bowl champion")
	}

	if bracket.AfcConferenceChampion != bracket.AfcDivisionalRound1Winner && bracket.AfcConferenceChampion != bracket.AfcDivisionalRound2Winner {
		errList = append(errList, "incorrect afc conference champion")
	}

	if bracket.NfcConferenceChampion != bracket.NfcDivisionalRound1Winner && bracket.NfcConferenceChampion != bracket.NfcDivisionalRound2Winner {
		errList = append(errList, "incorrect nfc conference champion")
	}

	if bracket.NfcDivisionalRound1Winner != standings[1].RankedTeams[0].Name && bracket.NfcDivisionalRound1Winner != bracket.NfcWildCard1Winner && bracket.NfcDivisionalRound1Winner != bracket.NfcWildCard2Winner && bracket.NfcDivisionalRound1Winner != bracket.NfcWildCard3Winner {
		errList = append(errList, "incorrect nfc divisional 1 winner")
	}

	if bracket.NfcDivisionalRound1Winner != standings[1].RankedTeams[0].Name && bracket.NfcDivisionalRound2Winner != bracket.NfcWildCard1Winner && bracket.NfcDivisionalRound2Winner != bracket.NfcWildCard2Winner && bracket.NfcDivisionalRound2Winner != bracket.NfcWildCard3Winner {
		errList = append(errList, "incorrect nfc divisional 2 winner")
	}

	if bracket.AfcDivisionalRound1Winner != standings[0].RankedTeams[0].Name && bracket.AfcDivisionalRound1Winner != bracket.AfcWildCard1Winner && bracket.AfcDivisionalRound1Winner != bracket.AfcWildCard2Winner && bracket.AfcDivisionalRound1Winner != bracket.AfcWildCard3Winner {
		errList = append(errList, "incorrect afc divisional 1 winner")
	}

	if bracket.AfcDivisionalRound1Winner != standings[0].RankedTeams[0].Name && bracket.AfcDivisionalRound2Winner != bracket.AfcWildCard1Winner && bracket.AfcDivisionalRound2Winner != bracket.AfcWildCard2Winner && bracket.AfcDivisionalRound2Winner != bracket.AfcWildCard3Winner {
		errList = append(errList, "incorrect afc divisional 2 winner")
	}

	return errList
}
