package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var key = flag.String("key", "", "Key Value")
var phoneNumbers = flag.String("phoneNumbers", "", "Phone Numbers")
var messageID = flag.String("messageID", "", "Message ID")
var messageContent = flag.String("messageContent", "", "Message Content")

func main() {
	flag.Parse()
	fmt.Println(*key)
	fmt.Println(*phoneNumbers)
	fmt.Println(*messageID)
	fmt.Println(*messageContent)

	splitKey := strings.Split(*key, "-")
	splitLen := len(splitKey)
	iss := fmt.Sprintf("%s-%s-%s-%s-%s", splitKey[splitLen-10], splitKey[splitLen-9], splitKey[splitLen-8], splitKey[splitLen-7], splitKey[splitLen-6])
	secret := fmt.Sprintf("%s-%s-%s-%s-%s", splitKey[splitLen-5], splitKey[splitLen-4], splitKey[splitLen-3], splitKey[splitLen-2], splitKey[splitLen-1])
	fmt.Println("creating jwt")
	token, err := createJWT(iss, secret)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)

}

func createJWT(iss, secretKey string) (string, error) {

	epoch := time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": iss,
		"iat": epoch,
	})

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
