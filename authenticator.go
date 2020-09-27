package imgurgo

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://api.imgur.com/oauth2/authorize",
	TokenURL: "https://api.imgur.com/oauth2/token",

	AuthStyle: oauth2.AuthStyleInParams,
}

// Returns Config file of oauth2.Config

// var reddit_config *oauth2.Config = &oauth2.Config{

// 	ClientID:     os.Getenv("CLIENT_ID"),
// 	ClientSecret: os.Getenv("CLIENT_SECRET"),
// 	Scopes:       Scopes,
// 	Endpoint:     Endpoint,
// 	RedirectURL:  "https://example.com/auth",
// }

type Authenticator struct {
	Config    *oauth2.Config
	Useragent string
}

// TODO: For now we use all scopes to while authentication
// We need a method to dynamically set scopes and also token permanent
// value
// var DefaultClient *Authenticator = &Authenticator{
// 	Config:    reddit_config,
// 	Useragent: os.Getenv("USER_AGENT"),
// }

type uaSetterTransport struct {
	config    *oauth2.Config
	useragent string
}

func SaveToken(path string, token *oauth2.Token) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)

}

func AuthUrl(isPermanent bool, authenticator *Authenticator) string {
	codeParam := oauth2.SetAuthURLParam("response_type", "code")
	url := authenticator.Config.AuthCodeURL(uuid.New().String(), codeParam)
	return url

}

// basic Authourization with username and password
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))

}

func (t *uaSetterTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", t.useragent)
	// set a non-standard Authorization header because reddit demands it
	// https://github.com/reddit/reddit/wiki/OAuth2#retrieving-the-access-token
	req.Header.Set("Authorization", basicAuth(t.config.ClientID, t.config.ClientSecret))
	return http.DefaultTransport.RoundTrip(req)
}

func TokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func GetToken(code string, authenticator *Authenticator) (*oauth2.Token, error) {
	client := &http.Client{
		Transport: &oauth2.Transport{
			Source: authenticator.Config.TokenSource(oauth2.NoContext, &oauth2.Token{
				AccessToken: code,
			}),
			Base: &uaSetterTransport{
				config:    authenticator.Config,
				useragent: authenticator.Useragent,
			},
		},
	}
	ctx := context.WithValue(oauth2.NoContext, oauth2.HTTPClient, client)

	return authenticator.Config.Exchange(ctx, code)
}

type rfToken struct {
	config    *oauth2.Config
	useragent string
}

func (r *rfToken) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", r.useragent)
	req.Header.Set("Authorization", basicAuth(r.config.ClientID, r.config.ClientSecret))
	return http.DefaultTransport.RoundTrip(req)
}

// parameter token is token from cached file
// used for refresh token
func UpdateToken(token *oauth2.Token, authenticator *Authenticator) {

	client := &http.Client{
		Transport: &rfToken{
			config:    authenticator.Config,
			useragent: authenticator.Useragent,
		},
	}
	ctx := context.WithValue(oauth2.NoContext, oauth2.HTTPClient, client)

	if !token.Valid() {
		newtoken, err := authenticator.Config.TokenSource(ctx, &oauth2.Token{
			RefreshToken: token.RefreshToken,
		}).Token()

		if err != nil {
			log.Fatal("Error from UpdateToken in getting refresh token\n", err)
		}

		token.AccessToken = newtoken.AccessToken

		//		fmt.Println("Expiry time : ", newtoken.Expiry)

		token.Expiry = newtoken.Expiry

		SaveToken("token.json", token)

	}
}
