package apihealth

import (
	"io/ioutil"
	"github.com/parnurzeal/gorequest"
	"fmt"
)

type HttpResponse struct {
	HttpStatus string
	Body       []byte
}

type HttpCaller struct{}

func (h HttpCaller) Get(url string, ch chan<- HttpResponse) {
	httpResponse := doGet(url)
	httpBody, _ := ioutil.ReadAll(httpResponse.Body)
	ch <- HttpResponse{httpResponse.Status, httpBody}
}

func doGet(url string) gorequest.Response {

	req := gorequest.New()

	resp, _, errors := req.Get(url).End()

	if errors != nil {
		fmt.Printf("Error calling Health-API %s", url)
	}

	return resp
}
