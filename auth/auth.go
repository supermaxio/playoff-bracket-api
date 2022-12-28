package auth

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/supermaxio/nflplayoffbracket/config"
	"github.com/supermaxio/nflplayoffbracket/constants"
	"github.com/supermaxio/nflplayoffbracket/database"
	"github.com/supermaxio/nflplayoffbracket/types"
	"golang.org/x/crypto/bcrypt"
)

// Create a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Token struct {
	Token          string    `json:"token"`
	ExpirationTime time.Time `json:"expiration_time"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user types.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	var res types.ResponseResult
	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	result := database.FindUser(user.Username)

	if result.Username == "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)

		if err != nil {
			res.Error = "Error While Hashing Password, Try Again"
			json.NewEncoder(w).Encode(res)
			return
		}
		user.Password = string(hash)

		_ = database.CreateUser(user)
		if err != nil {
			res.Error = "Error While Creating User, Try Again"
			json.NewEncoder(w).Encode(res)
			return
		}

		res.Result = "Register Successful"
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Result = "Username already exists!!"
	json.NewEncoder(w).Encode(res)
	return
}

// Create the Signin handler
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	var res types.ResponseResult

	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the expected password from our in memory map
	resultUser := database.FindUser(creds.Username)
	// if err != nil {
	if resultUser.Username == "" {
		res.Error = "Invalid username"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	errf := bcrypt.CompareHashAndPassword([]byte(resultUser.Password), []byte(creds.Password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		var res = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(res)
		return
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte(config.GetJwtKey()))
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Finally, we set the client cookie for constants.COOKIE_TOKEN as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    constants.COOKIE_TOKEN,
		Value:   tokenString,
		Expires: expirationTime,
	})

	json.NewEncoder(w).Encode(Token{tokenString, expirationTime})
}

func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(constants.COOKIE_TOKEN)
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	claims := &Claims{}
	jwtKey := []byte(config.GetJwtKey())
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// (END) The code until this point is the same as the first part of the `Welcome` route

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Set the new token as the users `token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    constants.COOKIE_TOKEN,
		Value:   tokenString,
		Expires: expirationTime,
	})

	json.NewEncoder(w).Encode(Token{tokenString, expirationTime})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// immediately clear the token cookie
	http.SetCookie(w, &http.Cookie{
		Name:    constants.COOKIE_TOKEN,
		Expires: time.Now(),
	})
}

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// We can obtain the session token from the requests cookies, which come with every request
		c, err := r.Cookie(constants.COOKIE_TOKEN)

		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Get the JWT string from the cookie
		tknStr := c.Value

		// Initialize a new instance of `Claims`
		claims := &Claims{}

		jwtKey := []byte(config.GetJwtKey())
		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// Finally, return the welcome message to the user, along with their
		// username given in the token

		ctx := context.WithValue(r.Context(), "user", claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CorsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		switch config.GetEnv() {
		case constants.ENV_PROD,
			constants.ENV_STAGE:
			w.Header().Set("Access-Control-Allow-Origin", r.RemoteAddr)
		default:
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		}
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "append,delete,entries,foreach,get,has,keys,set,values,Authorization,content-type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			//handle preflight in here
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
