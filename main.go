package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// define command line parameters.
var phoneNumbers = flag.String("phoneNumbers", "", "Phone Numbers")
var messageID = flag.String("messageID", "", "Message ID")
var messageContent = flag.String("messageContent", "", "Message Content")
var senderID = flag.String("senderID", "", "Sender ID")

func main() {
	//load env file and read command line arguments.
	godotenv.Load()
	flag.Parse()

	// assign values from env
	key := os.Getenv("API_KEY")
	urlBase := os.Getenv("URL_BASE")
	smsExt := os.Getenv("SMS_EXT")
	callUrl := fmt.Sprintf("%s%s", urlBase, smsExt)

	//split phone numbers into array
	phoneNos := strings.Split(*phoneNumbers, ",")

	//create token from key
	token, tokenCreated, err := createJWT(key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(token)
	fmt.Println(tokenCreated)
	fmt.Println(*messageContent)

	for _, phoneNo := range phoneNos {
		if tokenCreated < (time.Now().Unix() - 25) {
			token, tokenCreated, err = createJWT(key)
			if err != nil {
				log.Fatal(err)
			}
		}
		body := fmt.Sprintf(`{"phone_number":"%s",template_id":"%s","personalisation": %s,"sms_sender_id": "%s"}`, phoneNo, *messageID, *messageContent, *senderID)
		fmt.Println(callUrl)
		fmt.Println(body)
	}
}

func createJWT(key string) (string, int64, error) {
	//Split Iss and secret out of key.
	splitKey := strings.Split(key, "-")
	splitLen := len(splitKey)
	iss := fmt.Sprintf("%s-%s-%s-%s-%s", splitKey[splitLen-10], splitKey[splitLen-9], splitKey[splitLen-8], splitKey[splitLen-7], splitKey[splitLen-6])
	secret := fmt.Sprintf("%s-%s-%s-%s-%s", splitKey[splitLen-5], splitKey[splitLen-4], splitKey[splitLen-3], splitKey[splitLen-2], splitKey[splitLen-1])

	//get time stamp (now!)
	epoch := time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": iss,
		"iat": epoch,
	})

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", epoch, err
	}

	return signedToken, epoch, nil
}
