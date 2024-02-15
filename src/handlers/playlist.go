package handlers

import (
	"fmt"
	"time"
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

	return fmt.Sprintf("I would build a date starting on %s", parsedDate), nil
}