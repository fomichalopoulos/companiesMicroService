package swagger

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/fomichalopoulos/companiesMicroService/models"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret")

func createToken(user models.User) (string, error) {
	/*
		token := jwt.NewWithClaims(jwt.SigningMethodHS256,
			jwt.MapClaims{
				"username": user.Username,
				"exp":      time.Now().Add(time.Hour * 24).Unix(),
			})
	*/

	// Define the expiration time for the token (e.g., 15 minutes from now)
	expirationTime := time.Now().Add(15 * time.Minute)

	// Create the JWT claims, which includes the subject (user ID), issuer, and expiration time
	claims := &jwt.RegisteredClaims{
		Subject:   user.Username,                      // User ID (subject of the token)
		ExpiresAt: jwt.NewNumericDate(expirationTime), // Expiration time
		IssuedAt:  jwt.NewNumericDate(time.Now()),     // Issued at
		Issuer:    "copmanyAPI",                       // Issuer (who issued the token)
	}

	// Create a new JWT token with claims, and sign it using the HMAC-SHA256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) error {
	claims := &jwt.RegisteredClaims{}
	/*
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
	*/

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Provide the secret key to validate the signature
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
		return errors.New("Token has expired")
	}

	if claims.Subject != "Fotis" {
		fmt.Println("sub not Fotis")
		return errors.New("User is not allowed here")
	}

	if !token.Valid {
		//return fmt.Errorf("invalid token")
		return errors.New("invalid token")
	}

	return nil
}

func checkJWT(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		//w.WriteHeader(http.StatusUnauthorized)
		//fmt.Fprint(w, "Missing authorization header")
		return errors.New("Missing authorization header")
	}

	// Ensure the token has the "Bearer" prefix
	if !strings.HasPrefix(tokenString, "Bearer ") {
		fmt.Println("No prefix!!!")
		return errors.New("Invalid token format: expected 'Bearer <token>'")
	}

	tokenString = tokenString[len("Bearer "):]

	err := verifyToken(tokenString)
	if err != nil {
		//w.WriteHeader(http.StatusUnauthorized)
		//fmt.Fprint(w, "Invalid token")
		fmt.Println(err)
		return errors.New("Invalid token")
	}

	fmt.Println("Welcome to the protected area")
	return nil
}
