package util
 
import (
	"strings"
)
 
// BearerAuthHeader validates incoming `r.Header.Get("Authorization")` header
// and returns token otherwise an empty string.
func BearerAuthHeader(authHeader string) string {
	if authHeader == "" {
		return ""
	}
 
	parts := strings.Split(authHeader, "Bearer")
	if len(parts) != 2 {
		return ""
	}
 
	token := strings.TrimSpace(parts[1])
	if len(token) < 1 {
		return ""
	}
 
	return token
}