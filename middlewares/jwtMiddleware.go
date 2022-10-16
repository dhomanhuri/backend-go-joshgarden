package middlewares

import (
	"backend-go/config"
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(userID int, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = userID
	claims["email"] = email
	claims["exp"] = time.Now().Add(5 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SecretKey()))
}

func JWTTokenCheck(c *gin.Context) (int, string, error) {
	jwtToken, errExtract := extractBearerToken(c.GetHeader("Authorization"))
	if errExtract != nil {
		return 0, "", errExtract
	}

	token, errParse := parseToken(jwtToken)
	if errParse != nil {
		return 0, "", errParse
	}

	claims, OK := token.Claims.(jwt.MapClaims)
	if OK {
		userID := claims["userID"].(float64)
		email := claims["email"].(string)

		return int(userID), email, nil
	}
	return 0, "", errors.New("failed extract token")
}

func parseToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("bad signed method received")
		}
		return []byte(config.SecretKey()), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}
	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}
	return jwtToken[1], nil
}
