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
func PrintHeader(resp *http.Response) {
	fmt.Println(resp.Status)

	fmt.Println("--------REQUEST HEADER--------")
	fmt.Printf("Ratelimit Used : %s\n", resp.Header.Get("X-Ratelimit-Used"))
	fmt.Printf("Ratelimit Remaining : %s\n", resp.Header.Get("X-Ratelimit-Remaining"))
	fmt.Printf("Ratelimit Reset : %s\n", resp.Header.Get("X-Ratelimit-Reset"))
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
