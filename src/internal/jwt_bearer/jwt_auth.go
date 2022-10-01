package jwt_bearer

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/golang-jwt/jwt"
	"log"
	"simple-game-golang/src/model"
	"time"
)

const (
	SECRET_JWT_KEY     = "e18924e1982e4wqa4sd89asd"
	JWT_EXPIRE_MINUTES = 60
	JWT_ALGORITHM      = "HS512"
)

func GenerateToken(user model.User) string {
	claims := jwt.MapClaims{
		"user_id":  user.Id,
		"username": user.Username,
		"role":     user.Role,
	}
	jwtauth.SetExpiry(claims, time.Now().Add(time.Minute*JWT_EXPIRE_MINUTES))
	jwtauth.SetIssuedAt(claims, time.Now())

	_, token, _ := jwtauth.New(JWT_ALGORITHM, []byte(SECRET_JWT_KEY), nil).Encode(claims)

	return token
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecret := []byte(SECRET_JWT_KEY)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}

func TokenGetClaims(token string) jwt.MapClaims {
	claims, _ := ExtractClaims(token)

	if claims == nil {
		return nil
	}
	return claims
}
