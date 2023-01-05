package types

type Bracket struct {
	Username string `json:"username" bson:"username"`
	AfcWildCard1Winner string `json:"afc_wild_card_1_winner" bson:"afc_wild_card_1_winner"`
	AfcWildCard2Winner string `json:"afc_wild_card_2_winner" bson:"afc_wild_card_2_winner"`
	AfcWildCard3Winner string `json:"afc_wild_card_3_winner" bson:"afc_wild_card_3_winner"`
	AfcDivisionalRound1Winner string `json:"afc_divisional_round_1_winner" bson:"afc_divisional_round_1_winner"`
	AfcDivisionalRound2Winner string `json:"afc_divisional_round_2_winner" bson:"afc_divisional_round_2_winner"`
	AfcConferenceChampion string `json:"afc_conference_champion" bson:"afc_conference_champion"`
	NfcWildCard1Winner string `json:"nfc_wild_card_1_winner" bson:"nfc_wild_card_1_winner"`
	NfcWildCard2Winner string `json:"nfc_wild_card_2_winner" bson:"nfc_wild_card_2_winner"`
	NfcWildCard3Winner string `json:"nfc_wild_card_3_winner" bson:"nfc_wild_card_3_winner"`
	NfcDivisionalRound1Winner string `json:"nfc_divisional_round_1_winner" bson:"nfc_divisional_round_1_winner"`
	NfcDivisionalRound2Winner string `json:"nfc_divisional_round_2_winner" bson:"nfc_divisional_round_2_winner"`
	NfcConferenceChampion string `json:"nfc_conference_champion" bson:"nfc_conference_champion"`
	SuperBowlChampion string `json:"super_bowl_champion" bson:"super_bowl_champion"`
	FinalScoreSum int `json:"final_score_sum" bson:"final_score_sum"`
}