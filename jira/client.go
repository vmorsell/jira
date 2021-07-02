package jira

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"github.com/vmorsell/jira/authstore"
)

type Client struct {
	Tenant string
	Email  string
	Token  string
}

func New() *Client {
	return &Client{}
}

func (c *Client) WithCredsFromFile() *Client {
	s := authstore.New()

	var creds *authstore.Credentials
	creds, err := s.Read()
	if err != nil {
		creds = &authstore.Credentials{}
	}

	c.Tenant = creds.Tenant
	c.Email = creds.Email
	c.Token = creds.Token

	return c
}

func (c *Client) Req(method, url string) ([]byte, error) {
	fullURL := fmt.Sprintf("%s/%s", baseURL(c.Tenant), url)

	req, err := http.NewRequest(method, fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("req: %w", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", c.basicAuthToken()))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do: %w", err)
	}

	if res.Status != "200 OK" {
		return nil, fmt.Errorf("res status %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}
	return body, nil
}

func (c *Client) basicAuthToken() string {
	data := []byte(fmt.Sprintf("%s:%s", c.Email, c.Token))
	return base64.StdEncoding.EncodeToString(data)
}

func baseURL(tenant string) string {
	return fmt.Sprintf("https://%s.atlassian.net/rest/api/3", tenant)
}
