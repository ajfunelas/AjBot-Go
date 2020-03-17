package funcs

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/mmcdole/gofeed"
)



func getURL(u *url.URL) (string, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return "", fmt.Errorf("build request: %v", err)
	}
	req.Header.Set("User-Agent", "Reddit_RSS_Reader/0.0.1")
	rssResponse, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("get url: %v", err)
	}
	b, err := ioutil.ReadAll(rssResponse.Body)
	if err != nil {
		return "", fmt.Errorf("read io.Reader: %w", err)
	}
	defer rssResponse.Body.Close()
	return string(b), nil
}

// ReadRSSFeed receives URL u and returns parsed feed *gofeed.Feed for further processing
func ReadRSSFeed(u string) (*gofeed.Feed, error) {
	uri, err := url.Parse(u)
	if err != nil {
		return nil, fmt.Errorf("parse url: %w", err)
	}
	body, err := getURL(uri)
	if err != nil {
		return nil, fmt.Errorf("get feed: %w", err)
	}
	fp := gofeed.NewParser()
	feed, err := fp.ParseString(body)
	if err != nil {
		return nil, fmt.Errorf("parse feed: %w", err)
	}
	return feed, nil

}