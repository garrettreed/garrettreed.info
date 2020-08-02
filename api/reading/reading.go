package reading

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// GoodreadsCurrentlyReading represents the relevant response structure
// of the Goodreads reviews api.
// Selection query of XML starts with top tag,
// to prevent "GoodreadsResponse" from being included.
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

// Book represents a publication with a title and author(s)
type Book struct {
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
}

// GetCurrentlyReading requests Goodreads' api for items in the user's currently
// reading shelf, and parses the response to build a list of Books.
// TODO: impelement unmarshal method that uses the decoder api
// to enforce DisallowUnknownFields
func GetCurrentlyReading() (books []Book, err error) {
	var endpoint string = "https://www.goodreads.com/review/list?v=2&id=" +
		os.Getenv("GOODREADS_USER_ID") +
		"&shelf=currently-reading&key=" +
		os.Getenv("GOODREADS_API_KEY")

	goodreadsClient := http.Client{
		Timeout: time.Second * 5,
	}

	req, reqErr := http.NewRequest(http.MethodGet, endpoint, nil)
	if reqErr != nil {
		return nil, reqErr
	}

	res, getErr := goodreadsClient.Do(req)
	if getErr != nil {
		return nil, getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	var goodreadsBooks GoodreadsCurrentlyReading
	xmlErr := xml.Unmarshal(body, &goodreadsBooks)
	if xmlErr != nil {
		return nil, xmlErr
	}

	books = []Book{}
	for _, review := range goodreadsBooks.Reviews {
		var authors []string
		for _, author := range review.Book.Authors {
			authors = append(authors, author.Name)
		}
		books = append(books, Book{Title: review.Book.Title, Authors: authors})
	}

	return books, err
}
