package aggregate

import (
	"fmt"

	"github.com/garrettreed/garrettreed.info/api/listening"
	"github.com/garrettreed/garrettreed.info/api/reading"
	"github.com/garrettreed/garrettreed.info/api/working"
)

type SiteData struct {
	Listening []listening.Track `json:"listening"`
	Reading   []reading.Book    `json:"reading"`
	Working   []working.Work    `json:"working"`
}

type readingResult struct {
	Reading []reading.Book
	Error   error
}

type listeningResult struct {
	Listening []listening.Track
	Error     error
}

type workingResult struct {
	Working []working.Work
	Error   error
}

func getReadingData(readingChan chan readingResult) {
	reading, readingErr := reading.GetCurrentlyReading()
	readingChan <- readingResult{reading, readingErr}
}

func getListeningData(listeningChan chan listeningResult) {
	listening, listeningErr := listening.GetRecentTracks()
	listeningChan <- listeningResult{listening, listeningErr}
}

func getWorkingData(workingChan chan workingResult) {
	working, workingErr := working.GetWorking()
	workingChan <- workingResult{working, workingErr}
}

func GetAggregateData() (sd *SiteData, err error) {
	readingChan := make(chan readingResult)
	listeningChan := make(chan listeningResult)
	workingChan := make(chan workingResult)
	go getReadingData(readingChan)
	go getListeningData(listeningChan)
	go getWorkingData(workingChan)
	rs := <-readingChan
	ls := <-listeningChan
	ws := <-workingChan

	if rs.Error != nil {
		return sd, fmt.Errorf("failed to get reading data: %v", rs.Error)
	}

	if ls.Error != nil {
		return sd, fmt.Errorf("failed to get listening data: %v", ls.Error)
	}

	if ws.Error != nil {
		return sd, fmt.Errorf("failed to get working data: %v", ls.Error)
	}

	return &SiteData{ls.Listening, rs.Reading, ws.Working}, nil
}
