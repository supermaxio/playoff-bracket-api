package types

type TeamStanding struct {
	ID           string `json:"id" bson:"id"`
	Rank         int    `json:"rank" bson:"rank"`
	Record       string `json:"record" bson:"record"`
	Location     string `json:"location" bson:"location"`
	Name         string `json:"name" bson:"name"`
	Abbreviation string `json:"abbreviation" bson:"abbreviation"`
	DisplayName  string `json:"display_name" bson:"display_name"`
	Conference   string `json:"conference" bson:"conference"`
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
