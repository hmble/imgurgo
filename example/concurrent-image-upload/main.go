package main

import (
	"log"
	"os"
	"path/filepath"
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

	c.SetPrint()
	type fileinfo struct {
		path string
		name string
	}
	var wg sync.WaitGroup
	var files []fileinfo
	err := filepath.Walk("data", func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() {

			files = append(files, fileinfo{path: path, name: info.Name()})
		}
		return nil

	})

	if err != nil {
		panic(err)
	}

	wg.Add(len(files))
	for _, file := range files {
		go func(f fileinfo, wg *sync.WaitGroup) {

			defer wg.Done()

			c.Image.UplaodImage(f.path, imgurgo.PostData{
				"title": f.name,
			})
		}(file, &wg)

	}

	wg.Wait()

}
