package service

import (
	"github.com/supermaxio/nflplayoffbracket/constants"
	"github.com/supermaxio/nflplayoffbracket/database"
	"github.com/supermaxio/nflplayoffbracket/requests"
	"github.com/supermaxio/nflplayoffbracket/types"
	"github.com/supermaxio/nflplayoffbracket/util"
)

func GetStandings() (conferenceStandings []types.ConferenceStandingResponse, err error) {
	afcConferenceStandingResponse := types.ConferenceStandingResponse{Conference: constants.AFC}
	nfcConferenceStandingResponse := types.ConferenceStandingResponse{Conference: constants.NFC}

	scoreboard, _ := requests.EspnScoreboard()
	teamsToDb := []types.TeamStanding{}

	for _, event := range scoreboard.Events {
		for _, competitor := range event.Competitions[0].Competitors {
			competitorRecord := ""
			for _, record := range competitor.Records {
				if record.Name == "overall" {
					competitorRecord = record.Summary
				}
			}

			team := types.TeamStanding{
				ID:           competitor.Team.ID,
				DisplayName:  competitor.Team.DisplayName,
				Abbreviation: competitor.Team.Abbreviation,
				Name:         competitor.Team.Name,
				Location:     competitor.Team.Location,
				Record:       competitorRecord,
			}

			// Query for rank of team
			team.Rank, err = getRank(team)
			if err != nil {
				return
			}

			// Set team in array in the correct rank
			switch util.TeamConference(team.Name) {
			case constants.AFC:
				team.Conference = constants.AFC
				afcConferenceStandingResponse.RankedTeams = append(afcConferenceStandingResponse.RankedTeams, team)
			case constants.NFC:
				team.Conference = constants.NFC
				nfcConferenceStandingResponse.RankedTeams = append(nfcConferenceStandingResponse.RankedTeams, team)
			}

			teamsToDb = append(teamsToDb, team)
		}
	}

	conferenceStandings = append(conferenceStandings, afcConferenceStandingResponse, nfcConferenceStandingResponse)

	database.RefreshStandings(teamsToDb)

	return
}

func getRank(team types.TeamStanding) (rank int, err error) {
	teamRecord, err := requests.EspnTeamRecord(team.ID)
	if err != nil {
		return
	}

	for _, item := range teamRecord.Items {
		if item.Name == "overall" {
			for _, stats := range item.Stats {
				if stats.Name == "playoffSeed" {
					rank = int(stats.Value)
					break
				}
			}

			break
		}
	}

	return
}
