package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Song struct {
	Title string `json:"title"`
	Artist string `json:"artist"`
}

type BillboardResponse struct {
	Songs []Song `json:"content"`
}
// func (c *Client) Get(url string) (resp *Response, err error) {
// 	req, err := NewRequest("GET", url, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return c.Do(req)
// }

func GetTopSongForDate() (Song, error){
	req, err := http.NewRequest("GET", "https://billboard-api2.p.rapidapi.com/hot-100?date=2019-05-11&range=1-1", nil)
	// resp, err := http.Get("https://billboard-api2.p.rapidapi.com/hot-100?date=2019-05-11&range=1-1")

	if (err != nil) {
		return Song{}, err
	}
	
	req.Header.Add("X-RapidAPI-Key", "c85b222c9emshc81e5e2b9971431p10dbb7jsn6d3def4dad64")
	req.Header.Add("X-RapidAPI-Host", "billboard-api2.p.rapidapi.com")

	resp, err := http.DefaultClient.Do(req)

	if (err != nil) {
		return Song{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if (err != nil) {
		return Song{}, err
	}

	billboardResponse := BillboardResponse{}

	err = json.Unmarshal(body, &billboardResponse)

	if (err != nil) {
		return Song{}, fmt.Errorf("failed to unmarshal JSON: %s", err)
	}

	return billboardResponse.Songs[0], nil;
}