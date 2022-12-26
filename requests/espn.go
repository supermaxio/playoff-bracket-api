package requests

import (
	"encoding/json"
	"io/ioutil"

	"net/http"
	"time"

	"github.com/supermaxio/nflplayoffbracket/types"
)

func EspnScoreboard() (types.Scoreboard, error) {
	scoreboardResponse := types.Scoreboard{}
	req, err := http.NewRequest("GET", "http://site.api.espn.com/apis/site/v2/sports/football/nfl/scoreboard", nil)
	if err != nil {
		return types.Scoreboard{}, err
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	resp, err := client.Do(req)
	if err != nil {
		return types.Scoreboard{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &scoreboardResponse)
	if err != nil {
		return types.Scoreboard{}, err
	}

	return scoreboardResponse, nil
}
