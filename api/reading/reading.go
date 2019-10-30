package reading

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Selection query of XML starts with top tag,
// so "GoodreadsResponse" won't be included.
type GoodreadsCurrentlyReading struct {
	Reviews []struct {
		Book struct {
			Title   string `xml:"title" json:"title"`
			Authors []struct {
				Name string `xml:"name" json:"name"`
			} `xml:"authors>author" json:"authors"`
		} `xml:"book" json:"book"`
	} `xml:"reviews>review" json:"reviews"`
}

// TODO: impelement unmarshal method that uses the decoder api to enforce DisallowUnknownFields
// TODO: export as non-goodreads normalized structure
func GetCurrentlyReading() (books GoodreadsCurrentlyReading, err error) {
	var endpoint string = "https://www.goodreads.com/review/list?v=2&id=" + os.Getenv("GOODREADS_USER_ID") + "&shelf=currently-reading&key=" + os.Getenv("GOODREADS_API_KEY")

	goodreadsClient := http.Client{
		Timeout: time.Second * 5,
	}

	req, reqErr := http.NewRequest(http.MethodGet, endpoint, nil)
	if reqErr != nil {
		return books, reqErr
	}

	res, getErr := goodreadsClient.Do(req)
	if getErr != nil {
		return books, getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return books, readErr
	}

	xmlErr := xml.Unmarshal(body, &books)
	if xmlErr != nil {
		return GoodreadsCurrentlyReading{}, xmlErr
	}

	return books, err
}
