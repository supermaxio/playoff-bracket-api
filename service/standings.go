package service

import (
	"sort"

	"github.com/supermaxio/nflplayoffbracket/constants"
	"github.com/supermaxio/nflplayoffbracket/database"
	"github.com/supermaxio/nflplayoffbracket/requests"
	"github.com/supermaxio/nflplayoffbracket/types"
	"github.com/supermaxio/nflplayoffbracket/util"
)

func GetPlayoffStandings() (conferenceStandings []types.ConferenceStandingResponse, err error) {
	afcConferenceStandingResponse := types.ConferenceStandingResponse{Conference: constants.AFC}
	nfcConferenceStandingResponse := types.ConferenceStandingResponse{Conference: constants.NFC}

	// Add only the top 7 teams from each conference to respective structs
	teamsFromDb := database.GetStandings()
	for _, team := range teamsFromDb {
		if team.Rank < 8 {
			switch util.TeamConference(team.Name) {
			case constants.AFC:
				afcConferenceStandingResponse.RankedTeams = append(afcConferenceStandingResponse.RankedTeams, team)
			case constants.NFC:
				nfcConferenceStandingResponse.RankedTeams = append(nfcConferenceStandingResponse.RankedTeams, team)
			}
		}
	}

	// Sort the array of 7 teams per conference
	sort.Slice(afcConferenceStandingResponse.RankedTeams, func(i, j int) bool {
		return afcConferenceStandingResponse.RankedTeams[i].Rank < afcConferenceStandingResponse.RankedTeams[j].Rank
	})
	sort.Slice(nfcConferenceStandingResponse.RankedTeams, func(i, j int) bool {
		return nfcConferenceStandingResponse.RankedTeams[i].Rank < nfcConferenceStandingResponse.RankedTeams[j].Rank
	})

	conferenceStandings = append(conferenceStandings, afcConferenceStandingResponse, nfcConferenceStandingResponse)

	return
}

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
