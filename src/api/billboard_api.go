package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Song struct {
	Title string `json:"title"`
	Artist string `json:"artist"`
}

type BillboardResponse struct {
	Songs map[string]Song `json:"content"`
}

// TODO: Make this cleaner. This is good enoughf or testing so we dont exceed the 30 req/month quota
const BASE_URL = "https://billboard-api2.p.rapidapi.com/hot-100?date=%s&range=1-1";
const STUB_URL = "https://gist.githubusercontent.com/braydend/561312748fec29d78203949e21c49402/raw/623fce242a8e0a2e0f587ea9695fe72b4925e3e5/stub.json"

func GetTopSongForDate(dateString string) (Song, error){
	// urlWithDate := fmt.Sprintf(BASE_URL, dateString);
	urlWithDate := STUB_URL
	req, err := http.NewRequest("GET", urlWithDate, nil)

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
		log.Print(err)
		return Song{}, err
	}

	billboardResponse := BillboardResponse{}

	err = json.Unmarshal(body, &billboardResponse)

	if (err != nil) {
		log.Print(err)
		return Song{}, fmt.Errorf("failed to unmarshal JSON: %s", err)
	}

	return billboardResponse.Songs["1"], nil;
}