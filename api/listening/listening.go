package listening

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type LastFmRecentTracks struct {
	RecentTracks struct {
		Track []struct {
			TrackName string `json:"name"`
			TrackUrl  string `json:"url"`
			Album     struct {
				Title string `json:"#text"`
			} `json:"album"`
			Artist struct {
				Title string `json:"#text"`
			} `json:"artist"`
		} `json:"track"`
	} `json:"recenttracks"`
}

type RecentTracks struct {
	Tracks []Track `json:"tracks"`
}

type Track struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Album  string `json:"album"`
	Artist string `json:"artist"`
}

// TODO: impelement unmarshal method that uses the decoder api to enforce DisallowUnknownFields
func GetRecentTracks() (recentTracks *RecentTracks, err error) {
	var endpoint string = "http://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=" + os.Getenv("LASTFM_USER") + "&api_key=" + os.Getenv("LASTFM_API_KEY") + "&format=json"

	lastFmClient := http.Client{
		Timeout: time.Second * 5,
	}

	req, reqErr := http.NewRequest(http.MethodGet, endpoint, nil)
	if reqErr != nil {
		return nil, reqErr
	}

	res, getErr := lastFmClient.Do(req)
	if getErr != nil {
		return nil, getErr
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

	recentTracks = &RecentTracks{}
	for _, track := range lastFmTracks.RecentTracks.Track {
		recentTracks.Tracks = append(
			recentTracks.Tracks,
			Track{
				Name: track.TrackName,
				Url: track.TrackUrl,
				Album: track.Album.Title,
				Artist: track.Artist.Title,
			},
		)
	}

	return recentTracks, nil
}
