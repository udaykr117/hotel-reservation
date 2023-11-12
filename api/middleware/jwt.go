package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vUdayKumarr/hotel-reservation/db"
)

func JWTAuthentication(userStore db.UserStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokens, ok := c.GetReqHeaders()["X-Api-Token"]
		if len(tokens) == 0 {
			return fmt.Errorf("unauthorized")
		}

		// Assuming you want to use the first token if multiple tokens are present
		token := tokens[0]

		if !ok {
			return fmt.Errorf("unauthorized")
		}
		claims, err := validateToken(token)
		if err != nil {
			return err
		}

		expiresFloat := claims["expires"].(float64)
		expires := int64(expiresFloat)
		// Check token expiration
		if time.Now().Unix() > expires {
			return fmt.Errorf("token expired")
		}
		userID := claims["id"].(string)
		user, err := userStore.GetUserByID(c.Context(), userID)
		if err != nil {
			return fmt.Errorf("unauthorized")
		}
		// Set the current authenticated user to the context
		c.Context().SetUserValue("user", user)
		return c.Next()
	}
}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, fmt.Errorf("unauthorized")
		}
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("failed to parse JWT token", err)
		return nil, fmt.Errorf("unauthorized")
	}
	if !token.Valid {
		fmt.Println("invalid token")
		return nil, fmt.Errorf("unauthorized")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("invalid claims")
		return nil, fmt.Errorf("unauthorized")
	}
	return claims, nil
}
