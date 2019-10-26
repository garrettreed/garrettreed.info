package music

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
    "time"
    "os"
)

type LastFmRecentTracks struct {
	RecentTracks struct {
        Track []struct {
            TrackName string `json:"name"`
            TrackUrl string `json:"url"`
            Album struct {
                Title string `json:"#text"`
            } `json:"album"`
            Artist struct {
                Title string `json:"#text"`
            } `json:"artist"`
        } `json:"track"`
	} `json:"recenttracks"`
}

// TODO: impelement unmarshal method that uses the decoder api to enforce DisallowUnknownFields
func GetRecentTracks() (tracks LastFmRecentTracks, err error) {
	var endpoint string = "http://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user="+ os.Getenv("LASTFM_USER") + "&api_key=" + os.Getenv("LASTFM_API_KEY") + "&format=json"

	lastFmClient := http.Client{
		Timeout: time.Second * 5,
	}

	req, reqErr := http.NewRequest(http.MethodGet, endpoint, nil)
	if reqErr != nil {
		return tracks, reqErr
	}

	res, getErr := lastFmClient.Do(req)
	if getErr != nil {
		return tracks, getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return tracks, readErr
	}

	jsonErr := json.Unmarshal(body, &tracks)
	if jsonErr != nil {
		return LastFmRecentTracks{}, jsonErr
    }

	return tracks, err
}
