package herokuas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	authToken  string
	httpClient *http.Client
}

func NewClient(token string) *Client {
	return &Client{
		authToken:  token,
		httpClient: &http.Client{},
	}
}

type Trigger struct {
	UUID          string `json:"uuid"`
	Name          string `json:"name"`
	State         string `json:"state"`
	Dyno          string `json:"dyno"`
	FrequencyType string `json:"frequencyType"`
	Schedule      string `json:"schedule"`
	Timezone      string `json:"timezone"`
	Value         string `json:"value"`
	Timeout       int    `json:"timeout,omitempty"`
}

type TriggerResponse struct {
	Message string  `json:"message"`
	Code    int     `json:"code"`
	Trigger Trigger `json:"trigger"`
}

func (c *Client) GetAll() (*map[string]Trigger, error) {
	body, err := c.httpRequest("triggers", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	triggers := map[string]Trigger{}
	err = json.NewDecoder(body).Decode(&triggers)
	if err != nil {
		return nil, err
	}
	return &triggers, nil
}

func (c *Client) GetTrigger(uuid string) (*Trigger, error) {
	body, err := c.httpRequest(fmt.Sprintf("triggers/%v", uuid), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	triggerResponse := &TriggerResponse{}
	err = json.NewDecoder(body).Decode(triggerResponse)
	if err != nil {
		return nil, err
	}
	return &triggerResponse.Trigger, nil
}

func (c *Client) NewTrigger(newTrigger *Trigger) (*Trigger, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(newTrigger)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("triggers", "POST", buf)
	if err != nil {
		return nil, err
	}

	triggerResponse := &TriggerResponse{}
	err = json.NewDecoder(body).Decode(triggerResponse)
	if err != nil {
		return nil, err
	}
	return &triggerResponse.Trigger, nil
}

func (c *Client) UpdateTrigger(trigger *Trigger) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(trigger)
	if err != nil {
		return err
	}
	_, err = c.httpRequest(fmt.Sprintf("triggers/%s", trigger.UUID), "PUT", buf)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTrigger removes an trigger from the server
func (c *Client) DeleteTrigger(uuid string) error {
	_, err := c.httpRequest(fmt.Sprintf("triggers/%s", uuid), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) httpRequest(path, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, c.requestPath(path), &body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", c.authToken)
	switch method {
	case "GET":
	case "DELETE":
	default:
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
		}
		return nil, fmt.Errorf("got a non 200 status code: %v - %s", resp.StatusCode, respBody.String())
	}
	return resp.Body, nil
}

func (c *Client) requestPath(path string) string {
	return fmt.Sprintf("%s/%s", "https://api.advancedscheduler.io", path)
}
