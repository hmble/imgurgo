package imgurgo

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"golang.org/x/oauth2"
)

type User struct {
	Authenticator *Authenticator
}

const (
	BaseAuthURL = "https://api.imgur.com/3/"
)

// TODO(hmble): Add authenticated user info to client so that we can use it
// along different methods.
type Client struct {
	Http  *http.Client
	Token *oauth2.Token
	x     ratelimit

	common service // Reuse same struct instead of creating

	Print   bool
	Account *AccountService
	Image   *ImageService
}

type ratelimit struct {
	used      int
	remaining int
	reset     int
}
type service struct {
	client *Client
}

var NoAuthClient = &Client{
	Http: new(http.Client),
}

//var defaultAuth *internal.Authenticator = internal.DefaultClient

func (u *User) UserClient(token *oauth2.Token) *Client {
	c := &Client{
		Http:  u.Authenticator.Config.Client(oauth2.NoContext, token),
		Token: token,
	}

	c.Print = false
	c.common.client = c
	c.Account = (*AccountService)(&c.common)
	c.Image = (*ImageService)(&c.common)
	// here we can't use service struct because we included `path` member in
	// CommentService
	return c
}

func (u *User) Authenticate() (*oauth2.Token, error) {

	fmt.Println("Authentication starts from here:  ")
	fmt.Printf("Visit the url given below and paste the code given in url : \n %s", AuthUrl(true, u.Authenticator))

	fmt.Println("\n Enter the code here : ")

	var code string
	fmt.Scan(&code)

	token, err := GetToken(code, u.Authenticator)

	if err != nil {
		log.Fatal("Error in getting token")
		return nil, err
	}

	u.SaveToken("token.json", token)

	return token, nil

}

func (u *User) SaveToken(path string, token *oauth2.Token) {
	SaveToken(path, token)
}

func (u *User) UpdateToken(token *oauth2.Token) {
	UpdateToken(token, u.Authenticator)
}

// func TokenFromFile(filepath string) (*oauth2.Token, error) {
// 	return TokenFromFile(filepath)
// }

type Option map[string]string

var NoOptions Option = Option{}

func (c *Client) Get(endpoint string, opts Option) (res *http.Response, err error) {

	temp := BaseAuthURL + endpoint
	u, _ := url.Parse(temp)
	q, _ := url.ParseQuery(u.RawQuery)

	for k, v := range opts {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	path := u.String()

	// if c.x.remaining < 10 {
	// 	log.Fatal("---YOUR LIMIT HAS EXTENDED Wait for ", c.x.reset)
	// }

	req, err := http.NewRequest("GET", path, nil)

	if err != nil {
		log.Fatal("Error getting request")
		return nil, err
	}

	str := fmt.Sprintf("Bearer %s", c.Token.AccessToken)
	req.Header.Add("Authorization", str)

	return c.Http.Do(req)

}

type PostData map[string]string

var NoPostdata = PostData{}

func (c *Client) Post(endpoint string, postdata PostData) (*http.Response, error) {
	data := url.Values{}

	fullurl := BaseAuthURL + endpoint

	for k, v := range postdata {
		data.Set(k, v)
	}
	body := bytes.NewBufferString(data.Encode())
	req, err := http.NewRequest(http.MethodPost, fullurl, body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	return c.Http.Do(req)

}

// Reference
// https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/04.5.html
func (c *Client) PostImageUpload(endpoint string, postdata PostData, filename string) (*http.Response, error) {

	fullurl := BaseAuthURL + endpoint

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	for k, v := range postdata {
		bodyWriter.WriteField(k, v)
	}

	fileWriter, err := bodyWriter.CreateFormFile("image", filename)

	if err != nil {
		log.Fatal("error in writing to buffer from fileWriter:  ", err)

		return nil, err

	}

	fh, err := os.Open(filename)

	if err != nil {
		log.Fatal("Error in openeing file ", err)
	}
	defer fh.Close()

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return nil, err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	body := bytes.NewBufferString(bodyBuf.String())
	req, err := http.NewRequest(http.MethodPost, fullurl, body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	str := fmt.Sprintf("Bearer %s", c.Token.AccessToken)
	req.Header.Add("Authorization", str)

	return c.Http.Do(req)

}
func (c *Client) Delete(endpoint string, opts Option) (res *http.Response, err error) {

	temp := BaseAuthURL + endpoint
	u, _ := url.Parse(temp)
	q, _ := url.ParseQuery(u.RawQuery)

	for k, v := range opts {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	path := u.String()
	req, err := http.NewRequest("DELETE", path, nil)

	if err != nil {
		log.Fatal("Error getting request")
		return nil, err
	}

	str := fmt.Sprintf("Bearer %s", c.Token.AccessToken)
	req.Header.Add("Authorization", str)

	return c.Http.Do(req)

}

func (c *Client) Put(endpoint string, postdata PostData) (*http.Response, error) {
	fullurl := BaseAuthURL + endpoint

	data := url.Values{}

	data.Set("api_type", "json")

	for k, v := range postdata {
		data.Set(k, v)
	}
	body := bytes.NewBufferString(data.Encode())

	req, err := http.NewRequest(http.MethodPut, fullurl, body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	req.Header.Add("Content-Length", strconv.Itoa(len(data)))

	return c.Http.Do(req)

}

func (c *Client) savelimit(resp *http.Response) {
	used, errUser := strconv.Atoi(resp.Header.Get("X-Ratelimit-Used"))
	if errUser != nil {
		log.Fatal("Error in converting ratelimit used")
	}

	remaining, errRemaining := strconv.ParseFloat(resp.Header.Get("X-Ratelimit-Remaining"), 32)
	if errRemaining != nil {
		log.Fatal("Error in converting ratelimit remaining")
	}

	reset, errReset := strconv.Atoi(resp.Header.Get("X-Ratelimit-Reset"))
	if errReset != nil {
		log.Fatal("Error in converting ratelimit reset")
	}

	c.x = ratelimit{
		used:      used,
		remaining: int(remaining),
		reset:     reset,
	}
}

func (c *Client) SetPrint() {
	if !c.Print {
		c.Print = true
	}
}
