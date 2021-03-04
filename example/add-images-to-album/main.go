package main

import (
	"fmt"
	"log"
	"os"
	"sync"

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

	token, er := imgurgo.TokenFromFile("imgurtoken.json")
	if er != nil {
		log.Fatal("Error in reading token")
	}

	u.UpdateToken(token)

	c := u.UserClient(token)

	//{"data":{"id":"GUNHu2v","deletehash":"S8DceXQW8gTN8RV"},"success":true,"status":200}

	var wg sync.WaitGroup
	ids := []string{"mQw8M6Y", "VMabdhR", "SwKJHHL"}
	wg.Add(len(ids))
	for _, id := range ids {
		go func(id string) {
			defer wg.Done()
			fmt.Println("For id : ", id)
			c.Album.AddImagesToAlbum("GUNHu2v", imgurgo.PostData{
				"ids[]": id,
			})
		}(id)
	}

	wg.Wait()
}
