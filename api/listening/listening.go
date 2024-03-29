package listening

import (
	"encoding/json"
	"fmt"
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
func GetRecentTracks() (tracks []Track, err error) {
	endpoint := fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=%s&api_key=%s+&format=json", os.Getenv("LASTFM_USER"), os.Getenv("LASTFM_API_KEY"))

	lastFmClient := http.Client{
		Timeout: time.Second * 5,
	}

	req, reqErr := http.NewRequest(http.MethodGet, endpoint, nil)
	if reqErr != nil {
		return nil, reqErr
	}

	res, getErr := lastFmClient.Do(req)
	if getErr != nil {
		return nil, fmt.Errorf("last.fm api request failed: %v", getErr)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("last.fm api request failed with status code: %d", res.StatusCode)
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
