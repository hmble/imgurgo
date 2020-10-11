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
	u := &imgurgo.User{
		Authenticator: authenticator,
	}

	token, er := imgurgo.TokenFromFile("token.json")
	if er != nil {
		log.Fatal("Error in reading token")
	}

	u.UpdateToken(token)

	c := u.UserClient(token)

	images := c.Account.Images(0)

	for _, image := range images {
		fmt.Printf("ID: %s\n", image.ID)
	}

}
