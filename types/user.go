package types

type User struct {
	Username   string `bson:"username"`
	Password   string `bson:"password"`
	Score      int    `bson:"score"`
	Rank       int    `bson:"rank"`
	TieBreaker int    `bson:"tie_breaker"`
	Champion   int    `bson:"champion"`
}

type UserResponse struct {
	Username   string `bson:"username"`
	Score      int    `bson:"score"`
	Rank       int    `bson:"rank"`
	TieBreaker int    `bson:"tie_breaker"`
	Champion   int    `bson:"champion"`
}

type UserUpdate struct {
	Score      int `bson:"score"`
	Rank       int `bson:"rank"`
	TieBreaker int `bson:"tie_breaker"`
}

func (u *User) Response() (userResponse UserResponse) {
	userResponse.Username = u.Username
	userResponse.Rank = u.Rank
	userResponse.Score = u.Score
	userResponse.TieBreaker = u.TieBreaker
	userResponse.Champion = u.Champion
	return
}
