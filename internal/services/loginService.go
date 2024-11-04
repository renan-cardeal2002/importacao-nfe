package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("my_secret_key")

func createToken(id int) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token expira em 1 hora

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Método de assinatura não suportado")
		}
		return jwtKey, nil
	})

	return token, err
}

func implementacao() {
	tokenString, err := createToken(1)
	if err != nil {
		fmt.Println("Erro ao criar token:", err)
		return
	}
	fmt.Println("Token JWT:", tokenString)

	token, err := parseToken(tokenString)
	if err != nil {
		fmt.Println("Erro ao verificar token:", err)
		return
	}

	if token.Valid {
		fmt.Println("Token válido")
		claims, _ := token.Claims.(jwt.MapClaims)
		fmt.Printf("ID: %v\n", claims["id"])
		fmt.Printf("Expiração: %v\n", time.Unix(int64(claims["exp"].(float64)), 0))
	} else {
		fmt.Println("Token inválido")
	}
}
