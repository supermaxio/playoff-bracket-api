package service

import (
	"reflect"
	"testing"

	"github.com/supermaxio/nflplayoffbracket/types"
)


func TestSetRanksForUsers(t *testing.T) {
	testCases := []struct {
		name        string
		scores      []int
		users       []types.User
		want        []types.User
	}{
		{
			name:        "with score and tie breaker",
			scores:      []int{9, 8, 10, 7, 8, 9},
			users:       []types.User{UserTestdata(1, 0, 55, 9), UserTestdata(2, 0, 33, 8), UserTestdata(3, 0, 44, 10), UserTestdata(4, 0, 77, 7), UserTestdata(5, 0, 66, 8), UserTestdata(6, 0, 22, 9)},
			want:        []types.User{UserTestdata(1, 3, 55, 9), UserTestdata(2, 4, 33, 8), UserTestdata(3, 1, 44, 10), UserTestdata(4, 6, 77, 7), UserTestdata(5, 5, 66, 8), UserTestdata(6, 2, 22, 9)},
		},
		{
			name:        "No tie breaker",
			scores:      []int{9, 8, 10, 7, 8, 9},
			users:       []types.User{UserTestdata(1, 0, 0, 9), UserTestdata(2, 0, 0, 8), UserTestdata(3, 0, 0, 10), UserTestdata(4, 0, 0, 7), UserTestdata(5, 0, 0, 8), UserTestdata(6, 0, 0, 9)},
			want:        []types.User{UserTestdata(1, 2, 0, 9), UserTestdata(2, 4, 0, 8), UserTestdata(3, 1, 0, 10), UserTestdata(4, 6, 0, 7), UserTestdata(5, 4, 0, 8), UserTestdata(6, 2, 0, 9)},
		},
		{
			name:        "No scores",
			scores:      []int{0, 0, 0, 0, 0, 0},
			users:       []types.User{UserTestdata(1, 0, 0, 0), UserTestdata(2, 0, 0, 0), UserTestdata(3, 0, 0, 0), UserTestdata(4, 0, 0, 0), UserTestdata(5, 0, 0, 0), UserTestdata(6, 0, 0, 0)},
			want:        []types.User{UserTestdata(1, 1, 0, 0), UserTestdata(2, 1, 0, 0), UserTestdata(3, 1, 0, 0), UserTestdata(4, 1, 0, 0), UserTestdata(5, 1, 0, 0), UserTestdata(6, 1, 0, 0)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			setRanksForUsers(tc.scores, tc.users)
			if !reflect.DeepEqual(tc.users, tc.want) {
				t.Errorf("rank(%v) = %v, want %v", tc.scores, tc.users, tc.want)
			}
		})
	}
}

func UserTestdata(userNumber int, rank int, tieBreaker int, score int) (user types.User) {
	switch userNumber {
	case 1:
		user = types.User{Username: "1", Rank: rank, TieBreaker: tieBreaker, Score: score}
	case 2:
		user = types.User{Username: "2", Rank: rank, TieBreaker: tieBreaker, Score: score}
	case 3:
		user = types.User{Username: "3", Rank: rank, TieBreaker: tieBreaker, Score: score}
	case 4:
		user = types.User{Username: "4", Rank: rank, TieBreaker: tieBreaker, Score: score}
	case 5:
		user = types.User{Username: "5", Rank: rank, TieBreaker: tieBreaker, Score: score}
	case 6:
		user = types.User{Username: "6", Rank: rank, TieBreaker: tieBreaker, Score: score}
	}
	return user
}
