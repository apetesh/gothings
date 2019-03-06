package greeter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/apetesh/gothings/greeter/api"
	"io/ioutil"
	"log"
	"net/http"
)

type HttpClient interface {
	Do(r *http.Request) (*http.Response, error)
}

type Client struct {
	ApiURL string
	client HttpClient
}

func NewClient(apiUrl string) *Client {
	return &Client{apiUrl, &http.Client{}}
}

func (c *Client) Greet(name string) (*api.GreetResponse, error) {
	greetRequest := &api.GreetRequest{Name: name}
	greetResponse := &api.GreetResponse{}
	err := c.send(http.MethodPost, "greet", greetRequest, greetResponse, 200)
	if err != nil {
		return nil, err
	}
	return greetResponse, err
}


func (c *Client) send(method, requestPath string, requestBody, responseBody interface{}, expectedStatus int) error {
	requestURL := fmt.Sprintf("%s/%s", c.ApiURL, requestPath)
	parsedRequestBody, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}
	request, err := http.NewRequest(method, requestURL, bytes.NewBuffer(parsedRequestBody))
	if err != nil {
		return fmt.Errorf("error creating request object. error:  %s", err)
	}
	request.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(request)
	if err != nil {
		return err
	}
	if resp.StatusCode != expectedStatus {
		return fmt.Errorf("server responded with status %d", resp.StatusCode)
	}
	b, err := ioutil.ReadAll(resp.Body)
	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Print("error closing response body")
		}
	}()
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, responseBody)
	if err != nil {
		return err
	}
	return nil
}
