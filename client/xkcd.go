package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

const (
	baseURL       string        = "https://xkcd.com"
	clientTimeout time.Duration = 30 * time.Second
)

type xkcdClient struct {
	client  *http.Client
	baseURL string
}

func (c *xkcdClient) buildURL(comicNo int) string {
	if comicNo == 0 {
		return fmt.Sprintf("%s/info.0.json", c.baseURL)
	}

	return fmt.Sprintf("%s/%d/info.0.json", c.baseURL, comicNo)
}

func (c *xkcdClient) Fetch(comicNo int, save bool) (xkcdResponse, error) {
	resp, err := c.client.Get(c.buildURL(comicNo))
	if err != nil {
		return xkcdResponse{}, err
	}
	defer resp.Body.Close()

	var xkcdResp xkcdResponse
	if err := json.NewDecoder(resp.Body).Decode(&xkcdResp); err != nil {
		return xkcdResponse{}, nil
	}

	if save {
		if err := c.save(xkcdResp.Img, path.Base(xkcdResp.Img)); err != nil {
			return xkcdResponse{}, err
		}
	}

	return xkcdResp, nil
}

func (c *xkcdClient) save(imgURL string, path string) error {
	resp, err := c.client.Get(imgURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	img, err := os.Create(path)
	if err != nil {
		return err
	}
	defer img.Close()

	_, err = io.Copy(img, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func NewXKCDClient() *xkcdClient {
	return &xkcdClient{
		client: &http.Client{
			Timeout: clientTimeout,
		},
		baseURL: baseURL,
	}
}
