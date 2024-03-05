package util

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	neturl "net/url"
	"strconv"
	"strings"
	"time"
)

var HttpClinet *hClient

type hClient struct {
	HttpClinet http.Client
}

type HError struct {
	ErrMsg  string
	Timeout bool
}

func (he *HError) Error() string {
	return he.ErrMsg
}

func (he *HError) IsTimeout() bool {
	return he.Timeout
}

func IsTimeout(err error) bool {
	if err == nil {
		return false
	}
	if hErr, ok := err.(*HError); ok {
		return hErr.IsTimeout()
	}
	return false
}

// url - post url
// body - post body data
// ret - response body data, mast points to a block of memory
func (h *hClient) Post(url, contentType, body string, ret interface{}) error {
	if h.HttpClinet.Timeout == 0 {
		h.HttpClinet.Timeout = 5 * time.Second
	}
	bodyIo := strings.NewReader(body)
	resp, err := h.HttpClinet.Post(url, contentType, bodyIo)
	if err != nil {
		netErr, ok := err.(*neturl.Error)
		hErr := &HError{}
		if ok && netErr.Timeout() {
			hErr.ErrMsg = "timeout"
			hErr.Timeout = true
			return hErr
		}
		hErr.ErrMsg = err.Error()
		return err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.Unmarshal(respBody, ret)
	if resp.StatusCode != http.StatusOK {
		return errors.New("statusCode is " + strconv.Itoa(resp.StatusCode) + "body:" + string(respBody))
	}
	return err
}

func (h *hClient) JsonPost(url string, body string, ret interface{}) error {
	err := h.Post(url, "application/json", body, ret)
	return err
}

func (h *hClient) FormPost(url string, body string, ret interface{}) error {
	err := h.Post(url, "application/x-www-form-urlencoded", body, ret)
	return err
}

func (h *hClient) Do(req *http.Request, ret interface{}) error {
	resp, err := h.HttpClinet.Do(req)

	if err != nil {
		netErr, ok := err.(*neturl.Error)
		hErr := &HError{}
		if ok && netErr.Timeout() {
			hErr.ErrMsg = "timeout"
			hErr.Timeout = true
			return hErr
		}
		hErr.ErrMsg = err.Error()
		return err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.Unmarshal(respBody, ret)
	if resp.StatusCode != http.StatusOK {
		return errors.New("statusCode is " + strconv.Itoa(resp.StatusCode) + "body:" + string(respBody))
	}
	return err
}

func (h *hClient) DoReq(req *http.Request) (*http.Response, error) {
	resp, err := h.HttpClinet.Do(req)

	if err != nil {
		netErr, ok := err.(*neturl.Error)
		hErr := &HError{}
		if ok && netErr.Timeout() {
			hErr.ErrMsg = "timeout"
			hErr.Timeout = true
			return nil, hErr
		}
		hErr.ErrMsg = err.Error()
		return nil, err
	}

	return resp, err
}

// 默认60秒
func init() {
	HttpClinet = &hClient{
		HttpClinet: http.Client{
			Timeout: time.Duration(60) * time.Second,
		},
	}
}
