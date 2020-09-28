package imgurgo

import (
	"fmt"
	"strconv"
)

type AccountService service

// Request standard user information.
func (a *AccountService) AccountBase() {
	path := "/account/me"

	resp, err := a.client.Get(path, NoOptions)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)

}

// Returns the account settings, only accessible if you're logged in as the
// user.
func (a *AccountService) Settings() {
	path := "/account/me/settings"

	resp, err := a.client.Get(path, NoOptions)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	//printBytes(resp.Body, a.client)
	SaveResponse(resp.Body, "settings.json")

}

// Return the images the user has favorited in the gallery.
//
// Key                  Description
//=================================================================
// page          (optional) integer - allows you to set the page number so you don't have to
//               retrieve all the data at once.
// sort          (optional) 'oldest', or 'newest'. Defaults
//               to 'newest'.
func (a *AccountService) GalleryFavourites(sort string, page int) {
	path := "/account/me/gallery_favourites"
	opts := Option{
		"sort": sort,
		"page": strconv.Itoa(page),
	}

	resp, err := a.client.Get(path, opts)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	PrintHeader(resp)

	printBytes(resp.Body, a.client)
}

// Returns the users favorited images, only accessible if you're logged in as the
// user.
//
// Key                  Description
//=================================================================
// page          (optional) integer - allows you to set the page number so you don't have to
//               retrieve all the data at once.
// sort          (optional) 'oldest', or 'newest'. Defaults
//               to 'newest'.

func (a *AccountService) FavouriteImages(sort string, page int) {
	path := "/account/me/favorites"
	opts := Option{
		"sort": sort,
		"page": strconv.Itoa(page),
	}

	resp, err := a.client.Get(path, opts)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	PrintHeader(resp)

	printBytes(resp.Body, a.client)
}

// Return the images a user has submitted to the gallery
func (a *AccountService) Submissions(page string) {
	path := fmt.Sprintf("/account/me/submissions/%s", page)

	resp, err := a.client.Get(path, NoOptions)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

// Updates the account settings for a given user, the user must be logged in.
//
// Key                                  Description
// =======================================================================================
// bio                       (optional) The biography of the user, is displayed in the gallery profile page.
// public_images             (optional) Set the users images to private or public by default
// messaging_enabled         (optional) true | false - Allows the user to enable / disable private messages
// album_privacy             (optional) public | hidden | secret - Sets the default privacy level of albums the users creates
// accepted_gallery_terms    (optional) true | false - The user agreement to the Imgur Gallery terms.
// username                  (optional) A valid Imgur username (between 4 and 63 alphanumeric characters)
// show_mature               (optional) true | false - Toggle display of mature images in gallery list endpoints.
// newsletter_subscribed     (optional) true | false - Toggle subscription to email newsletter.

func (a *AccountService) ChangeSettings(postdata PostData) {
	path := "/account/me/settings"

	resp, err := a.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)

}

