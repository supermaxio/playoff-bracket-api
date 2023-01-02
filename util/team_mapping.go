package util

import (
	"github.com/supermaxio/nflplayoffbracket/constants"
)

func TeamConference(team_name string) (conference_name string) {
	switch team_name {
	case constants.BILLS,
		constants.BENGALS,
		constants.BROWNS,
		constants.BRONCOS,
		constants.TEXANS,
		constants.COLTS,
		constants.JAGUARS,
		constants.CHIEFS,
		constants.CHARGERS,
		constants.RAIDERS,
		constants.DOLPHINS,
		constants.PATRIOTS,
		constants.RAVENS,
		constants.JETS,
		constants.STEELERS,
		constants.TITANS:
		conference_name = constants.AFC
	case constants.CARDINALS,
		constants.FALCONS,
		constants.PANTHERS,
		constants.BEARS,
		constants.COWBOYS,
		constants.LIONS,
		constants.PACKERS,
		constants.RAMS,
		constants.VIKINGS,
		constants.GIANTS,
		constants.EAGLES,
		constants.SAINTS,
		constants.SEAHAWKS,
		constants.FORTYNINERS,
		constants.BUCCANEERS,
		constants.COMMANDERS:
		conference_name = constants.NFC
	default:
	}

	return
}
