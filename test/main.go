package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hmble/imgurgo"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/oauth2"
)

func main() {

	authenticator := &imgurgo.Authenticator{
		Config: &oauth2.Config{
			ClientID:     os.Getenv("CLIENT_ID"),
			ClientSecret: os.Getenv("CLIENT_SECRET"),
			Endpoint:     imgurgo.Endpoint,
			RedirectURL:  os.Getenv("REDIRECT_URL"),
		},
		Useragent: os.Getenv("USER_AGENT"),
	}
	imgurUser := &imgurgo.User{
		Authenticator: authenticator,
	}

	token, er := imgurgo.TokenFromFile("imgurtoken.json")
	if er != nil {
		log.Fatal("Error in reading token")
	}

	imgurUser.UpdateToken(token)

	imgurClient := imgurUser.UserClient(token)

resp := imgurClient.Image.UplaodImage("./BT0BWyTFuWd.jpg", imgurgo.PostData{
		"title": "BT0BWyTFuWd",
	})

	fmt.Println("ID is ", resp.Id)

	//mJSJkQY
}
