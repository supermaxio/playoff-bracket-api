package service

import (
	"fmt"
	"log"
	"math"
	"reflect"
	"sort"
	"strconv"

	"github.com/supermaxio/nflplayoffbracket/database"
	"github.com/supermaxio/nflplayoffbracket/requests"
	"github.com/supermaxio/nflplayoffbracket/types"
)

func GetGames() ([]types.Game, error) {
	gamesInDB, err := database.GetGames()
	if err != nil {
		return []types.Game{}, err
	}

	return gamesInDB, nil
}

func RefreshScores() ([]types.Game, error) {
	gamesToSave := []types.Game{}
	scoreboard, _ := requests.EspnScoreboard()

	teamsInRank, err := GetPlayoffStandings()
	if err != nil {
		return []types.Game{}, err
	}

	gamesInDB, err := database.GetGames()
	if err != nil {
		return []types.Game{}, err
	}

	bracketsInDB, err := database.GetBrackets()
	if err != nil {
		return []types.Game{}, err
	}

	usersInDB, err := database.GetUsers()
	if err != nil {
		return []types.Game{}, err
	}

	// Use the scoreboard to make a game object
	if scoreboard.Season.Type == 3 {
		for _, event := range scoreboard.Events {
			gameID, _ := strconv.Atoi(event.ID)
			awayTeamName, awayTeamScore, isAwayTeamWinner := getTeamInfo(event.Competitions[0].Competitors, "away")
			homeTeamName, homeTeamScore, isHomeTeamWinner := getTeamInfo(event.Competitions[0].Competitors, "home")
			winner := ""
			if isAwayTeamWinner {
				winner = awayTeamName
			}
			if isHomeTeamWinner {
				winner = homeTeamName
			}

			gameName := getBracketGameName(homeTeamName, scoreboard.Week.Number, teamsInRank)

			game := types.Game{
				ID:            gameID,
				AwayTeam:      awayTeamName,
				AwayTeamScore: awayTeamScore,
				HomeTeam:      homeTeamName,
				HomeTeamScore: homeTeamScore,
				WeekNumber:    scoreboard.Week.Number,
				Finished:      (isAwayTeamWinner || isHomeTeamWinner),
				Winner:        winner,
				BracketName:   gameName,
			}

			gamesToSave = append(gamesToSave, game)
		}
	}

	var usersToSave []types.User
	var scoresList []int

	// Calculate bracket scores from games
	for _, bracket := range bracketsInDB {
		for _, user := range usersInDB {
			if bracket.Username == user.Username {
				newScore, newTieBreaker := getScoreFromBracket(bracket, gamesToSave)
				scoresList = append(scoresList, newScore)
				usersToSave = append(usersToSave, types.User{
					Username:   user.Username,
					Score:      newScore,
					TieBreaker: newTieBreaker,
				})
			}
		}
	}

	// rank users by score
	setRanksForUsers(scoresList, usersToSave)

	// save users
	for _, user := range usersToSave {
		for _, userInDB := range usersInDB {
			if user.Username == userInDB.Username {
				if !reflect.DeepEqual(user, userInDB) {
					update := types.UserUpdate{
						Score:      user.Score,
						Rank:       user.Rank,
						TieBreaker: user.TieBreaker,
					}

					database.UpdateUser(user.Username, update)
				}
			}
		}
	}

	// Update game in db if exists and changed or create game
	for _, gameToSave := range gamesToSave {
		isNewGame := true
		for _, gameInDB := range gamesInDB {
			if gameToSave.ID == gameInDB.ID {
				if !reflect.DeepEqual(gameToSave, gameInDB) {
					_, err := database.UpdateGame(gameToSave.ID, gameToSave)
					if err != nil {
						return []types.Game{}, err
					}
				}

				isNewGame = false
				break
			}
		}

		if isNewGame {
			database.CreateGame(gameToSave)
			if err != nil {
				return []types.Game{}, err
			}
		}
	}

	return gamesToSave, nil
}

func getTeamInfo(teams []types.Competitor, homeAway string) (teamName string, score int, isWinner bool) {
	for _, team := range teams {
		if team.HomeAway == homeAway {
			teamScore, _ := strconv.Atoi(team.Score)
			return team.Team.Name, teamScore, team.Winner
		}
	}

	return
}

