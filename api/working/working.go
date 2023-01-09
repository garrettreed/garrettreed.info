package working

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// Gist represents the content response structure of a Github Gist with the
// filename `work.csv`
type Gist struct {
	Files struct {
		Work struct {
			Content string `json:"content"`
		} `json:"work.csv"`
	} `json:"files"`
}

type Work struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

// GetWorking requests Github Gist api for lines in the work.csv file.
// It parses the file to build a list of Work items.
func GetWorking() (work []Work, err error) {
	endpoint := fmt.Sprintf("https://api.github.com/gists/%s", os.Getenv("GIST_ID"))

	gistClient := http.Client{
		Timeout: time.Second * 5,
	}

	req, reqErr := http.NewRequest(http.MethodGet, endpoint, nil)
	if reqErr != nil {
		return nil, reqErr
	}

	res, getErr := gistClient.Do(req)
	if getErr != nil || res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to query github gist api")
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	var gist Gist
	jsonErr := json.Unmarshal(body, &gist)
	if jsonErr != nil {
		return nil, jsonErr
	}

	lines := strings.Split(gist.Files.Work.Content, `\n`)
	for _, line := range lines {
		parts := strings.Split(line, ",")
		workEntry := Work{}
		for i, part := range parts {
			switch i {
			case 0:
				workEntry.Title = part
			case 1:
				workEntry.Url = part
			}
		}
		work = append(work, workEntry)
	}

	return work, nil
}
