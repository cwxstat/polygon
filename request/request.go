package request

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

func Request(url string) (string, error) {

	request := gorequest.New()
	resp, body, err := request.Get(url).End()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return body, fmt.Errorf("Error in request: %s", resp.Status)
	}
	return body, nil
}
