package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type searchData struct {
	Key   string
	Name  []interface{}
	Blank []interface{}
	Link  []interface{}
}

type individualPage struct {
	PageID  int
	Ns      int
	Title   string
	Extract string
}

type query struct {
	Pages map[string]individualPage
}

type pageData struct {
	BatchComplete string
	Warnings      map[string]interface{}
	Query         query
}

func (d *searchData) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&d.Key, &d.Name, &d.Blank, &d.Link}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Notification: %d != %d", g, e)
	}
	return nil
}

func searchWiki(keyword string) searchData {
	keyword = url.QueryEscape(keyword)
	url := "https://en.wikipedia.org//w/api.php?action=opensearch&format=json&search=" + keyword + "&profile=fuzzy&limit=5"

	body, err := apiCall(url)
	if err != nil {
		log.Panicln(err)
	}

	var resultGot searchData

	err = json.Unmarshal(body, &resultGot)
	if err != nil {
		log.Panicln(err)
	}

	return resultGot
}

func getWikiPage(keyword string) pageData {
	keyword = url.QueryEscape(keyword)
	url := "https://en.wikipedia.org/w/api.php?action=query&prop=extracts&format=json&titles=" + keyword

	body, err := apiCall(url)
	if err != nil {
		log.Panicln(err)
	}

	var resultGot pageData

	err = json.Unmarshal(body, &resultGot)
	if err != nil {
		log.Panicln(err)
	}

	return resultGot
}

func apiCall(url string) ([]byte, error) {
	wikiClient := http.Client{Timeout: time.Second * 20}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "wikiui-cli")

	res, getErr := wikiClient.Do(req)
	if getErr != nil {
		return nil, getErr
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
