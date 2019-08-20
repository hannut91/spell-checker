package checker

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SpellResult struct {
	Text string `json:"notag_html"`
}

type SpellMessage struct {
	Result SpellResult `json:"result"`
}

type SpellResponse struct {
	Message SpellMessage `json:"message"`
}

// Add returns sum of two numbers.
func Add(a int, b int) int {
	return a + b
}

// CheckFile return fixed text.
func CheckFile(filename string) (fixed string, err error) {
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return
	}

	fixed, err = Check(string(body))

	return
}

// Check returns fixed text.
func Check(text string) (fixed string, err error) {
	body, err := fetchURL(getSpellURL(text))

	if err != nil {
		return
	}

	response, err := parseSpellResponse(body)

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

func parseSpellResponse(body []byte) (response SpellResponse, err error) {
	err = json.Unmarshal(body, &response)

	return
}

func getSpellURL(text string) string {
	return "https://m.search.naver.com/p/csearch/ocontent/util/" +
		"SpellerProxy?color_blindness=0&q=" +
		url.QueryEscape(text)
}
