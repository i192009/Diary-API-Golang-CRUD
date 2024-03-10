package helper

import (
	"diaryApi/model"
	"fmt"
	//"go/token"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)



func GenerateJWT(user model.User) (string , error){
	var privateKey = os.Getenv("JWT_PRIVATE_KEY")

	fmt.Println("3key", privateKey)
	tokenTTL , _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	fmt.Println("3 :" , tokenTTL)
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"id" : user.ID, 
		"iat" : time.Now().Unix(),
		"eat" : time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	fmt.Println("3:" , token)
	tokenString , err := token.SignedString(privateKey)
	if err != nil {
		fmt.Println("3e",err)
	} else {
		fmt.Println("3",tokenString)
	}
	return token.SignedString(privateKey)
}