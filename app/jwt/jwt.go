package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"

	"time"
)


//var jwtSecret = []byte(setting.JwtSecret)
var jwtSecret = []byte("sasasa")


type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

//生成token
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)//token时间为三个小时

	claims := Claims{
		username,
		password,
		jwt.StandardClaims {
			ExpiresAt : expireTime.Unix(),
			Issuer : "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

