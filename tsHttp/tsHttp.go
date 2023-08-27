package tsHttp

import (
	"io"
	"log"
	"net/http"
)

func SendHttpRequest(method string, url string, param []string) (sitemap []byte, err error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	sitemap, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}
