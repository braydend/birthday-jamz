package handlers

import (
	"fmt"
	"time"

	"github.com/braydend/birthday-jamz/src/api"
)

type PlaylistResponse struct {
    Date string `json:"date"`
    Songs map[string]api.Song `json:"playlist"`
}

func formatDate(dateParam string) (string, error) {
    // Parse the date string into a time.Time object
    parsedDate, err := time.Parse(time.DateOnly, dateParam)
    if err != nil {
        return "", err
    }

    // Format the date into a string
    dateString := parsedDate.Format(time.DateOnly)
    return dateString, nil
}


func BuildPlaylistHandler(birthday string) (PlaylistResponse, error) {
	parsedDate, err := formatDate(birthday);

	if (err != nil) {
		return PlaylistResponse{}, fmt.Errorf("unable to parse date")
	}

    topSong, err := api.GetTopSongForDate(parsedDate)

    if (err != nil) {
		return PlaylistResponse{}, err
	}

    resp := PlaylistResponse{
        Date: parsedDate,
        Songs: map[string]api.Song{"foo":topSong},
    }

    return resp, err
}