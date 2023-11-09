package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

func makeHttpRequest[T any](fullUrl string, hhtpMethod string, headers map[string]string, queryParameters url.Values, body io.Reader, responseType T) (T, error) {
	client := http.Client{}

	u, err := url.Parse(fullUrl)
	if err != nil {
		return responseType, err
	}

	if hhtpMethod == "GET" {
		q := u.Query()
		for k, v := range queryParameters {
			q.Add(k, v[0])
		}
		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(hhtpMethod, u.String(), body)
	if err != nil {
		return responseType, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return responseType, err

	}

	if res == nil {
		return responseType, nil
	}

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return responseType, err

	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return responseType, err
	}

	var responseObject T
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return responseType, err
	}

	return responseObject, nil

}
