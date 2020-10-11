package imgurgo

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type ImageService service

// Image upload

type Image struct {
	ID          string        `json:"id"`
	Link        string        `json:"link"`
	Width       int32         `json:"width"`
	Height      int32         `json:"height"`
	Size        uint64        `json:"size"`
	Datetime    time.Duration `json:"datetime"`
	Deletehash  string        `json:"deletehash"`
	Views       int64         `json:"views"`
	NSFW        bool          `json:"nsfw"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
}

// Get information about an image.
func (i *ImageService) Info(id string) {
	path := fmt.Sprintf("/image/%s", id)

	resp, err := i.client.Get(path, NoOptions)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, i.client)
}

// Upload a new image
//
// Key                      Description
//==========================================================================================
// image       ( required ) A binary file, base64 data, or a URL for an image. (up to 10MB)
// album       ( optional ) The id of the album you want to add the image to. For anonymous albums
//                          , {album} should be the deletehash that is returned at creation.
// type        ( optional ) The type of the file that's being sent; file, base64 or URL
// name        ( optional ) The name of the file, this is automatically detected if
//                          uploading a file with a POST and multipart / form-data
// title       ( optional ) The title of the image.
// description ( optional ) The description of the image.

func (i *ImageService) UplaodImage(imagepath string, postdata PostData) {
	path := "/image"
	resp, err := i.client.PostImageUpload(path, postdata, imagepath)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	var result struct {
		Data Image `json:"data"`
	}

	if jerr := json.NewDecoder(resp.Body).Decode(&result); jerr != nil {
		log.Fatal("Error in decoding json response, Err: ", err)

	}

	fmt.Println(resp.Status)

}

// Deletes an image. For an anonymous image, {id} must be the image's deletehash.
// If the image belongs to your account then passing the ID of the image is
// sufficient.
func (i *ImageService) Delete(deletehash string) {
	path := fmt.Sprintf("/image/%s", deletehash)

	resp, err := i.client.Delete(path, NoOptions)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, i.client)
}

// Updates the title or description of an image. You can only update an image
// you own and is associated with your account. For an anonymous image, {id}
// must be the image's deletehash.
// Key              Description
//==================================================
// title       ( optional ) The title of the image.
// description ( optional ) The description of the image.

func (i *ImageService) UpdateInfo(id string, postdata PostData) {
	path := fmt.Sprintf("/image/%s", id)

	resp, err := i.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, i.client)
}

// Favorite an image with the given ID. The user is required to be logged in to
// favorite the image.
func (i *ImageService) FavoriteImage(id string) {
	path := fmt.Sprintf("/image/%s/favorite", id)

	resp, err := i.client.Post(path, NoPostdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, i.client)
}
