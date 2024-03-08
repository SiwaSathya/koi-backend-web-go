package middleware

import (
	"fmt"
	"koi-backend-web-go/domain"
	"koi-backend-web-go/utils/fiberutil"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

const SecretKey = "your-secret-key"

func CreateToken(fill *domain.TokenClaims) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Set expiration time

	claims := &domain.TokenClaims{
		User: fill.User,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(c *fiber.Ctx) error {
	authHeaders := c.Get("Authorization")
	if !strings.Contains(authHeaders, "Bearer") {
		return fiberutil.ReturnStatusUnauthorized(c)
	}

	tokens := strings.Replace(authHeaders, "Bearer ", "", -1)
	if tokens == "Bearer" {
		return fiberutil.ReturnStatusUnauthorized(c)
	}

	// SecretKey adalah kunci rahasia yang sama yang digunakan untuk menandatangani token
	secretKey := []byte("yourSecretKey")

	// Memverifikasi token
	resp, err := verifyToken(tokens, secretKey)
	fmt.Println(resp)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "Error",
			"message": "Token tidak valid",
			"error":   err.Error(),
		})
	} else {
		c.Locals("affiliate_id", resp.ID)
		return c.Next()
	}
}

func verifyToken(tokenString string, secretKey []byte) (*domain.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate that the token is signed with the correct method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return &domain.User{}, err
	}

	// Check if the token is valid
	if !token.Valid {
		return &domain.User{}, fmt.Errorf("invalid token")
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return &domain.User{}, fmt.Errorf("invalid claims")
	}

	// Check if the token has an expiration claim
	expiration, ok := claims["exp"].(float64)
	if !ok {
		return &domain.User{}, fmt.Errorf("expiration claim not found")
	}

	// Convert the expiration claim to a time.Time
	expirationTime := time.Unix(int64(expiration), 0)

	// Check if the token has expired
	if time.Now().After(expirationTime) {
		return &domain.User{}, fmt.Errorf("token has expired")
	}

	// Extract user ID from the payload
	_, ok = claims["user_id"].(string)
	if !ok {
		return &domain.User{}, fmt.Errorf("user id not found in claims")
	}

	// repoAf := postgresql.NewPostgreAffiliate(db.Postgres.DB)

	// resp, errs := repoAf.GetBySSOID(userID)
	// if errs != nil {
	// 	return &domain.User{}, fmt.Errorf("Failed to get affiliate data: %v", errs)
	// }

	return &domain.User{}, nil
}
