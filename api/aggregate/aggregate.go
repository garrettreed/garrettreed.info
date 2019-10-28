package aggregate

import (
    "github.com/garrettreed/garrettreed.info/api/listening"
    "github.com/garrettreed/garrettreed.info/api/reading"
)

type SiteData struct {
    Listening listening.LastFmRecentTracks `json:"listening"`
    Reading reading.GoodreadsCurrentlyReading `json:"reading"`
}

func GetAggregateData() (siteData SiteData, err error) {
    reading, readingErr := reading.GetCurrentlyReading()
    if readingErr != nil {
        return siteData, readingErr
    }

    listening, musicErr := listening.GetRecentTracks()
    if musicErr != nil {
        return siteData, musicErr
    }

    return SiteData{listening, reading}, nil
}
