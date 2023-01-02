package types

type TeamStanding struct {
	ID           string `json:"id"`
	Rank         int    `json:"rank"`
	Record       string `json:"record"`
	Location     string `json:"location"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	DisplayName  string `json:"display_name"`
	Conference   string `json:"conference"`
}

type ConferenceStandingResponse struct {
	Conference  string         `json:"conference"`
	RankedTeams []TeamStanding `json:"teams"`
}

type TeamStandingUpdateDB struct {
	Record       string `bson:"record"`
	Rank         int    `bson:"rank"`
	Location     string `bson:"location"`
	Name         string `bson:"name"`
	Abbreviation string `bson:"abbreviation"`
	DisplayName  string `bson:"display_name"`
	Conference   string `bson:"conference"`
}
