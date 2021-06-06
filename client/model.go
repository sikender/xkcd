package client

import "encoding/json"

type xkcdResponse struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func (r xkcdResponse) JSON() string {
	rJSON, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return ""
	}
	return string(rJSON)
}
