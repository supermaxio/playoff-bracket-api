package types

type TeamStandingResponse struct {
	TeamId           string `json:"team_id"`
	LeagueRank       int    `json:"league_rank"`
	Conference       string `json:"conference"`
	Division         string `json:"division"`
	ConferenceRank   int    `json:"conference_rank"`
	Record           string `json:"record"`
	Location         string `json:"location"`
	Name             string `json:"name"`
	Abbreviation     string `json:"abbreviation"`
	DisplayName      string `json:"display_name"`
	ShortDisplayName string `json:"short_display_name"`
	LogoLink         string `json:"logo_link"`
}

type TeamStandingDB struct {
	ID     string `bson:"id"`
	Record string `bson:"record"`
}
