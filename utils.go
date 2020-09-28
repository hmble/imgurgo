package imgurgo

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func printBytes(body io.Reader, client *Client) {

	if client.Print {
		bodyBytes, err := ioutil.ReadAll(body)

		if err != nil {
			panic(err)
		}

		fmt.Println(string(bodyBytes))
	} else {
		fmt.Println(`client Print setting is not set to true, set it to true by using method
      client.SetPrint()
    `)
	}

}
func respError(path string) {
	log.Fatalf("Error in getting reponse from path : %s\n", path)
}

// HTTP                            Header Description
// ======================================================================
// X-RateLimit-UserLimit           Total credits that can be allocated.
// X-RateLimit-UserRemaining       Total credits available.
// X-RateLimit-UserReset           Timestamp (unix epoch) for when the credits will be reset.
// X-RateLimit-ClientLimit         Total credits that can be allocated for the application in a day.
// X-RateLimit-ClientRemaining     Total credits remaining for the application in a day.

func PrintHeader(resp *http.Response) {
	fmt.Println(resp.Status)

	fmt.Println("--------REQUEST HEADER--------")
	fmt.Printf("Ratelimit UserLimit       : %s\n", resp.Header.Get("X-RateLimit-UserLimit"))
	fmt.Printf("Ratelimit UserRemaining   : %s\n", resp.Header.Get("X-Ratelimit-UserRemaining"))
	fmt.Printf("Ratelimit UserReset       : %s\n", resp.Header.Get("X-Ratelimit-UserReset"))
	fmt.Printf("Ratelimit ClientLimit     : %s\n", resp.Header.Get("X-Ratelimit-ClientLimit"))
	fmt.Printf("Ratelimit ClientRemaining : %s\n", resp.Header.Get("X-Ratelimit-ClientRemaining"))
	fmt.Println("--------REQUEST HEADER--------")

}

func SaveResponse(r io.Reader, filepath string) {

	f, err := os.Create(filepath)

	if err != nil {
		log.Fatal("Error in creating file")
	}

	defer f.Close()

	written, copyErr := io.Copy(f, r)

	if copyErr != nil {
		log.Fatal("Error in writing bytes to file")
	}

	fmt.Printf("Wrote %d bytes at %s\n", written, filepath)
}
