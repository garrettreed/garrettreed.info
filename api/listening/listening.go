package listening

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// LastFmRecentTracks represents the relevant response structure of last.fm's
// audioscrobbler recent tracks api.
type LastFmRecentTracks struct {
	RecentTracks struct {
		Track []struct {
			TrackName string `json:"name"`
			TrackURL  string `json:"url"`
			Album     struct {
				Title string `json:"#text"`
			} `json:"album"`
			Artist struct {
				Title string `json:"#text"`
			} `json:"artist"`
		} `json:"track"`
	} `json:"recenttracks"`
}

// Track represents a single song
type Track struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Album  string `json:"album"`
	Artist string `json:"artist"`
}

// GetRecentTracks requests Last.fm's api for a user's recently-listened
// tracks and parses the response to build a list of Books.
// TODO: impelement unmarshal method that uses the decoder api to enforce DisallowUnknownFields
func GetRecentTracks() (tracks []Track, err error) {
	var endpoint string = "http://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=" + os.Getenv("LASTFM_USER") + "&api_key=" + os.Getenv("LASTFM_API_KEY") + "&format=json"

	lastFmClient := http.Client{
		Timeout: time.Second * 5,
	}

	req, reqErr := http.NewRequest(http.MethodGet, endpoint, nil)
	if reqErr != nil {
		return nil, reqErr
	}

	res, getErr := lastFmClient.Do(req)
	if getErr != nil || res.StatusCode != http.StatusOK {
		return nil, errors.New("Failed to query last.fm api.")
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	var lastFmTracks LastFmRecentTracks
	jsonErr := json.Unmarshal(body, &lastFmTracks)
	if jsonErr != nil {
		return nil, jsonErr
	}

	tracks = []Track{}
	for _, track := range lastFmTracks.RecentTracks.Track {
		tracks = append(
			tracks,
			Track{
				Name:   track.TrackName,
				URL:    track.TrackURL,
				Album:  track.Album.Title,
				Artist: track.Artist.Title,
			},
		)
	}

	return tracks, nil
}
