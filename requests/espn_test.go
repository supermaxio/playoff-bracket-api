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
