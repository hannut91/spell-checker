package checker

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const MAX_LENGTH = 500

type SpellResult struct {
	Text string `json:"notag_html"`
}

type SpellMessage struct {
	Result SpellResult `json:"result"`
}

type SpellResponse struct {
	Message SpellMessage `json:"message"`
}

// Check returns fixed text.
func Check(text string) (fixed string, err error) {
	if len(text) >= MAX_LENGTH {
		text = text[:MAX_LENGTH]
	}

	body, err := fetchURL(getSpellURL(text))
	if err != nil {
		return
	}

	var response SpellResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	fixed = response.Message.Result.Text
	return
}

func fetchURL(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	return
}

func getSpellURL(text string) string {
	return "https://m.search.naver.com/p/csearch/ocontent/util/" +
		"SpellerProxy?color_blindness=0&q=" +
		url.QueryEscape(text)
}
