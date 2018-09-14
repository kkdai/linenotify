package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ErrorResponse struct {
	ErrorCode int
	Err       error
}

const (
	urlEncodedContent string = "application/x-www-form-urlencoded"
	apiToken          string = "https://notify-bot.line.me/oauth/token"
	apiNotify         string = "https://notify-api.line.me/api/notify"
)

func apiCall(mode string, inUrl string, data url.Values, token string) ([]byte, *ErrorResponse) {
	client := &http.Client{}
	fmt.Println("connected: ", mode, inUrl, data)

	r := &http.Request{}
	if data == nil {
		r, _ = http.NewRequest(mode, inUrl, nil)

	} else {
		r, _ = http.NewRequest(mode, inUrl, strings.NewReader(data.Encode()))
	}

	r.Header.Add("Content-Type", urlEncodedContent)
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	if len(token) != 0 {
		r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	ret := new(ErrorResponse)
	resp, err := client.Do(r)

	if err != nil {
		log.Println("er:", err)
		ret.Err = err
		return nil, ret
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("er:", err)
		ret.Err = err
		return nil, ret
	}

	if resp.StatusCode > http.StatusAccepted {
		ret.ErrorCode = resp.StatusCode
		ret.Err = errors.New("Error on:" + string(body))
		log.Println("Error happen! body:", string(body))
		return body, ret
	}

	return body, nil
}
