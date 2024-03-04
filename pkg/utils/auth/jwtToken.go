package auth

import (
	"gostart/pkg/config"
	authSchema "gostart/pkg/schema/auth"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id string) (authSchema.Token, error) {

	accessSecret := config.GetString("ACCESS_TOKEN_SECRET")
	refreshSecret := config.GetString("REFRESH_TOKEN_SECRET")

	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Minute * time.Duration(10)).Unix(),
	}

	accessToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(accessSecret))

	claims["exp"] = time.Now().Add(time.Hour * 24 * 90).Unix()
	refresToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(refreshSecret))

	return authSchema.Token{
		AccessToken:  accessToken,
		RefreshToken: refresToken,
	}, nil
}

func GetTokenID(c *fiber.Ctx) string {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}

func GetToken(c *fiber.Ctx) string {
	token := c.Locals("user").(*jwt.Token)
	return token.Raw
}
