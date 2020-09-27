package imgurgo

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type AccountService service

func (a *AccountService) AccountBase() {
	path := "/account/me"

	resp, err := a.client.Get(path, NoOptions)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)

}

func (a *AccountService) AccountSettings() {
	path := "/account/me/settings"

	resp, err := a.client.Get(path, NoOptions)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	//printBytes(resp.Body, a.client)
	SaveResponse(resp.Body, "settings.json")

}

// Image upload

type ImageRespone struct {
	Id         string        `json:"id"`
	Link       string        `json:"link"`
	Width      int32         `json:"width"`
	Height     int32         `json:"height"`
	Size       uint64        `json:"size"`
	Datetime   time.Duration `json:"datetime"`
	Deletehash string        `json:"deletehash"`
}

func (a *AccountService) UplaodImage(imagename string, postdata PostData) {
	path := "/image"
	resp, err := a.client.PostImageUpload(path, postdata, imagename)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	var result struct {
		Data ImageRespone `json:"data"`
	}

	if jerr := json.NewDecoder(resp.Body).Decode(&result); jerr != nil {
		log.Fatal("Error in decoding json response, Err: ", err)

	}
	fmt.Println(result)

}
