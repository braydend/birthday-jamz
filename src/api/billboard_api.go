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

func GetTopSongForDate(dateString string) (Song, error){
	urlWithDate := fmt.Sprintf("https://billboard-api2.p.rapidapi.com/hot-100?date=%s&range=1-1", dateString);
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