package twitch

import (
	"log"
	"net/http"
	"net/url"
	"streaming_app/server/config"
	"streaming_app/server/services/api"
	"strings"
	"time"
)

type QueryParams map[string]interface{}

func buildUrl(path string, queryParams QueryParams) (string, error) {
	var result string

	config := config.Read()
	baseUrl := config.Twitch.Url

	twitchURL, err := url.Parse(baseUrl)

	if err != nil {
		log.Println(err.Error())
		return result, err
	}

	twitchURL.Path += path
	params := url.Values{}

	for k, v := range queryParams {
		params.Add(
			strings.ToLower(k),
			strings.ToLower(v.(string)),
		)
	}

	twitchURL.RawQuery = params.Encode()
	result = twitchURL.String()

	return result, nil
}

func MakeRequest(method string, path string, queryParams QueryParams) (*http.Response, api.WebError) {
	var res *http.Response

	config := config.Read()
	clientId := config.Twitch.Id

	requestURL, err := buildUrl(path, queryParams)

	if err != nil {
		log.Println(err.Error())
		return res, api.ErrInternalServerError
	}

	req, err := http.NewRequest(method, requestURL, nil)

	if err != nil {
		log.Println(err.Error())
		return res, api.ErrInternalServerError
	}

	req.Header.Add("Client-Id", clientId)
	req.Header.Add("Accept", "application/vnd.twitchtv.v5+json")

	client := &http.Client{Timeout: 2 * time.Second} // 2 second timeout
	res, err = client.Do(req)

	if err != nil {
		log.Println(err.Error())
		return res, api.ErrInternalServerError
	}

	return res, api.WebError{}
}

func Get(uri string, queryParams QueryParams) (*http.Response, api.WebError) {
	return MakeRequest("GET", uri, queryParams)
}
