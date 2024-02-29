package handlers

import (
	"fmt"
	"time"

	"github.com/braydend/birthday-jamz/src/api"
)

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


func BuildPlaylistHandler(birthday string) (string, error) {
	parsedDate, err := formatDate(birthday);

	if (err != nil) {
		return "", fmt.Errorf("unable to parse date")
	}

    topSong, err := api.GetTopSongForDate(parsedDate)

    if (err != nil) {
		return "", err
	}

	return fmt.Sprintf("top song on %s is %s", parsedDate, topSong.Title), nil
}