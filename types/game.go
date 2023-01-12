package types

type Game struct {
	ID            int    `json:"id" bson:"id"`
	BracketName   string `json:"bracket_name" bson:"bracket_name"`
	WeekNumber    int    `json:"week_number" bson:"week_number"`
	AwayTeamScore int    `json:"away_team_score" bson:"away_team_score"`
	HomeTeamScore int    `json:"home_team_score" bson:"home_team_score"`
	AwayTeam      string `json:"away_team" bson:"away_team"`
	HomeTeam      string `json:"home_team" bson:"home_team"`
	Winner        string `json:"winner" bson:"winner"`
	Finished      bool   `json:"finished" bson:"finished"`
}
