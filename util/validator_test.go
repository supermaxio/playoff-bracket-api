package util
 
import "testing"
 
func TestBearerAuthHeader(t *testing.T) {
	tests := []struct{
		name  string
		auth  string
		token string
	}{
		{
			"EmptyInput",
			"",
			"",
		},
		{
			"EmptyStringInput",
			"   ",
			"",
		},
		{
			"BearerWithoutToken",
			"Bearer",
			"",
		},
		{
			"BearerPrefixWithEmptyStringToken",
			"Bearer   ",
			"",
		},
		{
			"WrongPrefixWithToken",
			"Basic token",
			"",
		},
		{
			"WrongBearerPrefixCaseWithToken",
			"BEARER token",
			"",
		},
		{
			"BearerPrefixWithNextLineToken",
			"Bearer \n",
			"",
		},
		{
			"BearerPrefixWithTabToken",
			"Bearer \t",
			"",
		},
		{
			"IncorrectlySpacedValidRequest",
			"   Bearer    token   ",
			"token",
		},
		{
			"CorrectlySpacedValidRequest",
			"Bearer token",
			"token",
		},
	}
 
	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			token := BearerAuthHeader(c.auth)
 
			if c.token != token {
				t.Fatal("expected", c.token, "but got", token)
			}
		})
	}
}