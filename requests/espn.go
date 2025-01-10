package requests

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/supermaxio/nflplayoffbracket/types"
)

func EspnScoreboard() (types.EspnScoreboard, error) {
	scoreboardResponse := types.EspnScoreboard{}
	resp, err := ApiGetRequest("http://site.api.espn.com/apis/site/v2/sports/football/nfl/scoreboard")
	if err != nil {
		return types.EspnScoreboard{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.EspnScoreboard{}, err
	}

	err = json.Unmarshal(body, &scoreboardResponse)
	if err != nil {
		return types.EspnScoreboard{}, err
	}

	return scoreboardResponse, nil
}

func EspnTeamRecord(teamId string) (types.EspnTeamRecord, error) {
	teamRecordResponse := types.EspnTeamRecord{}
	resp, err := ApiGetRequest(fmt.Sprintf("http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2024/types/2/teams/%s/record?lang=en&region=us", teamId))
	if err != nil {
		return types.EspnTeamRecord{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.EspnTeamRecord{}, err
	}

	err = json.Unmarshal(body, &teamRecordResponse)
	if err != nil {
		return types.EspnTeamRecord{}, err
	}

	return teamRecordResponse, nil
}
