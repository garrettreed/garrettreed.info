package reading

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// BooksResponse represents the relevant response structure of Open Library Books api
type BooksResponse struct {
	Entries []struct {
		Work struct {
			Title   string   `json:"title"`
			Authors []string `json:"author_names"`
		} `json:"work"`
	} `json:"reading_log_entries"`
}

// Book represents a publication with a title and author(s)
type Book struct {
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
}

// GetCurrentlyReading requests Open Library Books api for items in the user's currently
// reading bookshelf, and parses the response to build a list of Books.
func GetCurrentlyReading() (books []Book, err error) {
	endpoint := fmt.Sprintf("https://openlibrary.org/people/%s/books/currently-reading.json", os.Getenv("OPENLIBRARY_USER"))

	booksClient := http.Client{
		Timeout: time.Second * 5,
	}

	req, reqErr := http.NewRequest(http.MethodGet, endpoint, nil)
	if reqErr != nil {
		fmt.Println("Error creating request:", reqErr)
		return nil, reqErr
	}

	res, getErr := booksClient.Do(req)
	if getErr != nil || res.StatusCode != http.StatusOK {
		fmt.Println("Response error:", getErr)
		return nil, errors.New("failed to query open library books api")
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	var resp BooksResponse
	jsonErr := json.Unmarshal(body, &resp)
	if jsonErr != nil {
		return nil, jsonErr
	}

	books = []Book{}
	for _, book := range resp.Entries {
		books = append(
			books,
			Book{
				Title:   book.Work.Title,
				Authors: book.Work.Authors,
			},
		)
	}

	return books, nil
}
