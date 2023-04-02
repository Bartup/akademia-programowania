package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type response struct {
	Data struct {
		Children []struct {
			Data struct {
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type RedditFetcher interface {
	Fetch() error
	Save(io.Writer) error
}

type HttpFetcher struct {
	URL  string
	Resp response
}

func (fet *HttpFetcher) Fetch() error {
	request, err := http.NewRequest(http.MethodGet, fet.URL, nil)
	if err != nil {
		return fmt.Errorf("cannot create request: %w", err)
	}

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("cannot get data: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("cannot read body: %w", err)
	}

	err = json.Unmarshal(resBody, &fet.Resp)
	if err != nil {
		return fmt.Errorf("cannot unmarshal data: %w", err)
	}
	return nil
}

func (fet *HttpFetcher) Save(writer io.Writer) error {
	for _, child := range fet.Resp.Data.Children {
		_, err := writer.Write([]byte(child.Data.Title + "\n" + child.Data.URL + "\n"))
		if err != nil {
			return fmt.Errorf("cannot write data: %w", err)
		}
	}
	return nil
}
