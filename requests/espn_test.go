package requests_test

import (
	"fmt"
	"testing"

	"github.com/supermaxio/nflplayoffbracket/requests"
)

func TestEspnSchedule(t *testing.T) {
	testBody, err := requests.EspnScoreboard()
	if err != nil {
		fmt.Print(err)
		t.Fatal("Failed to retrieve schedule")
	}

	if len(testBody.Leagues) == 0 {
		t.Fatal("Fail on server")
	}
}

func TestEspnTeamRecord(t *testing.T) {
	teamId := "2"
	testBody, err := requests.EspnTeamRecord(teamId)
	if err != nil {
		fmt.Print(err)
		t.Fatal("Failed to retrieve schedule")
	}

	if len(testBody.Items) == 0 {
		t.Fatal("Fail on server")
	}
}

func TestEspnTeamRecord33(t *testing.T) {
	teamId := "33"
	testBody, err := requests.EspnTeamRecord(teamId)
	if err != nil {
		fmt.Print(err)
		t.Fatal("Failed to retrieve schedule")
	}

	if len(testBody.Items) == 0 {
		t.Fatal("Fail on server")
	}
}

func TestEspnTeamRecord34(t *testing.T) {
	teamId := "34"
	testBody, err := requests.EspnTeamRecord(teamId)
	if err != nil {
		fmt.Print(err)
		t.Fatal("Failed to retrieve schedule")
	}

	if len(testBody.Items) == 0 {
		t.Fatal("Fail on server")
	}
}

func TestEspnTeamRecord32(t *testing.T) {
	teamId := "32"
	testBody, err := requests.EspnTeamRecord(teamId)
	if err != nil {
		fmt.Print(err)
		t.Fatal("Failed to retrieve schedule")
	}

	if len(testBody.Items) != 0 {
		t.Fatal("Fail on server")
	}
}

func TestEspnTeamRecord31(t *testing.T) {
	teamId := "31"
	testBody, err := requests.EspnTeamRecord(teamId)
	if err != nil {
		fmt.Print(err)
		t.Fatal("Failed to retrieve schedule")
	}

	if len(testBody.Items) != 0 {
		t.Fatal("Fail on server")
	}
}
