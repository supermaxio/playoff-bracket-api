package service

import (
	"github.com/supermaxio/nflplayoffbracket/database"
	"github.com/supermaxio/nflplayoffbracket/requests"
	"github.com/supermaxio/nflplayoffbracket/types"
)

func GetStandings() (teamStandings []types.TeamStandingResponse, err error) {
	scoreboard, _ := requests.EspnScoreboard()
	teamsToDb := []types.TeamStandingDB{}

	for _, event := range scoreboard.Events {
		for _, competitor := range event.Competitions[0].Competitors {
			competitorRecord := ""
			for _, record := range competitor.Records {
				if record.Name == "overall" {
					competitorRecord = record.Summary
				}
			}

			team := types.TeamStandingResponse{
				TeamId:           competitor.Team.ID,
				DisplayName:      competitor.Team.DisplayName,
				ShortDisplayName: competitor.Team.ShortDisplayName,
				Abbreviation:     competitor.Team.Abbreviation,
				Name:             competitor.Team.Name,
				Location:         competitor.Team.Location,
				Record:           competitorRecord,
				LogoLink:         competitor.Team.Logo,
			}

			teamStandings = append(teamStandings, team)

			teamToDb := types.TeamStandingDB{
				ID: team.TeamId,
				Record: team.Record,
			}

			teamsToDb = append(teamsToDb, teamToDb)
		}
	}

	go database.RefreshStandings(teamsToDb)

	return
}
