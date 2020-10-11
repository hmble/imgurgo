package imgurgo

import "fmt"

type AlbumService service

// Get information about a specific album.
// Model Reference : https://api.imgur.com/models/album
func (a *AlbumService) Info(id string) {
	path := fmt.Sprintf("/album/%s", id)

	resp, err := a.client.Get(path, NoOptions)

	if err != nil {
		respError(path)
	}
	defer resp.Body.Close()

	printBytes(resp.Body, a.client)

}

// Return all of the images in the album
// Model Reference: https://api.imgur.com/models/image
func (a *AlbumService) AllImages(id string) {
	path := fmt.Sprintf("/album/%s/images", id)

	resp, err := a.client.Get(path, NoOptions)

	if err != nil {
		respError(path)
	}
	defer resp.Body.Close()

	printBytes(resp.Body, a.client)

}

// Get information about an image in an album, any additional actions found in
// Image Endpoint will also work.
func (a *AlbumService) Image(albumId, imageId string) {
	path := fmt.Sprintf("/album/%s/image/%s", albumId, imageId)

	resp, err := a.client.Get(path, NoOptions)

	if err != nil {
		respError(path)
	}
	defer resp.Body.Close()

	printBytes(resp.Body, a.client)

}

// TODO(hmble): Add support for anonymous image upload.
// TODO(hmble): More better way to create postdata
// Key                                  Description
// ================================================================================
// ids[]             ( optional ) The image ids that you want to be included in the album.
// deletehashes[]    ( optional ) The deletehashes of the images that you want to be included in the album.
// title	           ( optional )  The title of the album
// description       ( optional ) The description of the album
// privacy           ( optional ) Sets the privacy level of the album. Values are : public | hidden | secret.
//                                Defaults to user's privacy settings for logged in users.
// layout            ( optional ) Sets the layout to display the album. Values are : blog | grid | horizontal | vertical
// cover             ( optional ) The ID of an image that you want to be the cover of the album

// Create a new album. Optional parameter of ids[] is an array of image ids to add
// to the album.
// array of image ids to add to the album. If uploading anonymous images to
// anonymous album please use the optional parameter of deletehashes[] rather than
// ids[]. Note: including the optional deletehashes[] parameter will also work for
// authenticated user albums. There is no need to duplicate image ids with their
// corresponding deletehash.
// This method is available without authenticating an account, and may be used
// merely by sending "Authorization: Client-ID {client_id}" in the request headers.
// Doing so will create an anonymous album which is not tied to an account.
func (a *AlbumService) CrateAlbum(postdata PostData) {
	path := "/album"

	resp, err := a.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

// Key                                  Description
// ================================================================================
// ids[]             ( optional ) The image ids that you want to be included in the album.
// deletehashes[]    ( optional ) The deletehashes of the images that you want to be included in the album.
// title	           ( optional )  The title of the album
// description       ( optional ) The description of the album
// privacy           ( optional ) Sets the privacy level of the album. Values are : public | hidden | secret.
//                                Defaults to user's privacy settings for logged in users.
// layout            ( optional ) Sets the layout to display the album. Values are : blog | grid | horizontal | vertical
// cover             ( optional ) The ID of an image that you want to be the cover of the album

// Update the information of an album. For anonymous albums, {album} should be the
// deletehash that is returned at creation. If uploading anonymous images to
// anonymous album please use the optional parameter of deletehashes[] rather than
// ids[]. Note: including the optional deletehashes[] parameter will also work for
// authenticated user albums. There is no need to duplicate image ids with their
// corresponding deletehash.
func (a *AlbumService) UpdateAlbum(deletehash string, postdata PostData) {
	path := fmt.Sprintf("/album/%s", deletehash)

	resp, err := a.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

// Delete an album with a given ID. You are required to be logged in as the user to
// delete the album. For anonymous albums, {album} should be the deletehash that is
// returned at creation.
func (a *AlbumService) DeleteAlbum(deletehash string) {
	path := fmt.Sprintf("/album/%s", deletehash)

	resp, err := a.client.Delete(path, NoOptions)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

// Favorite an album with a given ID. The user is required to be logged in to
// favorite the album.
func (a *AlbumService) FavoriteAlbum(albumid string) {
	path := fmt.Sprintf("/album/%s/favorite", albumid)

	resp, err := a.client.Post(path, NoPostdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

//
// Key                              Description
//======================================================================
// ids[]             ( optional ) The image ids that you want to be added to the album.
//                                The [] represents the ability to use this variable as an array.
// deletehashes[]    ( optional ) The image deletehashes that you want to be added to the album.
//                            The [] represents the ability to use this variable as an array.

// Sets the images for an album, removes all other images and only uses the
// images in this request. For anonymous albums, {album} should be the
// deletehash that is returned at creation. If uploading anonymous images to
// anonymous album please use the optional parameter of deletehashes[] rather
// than ids[]. Note: including the optional deletehashes[] parameter will also
// work for authenticated user albums. There is no need to duplicate image ids
// with their corresponding deletehash.
func (a *AlbumService) SetAlbumImages(deletehash string, postdata PostData) {
	path := fmt.Sprintf("/album/%s", deletehash)

	resp, err := a.client.Post(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

//
// Key                              Description
//======================================================================
// ids[]             ( optional ) The image ids that you want to be added to the album.
//                                The [] represents the ability to use this variable as an array.
// deletehashes[]    ( optional ) The image deletehashes that you want to be added to the album.
//                            The [] represents the ability to use this variable as an array.

// NOTE : we can only add one id at a time
// Reference: https://apidocs.imgur.com/#b98029b6-5cc1-4a6f-b4bf-fe1db50869a2

// For anonymous albums, {album} should be the deletehash that is returned at
// creation. If uploading anonymous images to anonymous album please use the
// optional parameter of deletehashes[] rather than ids[]. Note: including the
// optional deletehashes[] parameter will also work for authenticated user
// albums. There is no need to duplicate image ids with their corresponding
// deletehash.
func (a *AlbumService) AddImagesToAlbum(deletehash string, postdata PostData) {
	path := fmt.Sprintf("/album/%s/add", deletehash)

	resp, err := a.client.Put(path, postdata)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}

//
// Key                              Description
//======================================================================
// ids[]             ( required ) The image ids that you want to be removed from the album.
//                                The [] represents the ability to use this variable as an array.

// Takes parameter, ids[], as an array of ids and removes from the album. For
// anonymous albums, {album} should be the deletehash that is returned at
// creation.
func (a *AlbumService) RemoveImages(deletehash string, opts Option) {
	path := fmt.Sprintf("/album/%s/remove_images", deletehash)

	resp, err := a.client.Delete(path, opts)

	if err != nil {
		respError(path)
	}

	defer resp.Body.Close()

	printBytes(resp.Body, a.client)
}