// Returns the totals for the gallery profile.
func (a *AccountService) GalleryProfile() {
	path := "/account/me/gallery_profile"

	resp, err := a.client.Get(path, NoOptions)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

// Checks to see if user has verified their email address

func (a *AccountService) VerifyEmail() {
	path := "/account/me/verifyemail"

	resp, err := a.client.Get(path, NoOptions)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

// Sends an email to the user to verify that their email is valid to upload to
// gallery. Must be logged in as the user to send.

func (a *AccountService) SendVerificationEmail() {
	path := "/account/me/verifyemail"

	resp, err := a.client.Post(path, NoPostdata)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

// Get all the albums associated with the account. Must be logged in as the user to
// see secret and hidden albums.

func (a *AccountService) Albums(page string) {
	path := fmt.Sprintf("/account/me/albums/%s", page)

	resp, err := a.client.Get(path, NoOptions)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

// Get additional information about an album, this endpoint works the same as the
// Album Endpoint. You can also use any of the additional routes that are used on
// an album in the album endpoint.
func (a *AccountService) AlbumInfo(id string) {
	path := fmt.Sprintf("/account/me/album/%s", id)

	resp, err := a.client.Get(path, NoOptions)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

// Return an array of all of the album IDs.
func (a *AccountService) AlbumIds(page int) {
	path := "/account/me/albums/ids"

	opts := Option{
		"page": strconv.Itoa(page),
	}
	resp, err := a.client.Get(path, opts)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

// Return the total number of albums associated with the account.
func (a *AccountService) AlbumCount() {
	path := "/account/me/albums/count"
	resp, err := a.client.Get(path, NoOptions)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

// Delete an Album with a given id.

func (a *AccountService) AlbumDeletion(id string) {
	path := fmt.Sprintf("/account/me/album/%s", id)
	resp, err := a.client.Delete(path, NoOptions)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

// Return the comments the user has created.
//
// Key                  Description
//=================================================================
// page          (optional) integer - allows you to set the page number so you don't have to
//               retrieve all the data at once.
// sort          (optional) 'oldest', or 'newest'. Defaults
//               to 'newest'.
func (a *AccountService) Comments(sort string, page int) {
	path := "/account/me/comments"
	opts := Option{
		"sort": sort,
		"page": strconv.Itoa(page),
	}
	resp, err := a.client.Get(path, opts)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

// Return information about a specific comment. This endpoint works the same as the
// Comment Endpoint. You can use any of the additional actions that the comment
// endpoint allows on this end point.

func (a *AccountService) CommentInfo(id string) {
	path := fmt.Sprintf("/account/me/comment/%s", id)

	resp, err := a.client.Get(path, NoOptions)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

// Return an array of all of the comment IDs.
func (a *AccountService) CommentIds(sort string, page int) {
	path := "/account/me/comments/ids"

	opts := Option{
		"sort": sort,
		"page": strconv.Itoa(page),
	}
	resp, err := a.client.Get(path, opts)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

// Return a count of all of the comments associated with the account.
func (a *AccountService) CommentsCount() {
	path := "/account/me/comments/count"
	resp, err := a.client.Get(path, NoOptions)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

// Delete an Album with a given id.

func (a *AccountService) CommentDeletion(id string) {
	path := fmt.Sprintf("/account/me/comment/%s", id)
	resp, err := a.client.Delete(path, NoOptions)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

//Return all of the images associated with the account. You can page through the
//images by setting the page, this defaults to 0.

func (a *AccountService) Images(page int) {
	path := fmt.Sprintf("/account/me/images/%d", page)
	resp, err := a.client.Delete(path, NoOptions)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

// Return information about a specific image. This endpoint works the same as the
// Image Endpoint. You can use any of the additional actions that the image
// endpoint with this endpoint.

func (a *AccountService) ImageInfo(id string) {
	path := fmt.Sprintf("/account/me/image/%s", id)

	resp, err := a.client.Get(path, NoOptions)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

// Returns an array of Image IDs that are associated with the account.

func (a *AccountService) ImageIds(page int) {

	path := "/account/me/images/ids"

	opts := Option{
		"page": strconv.Itoa(page),
	}

	resp, err := a.client.Get(path, opts)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

// Returns the total number of images associated with the account.
func (a *AccountService) ImageCount() {
	path := "/account/me/images/count"
	resp, err := a.client.Get(path, NoOptions)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

// Deletes an Image. This requires a delete hash rather than an ID.
func (a *AccountService) ImageDeletion(deleteHash string) {
	path := fmt.Sprintf("/account/me/image/%s", deleteHash)
	resp, err := a.client.Delete(path, NoOptions)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

// Returns all of the reply notifications for the user. Required to be logged in as that user
// Key         Description
// =====================================
// new     optional boolean - false for all notifications, true for only
//         non-viewed notification. Default is true.
func (a *AccountService) Replies(newreplies bool) {
	path := "/account/me/notifications/replies"

	opts := Option{
		"new": strconv.FormatBool(newreplies),
	}
	resp, err := a.client.Get(path, opts)

	if err != nil {
		respError(path)

	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}
