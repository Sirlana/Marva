package main

import (
	"fmt"
	"time"

	"sirlana.com/sirlana/sso/libs"
)

func mains() {
	key := "sir"
	jwt := libs.NewJWT(key)
	jwt.AddExpiredDate(1)
	jwt.AddDataString("email", "cs@srilana.com")
	jwt.AddDataInt("id", 1)

	token, err := jwt.Encode()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(token)

	_, claims, _ := jwt.Decode(token)
	fmt.Println(claims["exp"])

	time.Sleep(2 * time.Second)
	if jwt.IsExpired(claims["exp"].(float64)) {
		fmt.Println("Valid")
	} else {
		fmt.Println("Expired ")
	}
}
