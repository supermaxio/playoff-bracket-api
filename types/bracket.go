package types

type Bracket struct {
	Username string `json:"username" bson:"username"`
	WildCardWinners []string `json:"wild_card_winners" bson:"wild_card_winners"`
	DivisionalRoundWinners []string `json:"divisional_round_winners" bson:"divisional_round_winners"`
	ConferenceChampions []string `json:"conference_champions" bson:"conference_champions"`
	SuperBowlChampion string `json:"super_bowl_champion" bson:"super_bowl_champion"`
	FinalScoreSum int `json:"final_score_sum" bson:"final_score_sum"`
}