package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type user struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type myClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(mysql.Open("root:((root))@tcp(localhost:3306)/store"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to databases")
	}
	var users user
	json.NewDecoder(r.Body).Decode(&users)
	var result user
	res := db.Where("username = ?", users.Username).Where("password = ?", users.Password).First(&result)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Login Failed"))
		} else {
			w.Write([]byte(res.Error.Error()))
		}
		return
	}
	claims := myClaims{
		Username: result.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("Lindu"))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println("token :", signedToken)
	w.Write([]byte(signedToken))
}
