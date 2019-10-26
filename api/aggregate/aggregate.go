package aggregate

import (
    "github.com/garrettreed/garrettreed.info/api/music"
)

type SiteData struct {
    Listening music.LastFmRecentTracks `json:"listening"`
}

func GetAggregateData() (siteData SiteData, err error) {
    listening, musicErr := music.GetRecentTracks()

    if musicErr != nil {
        return siteData, musicErr
    }

    return SiteData{listening}, nil
}
