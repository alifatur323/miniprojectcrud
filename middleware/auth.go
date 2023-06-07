package middleware

import (
	"crmservice/dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

func CreateToken(sub uint, name string) string {
	// Inisialisasi klaim-klaim yang ingin Anda sertakan dalam token
	claims := jwt.MapClaims{
		"sub":  sub,
		"name": name,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}
	// Tandatangani token dengan kunci rahasia
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("Kmzwa8awaa"))
	if err != nil {
		return "Failed"
	}
	return signedToken
}
func AuthMiddleware(c *gin.Context) {
	// Ambil header Authorization
	receivedToken := c.GetHeader("Authorization")

	token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("Kmzwa8awaa"), nil
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.TokenResponse())
		c.Abort()
		return
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok {
		c.Next()
		return
	}
	c.JSON(http.StatusBadRequest, dto.TokenResponse())
	c.Abort()
	return
}
