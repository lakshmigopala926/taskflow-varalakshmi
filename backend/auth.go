package main

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
)

var secret = []byte("secret123")

func Login(c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": "demo",
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	t, _ := token.SignedString(secret)

	c.JSON(http.StatusOK, gin.H{"token": t})
}