func getBracketGameName(homeTeam string, weekNumber int, teamsInRank []types.ConferenceStandingResponse) string {
	switch weekNumber {
	case 1:
		switch homeTeam {
		case teamsInRank[0].RankedTeams[1].Name:
			return "afc_wild_card_1"
		case teamsInRank[0].RankedTeams[2].Name:
			return "afc_wild_card_2"
		case teamsInRank[0].RankedTeams[3].Name:
			return "afc_wild_card_3"
		case teamsInRank[1].RankedTeams[1].Name:
			return "nfc_wild_card_1"
		case teamsInRank[1].RankedTeams[2].Name:
			return "nfc_wild_card_2"
		case teamsInRank[1].RankedTeams[3].Name:
			return "nfc_wild_card_3"
		default:
			return ""
		}
	case 2:
		switch homeTeam {
		case teamsInRank[0].RankedTeams[0].Name:
			return "afc_divisional_round_1"
		case teamsInRank[0].RankedTeams[1].Name,
			teamsInRank[0].RankedTeams[2].Name,
			teamsInRank[0].RankedTeams[3].Name,
			teamsInRank[0].RankedTeams[4].Name:
			return "afc_divisional_round_2"
		case teamsInRank[1].RankedTeams[0].Name:
			return "nfc_divisional_round_1"
		case teamsInRank[1].RankedTeams[1].Name,
			teamsInRank[1].RankedTeams[2].Name,
			teamsInRank[1].RankedTeams[3].Name,
			teamsInRank[1].RankedTeams[4].Name:
			return "nfc_divisional_round_2"
		default:
			return ""
		}
	case 3:
		switch homeTeam {
		case teamsInRank[0].RankedTeams[0].Name,
			teamsInRank[0].RankedTeams[1].Name,
			teamsInRank[0].RankedTeams[2].Name,
			teamsInRank[0].RankedTeams[3].Name,
			teamsInRank[0].RankedTeams[4].Name,
			teamsInRank[0].RankedTeams[5].Name:
			return "afc_conference_championship"
		case teamsInRank[1].RankedTeams[0].Name,
			teamsInRank[1].RankedTeams[1].Name,
			teamsInRank[1].RankedTeams[2].Name,
			teamsInRank[1].RankedTeams[3].Name,
			teamsInRank[1].RankedTeams[4].Name,
			teamsInRank[1].RankedTeams[5].Name:
			return "nfc_conference_championship"
		default:
			return ""
		}
	case 4:
		return "pro_bowl"
	case 5:
		return "suber_bowl"
	default:
		return ""
	}
}

func getScoreFromBracket(bracket types.Bracket, games []types.Game) (newScore int, newTieBreaker int) {
	newTieBreaker = bracket.FinalScoreSum

	for _, game := range games {
		switch game.WeekNumber {
		case 1:
			if bracket.AfcWildCard1Winner == game.Winner {
				newScore += 3
			}
			if bracket.AfcWildCard2Winner == game.Winner {
				newScore += 3
			}
			if bracket.AfcWildCard3Winner == game.Winner {
				newScore += 3
			}
			if bracket.NfcWildCard1Winner == game.Winner {
				newScore += 3
			}
			if bracket.NfcWildCard2Winner == game.Winner {
				newScore += 3
			}
			if bracket.NfcWildCard3Winner == game.Winner {
				newScore += 3
			}
		case 2:
			if bracket.AfcDivisionalRound1Winner == game.Winner {
				newScore += 6
			}
			if bracket.AfcDivisionalRound2Winner == game.Winner {
				newScore += 6
			}
			if bracket.NfcDivisionalRound1Winner == game.Winner {
				newScore += 6
			}
			if bracket.NfcDivisionalRound2Winner == game.Winner {
				newScore += 6
			}
		case 3:
			if bracket.AfcConferenceChampion == game.Winner {
				newScore += 12
			}
			if bracket.NfcConferenceChampion == game.Winner {
				newScore += 12
			}
		case 5:
			if bracket.SuperBowlChampion == game.Winner {
				newScore += 24
			}

			newTieBreaker = int(math.Abs(float64(game.AwayTeamScore + game.HomeTeamScore - bracket.FinalScoreSum)))
		}

	}

	return
}

func rank(scores []int, isReversed bool) []int {
	// Create a copy of the scores slice, so we don't modify the original
	scoresCopy := make([]int, len(scores))
	copy(scoresCopy, scores)

	// Create a mapping of scores to ranks
	scoreMap := make(map[int]int, len(scores))
	sort.Ints(scoresCopy)
	if !isReversed {
		reverse(scoresCopy)
	}
	log.Println(scoresCopy)
	for i := range scoresCopy {
		if i != 0 && scoresCopy[i] == scoresCopy[i-1] {
			scoreMap[scoresCopy[i]] = scoreMap[scoresCopy[i-1]]
		} else {
			scoreMap[scoresCopy[i]] = i + 1
		}
	}

	// Use the mapping to assign ranks to the original scores
	ranks := make([]int, len(scores))
	for i := range scores {
		ranks[i] = scoreMap[scores[i]]
	}
	return ranks
}

func reverse(a []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
}

// rank users by score
func setRanksForUsers(scoresList []int, usersToSave []types.User) {

	log.Printf(fmt.Sprintf(" ************ "))
	log.Printf(fmt.Sprintf(" scoreslist %d", scoresList))
	log.Printf(fmt.Sprintf(" users %d", len(usersToSave)))

	// rank users by score
	rankOrder := rank(scoresList, false)
	log.Printf(fmt.Sprintf(" rank order %d", rankOrder))

	for i := range usersToSave {
		usersToSave[i].Rank = rankOrder[i]
	}

	// tie breaker rank
	for i, user := range usersToSave {
		var tieBreakerList []int
		var tieBreakerUserIndex []int
		for j, userToCompare := range usersToSave {
			if user.Username != userToCompare.Username {
				if user.Rank == userToCompare.Rank {
					tieBreakerList = append(tieBreakerList, userToCompare.TieBreaker)
					tieBreakerUserIndex = append(tieBreakerUserIndex, j)
				}
			}
		}

		if len(tieBreakerList) > 0 {
			// add this users data
			tieBreakerList = append(tieBreakerList, user.TieBreaker)
			tieBreakerUserIndex = append(tieBreakerUserIndex, i)

			tieBreakerOrder := rank(tieBreakerList, true)
			for j, index := range tieBreakerUserIndex {
				usersToSave[index].Rank += tieBreakerOrder[j]
				usersToSave[index].Rank -= 1
			}
		}
	}
}
