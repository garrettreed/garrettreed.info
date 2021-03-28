package aggregate

import (
	"errors"

	"github.com/garrettreed/garrettreed.info/api/listening"
	"github.com/garrettreed/garrettreed.info/api/reading"
)

type SiteData struct {
	Listening []listening.Track `json:"listening"`
	Reading   []reading.Book    `json:"reading"`
}

type readingResult struct {
	Reading []reading.Book
	Error   error
}

type listeningResult struct {
	Listening []listening.Track
	Error     error
}

func getReadingData(readingChan chan readingResult) {
	reading, readingErr := reading.GetCurrentlyReading()
	readingChan <- readingResult{reading, readingErr}
}

func getListeningData(listeningChan chan listeningResult) {
	listening, listeningErr := listening.GetRecentTracks()
	listeningChan <- listeningResult{listening, listeningErr}
}

func GetAggregateData() (sd *SiteData, err error) {
	readingChan := make(chan readingResult)
	listeningChan := make(chan listeningResult)
	go getReadingData(readingChan)
	go getListeningData(listeningChan)
	rs := <-readingChan
	ls := <-listeningChan

	if rs.Error != nil || ls.Error != nil {
		return sd, errors.New("Error")
	}

	return &SiteData{ls.Listening, rs.Reading}, nil
}
