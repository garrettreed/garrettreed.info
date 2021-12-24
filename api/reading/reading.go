package reading

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Volumes represents the relevant response structure of Google Books
// Bookshelf Volumes List api endpoint
type Volumes struct {
	Items []struct {
		VolumeInfo struct {
			Title   string   `json:"title"`
			Authors []string `json:"authors"`
		} `json:"volumeInfo"`
	} `json:"items"`
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
	endpoint := fmt.Sprintf("https://www.googleapis.com/books/v1/users/%s/bookshelves/%s/volumes", os.Getenv("GOOGLE_BOOKS_USER_ID"), os.Getenv("GOOGLE_BOOKS_BOOKSHELF_ID"))

	booksClient := http.Client{
		Timeout: time.Second * 5,
	}

	req, reqErr := http.NewRequest(http.MethodGet, endpoint, nil)
	if reqErr != nil {
		return nil, reqErr
	}

	res, getErr := booksClient.Do(req)
	if getErr != nil || res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to query google books api")
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	var volumes Volumes
	jsonErr := json.Unmarshal(body, &volumes)
	if jsonErr != nil {
		return nil, jsonErr
	}

	books = []Book{}
	for _, book := range volumes.Items {
		books = append(
			books,
			Book{
				Title:   book.VolumeInfo.Title,
				Authors: book.VolumeInfo.Authors,
			},
		)
	}

	return books, nil
}
