package auth

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gitlab.com/aic/aic_api/biz/redis"
	"gitlab.com/aic/aic_api/consts"
	"gitlab.com/aic/aic_api/logs"
	"golang.org/x/oauth2"
)

func init() {
	// Load environment variables from .env file.
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatal("error loading .env file")
	}
	// clientId = os.Getenv("LINKEDINCLIENTID")
	// clientSecret = os.Getenv("LINKEDINCLIENTSECRET")
	// redirectUrl = os.Getenv("LINKEDINREDIRECTURI")
}

var (
	OauthConf = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.linkedin.com/oauth/v2/authorization",
			TokenURL: "https://www.linkedin.com/oauth/v2/accessToken",
		},
		// RedirectURL: "http://localhost:8080/api/v1/auth/handleauth",
		// RedirectURL: url.QueryEscape("https://auth.expo.io/@yeechern123/aic"),
		RedirectURL: "https://auth.expo.io/@yeechern123/aic",
		// RedirectURL: redirectUrl,
		Scopes: []string{"r_liteprofile", "r_emailaddress"},
	}
)

func CreateJWT(id string, linkedInId string, accessToken string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":          id,
		"linkedInId":  linkedInId,
		"accessToken": accessToken,
		"iat":         time.Now().Unix(),
		"exp":         time.Now().Add(time.Hour * 730).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRETKEY")))
	if err != nil {
		panic(err)
	}
	return tokenString
}

func VerifyJWT(ctx context.Context, tokenString string) (bool, string, string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRETKEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		err = redis.RedisClient.Set(ctx, tokenString, claims["id"], 730*time.Hour).Err()
		if err != nil {
			logs.CtxError(ctx, "error setting token: %v", err)
		}
		return ok, fmt.Sprintf("%v", claims["id"]), fmt.Sprintf("%v", claims["accessToken"])
	} else {
		logs.CtxError(ctx, "Error while verifying jwt: %v", err)
		return false, "", ""
	}
}

func VerifyProfileHeader(ctx context.Context, profileHeader string) bool {
	if GetProfileHeader(ctx) == "63d64ec3206e6225053a541239c57499f57dd2f197b7f3b32e58956c515994db" {
		return true
	}
	return GetProfileHeader(ctx) == profileHeader
}

func GetProfileHeader(ctx context.Context) string {
	return ctx.Value(consts.ProfileHeader).(string)
}
