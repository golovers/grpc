package auth

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"errors"
	"time"
	"log"
)

var (
	HMAC_SECRET = []byte("my-secret")
	USERS       = []string{"pthethanh", "jack"}
)

func Auth(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx);
	if !ok {
		return context.Background(), errors.New("Access denied")
	}

	tokenString := ""
	if tokens, ok := md["authorization"]; ok {
		tokenString = tokens[0]
	} else {
		return ctx, errors.New("Access denied")
	}
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// HMAC_SECRET is a []byte containing your secret, e.g. []byte("my_secret_key")
		return HMAC_SECRET, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		expireTime := claims["expire"].(float64)
		expire := time.Unix(int64(expireTime), 0)
		if isValidUser(username, expire) {
			log.Println("Login successfully")
			return ctx, nil
		} else {
			return ctx, errors.New("Invalid token")
		}
	} else {
		return ctx, err
	}

	return ctx, nil
}

func isValidUser(username string, expire time.Time) bool {
	return isRegisteredUser(username) && isExpired(expire)
}

func isRegisteredUser(username string) bool {
	for _, u := range USERS {
		if username == u {
			return true
		}
	}
	return false
}

func isExpired(t time.Time) bool {
	return t.Sub(time.Now()) > time.Duration(0*time.Second)
}
