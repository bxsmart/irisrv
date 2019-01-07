package basic

import (
	"github.com/dgrijalva/jwt-go"
	jwtmw "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"time"
)

const (
	// default expire 7 * 24 Hours
	DefaultExpire = 24 * time.Hour
)

type JwtOptions struct {
	Secret             string
	AccessTokenExpire  int64
	RefreshTokenExpire int64
}

func UnauthorizedError(ctx iris.Context, err string) {
	jwtmw.OnError(ctx, err)
}

func JWTTokenExtractor(ctx iris.Context) (string, error) {
	return jwtmw.FromAuthHeader(ctx)
}

func JWTAuth(app *iris.Application, option JwtOptions) {
	jwtHandler := jwtmw.New(jwtmw.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(option.Secret), nil
		},
		// When set, the middleware verifies that tokens are signed with the specific signing algorithm
		// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
		// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		ContextKey:          jwtmw.DefaultContextKey,
		ErrorHandler:        UnauthorizedError,
		CredentialsOptional: false,
		Extractor:           JWTTokenExtractor,
		Debug:               false,
		EnableAuthOnOptions: false,
		SigningMethod:       jwt.SigningMethodHS256,
		Expiration:          false,
	})

	app.Use(jwtHandler.Serve)
}
