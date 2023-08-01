package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	c       *http.Client
	baseUrl string
	config  Configuration
}

type Configuration struct {
	Form   map[string]string
	Header map[string]string
	Query  map[string]string
	Body   interface{}
	Entity interface{}
}

type HubspotTransport struct {
	apiKey string
}

func (t *HubspotTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", t.apiKey)
	req.Header.Add("content-type", "application/json")
	return http.DefaultTransport.RoundTrip(req)
}

func NewApi(baseUrl string, apiKey string) (*Client, error) {

	c := &http.Client{Transport: &HubspotTransport{
		apiKey: apiKey,
	}}

	return &Client{
		c: c,
		config: Configuration{
			Form:   nil,
			Header: nil,
		},
		baseUrl: baseUrl,
	}, nil
}

func (cli Client) Get(path string, result interface{}, config ...Configuration) (interface{}, error) {
	var conf Configuration

	endpoint, err := url.ParseRequestURI(cli.baseUrl + path)
	if err != nil {
		return nil, err
	}

	var res *http.Response

	for i := 0; i < 10; i++ {

		req := http.Request{
			Method: "GET",
			URL:    endpoint,
		}

		if len(config) != 0 {
			conf = config[0]
		} else {
			conf = cli.config
		}

		cli.addHeaders(&req, conf)
		cli.addQuery(&req, conf)
		//cli.addForm(&req, conf)

		//fmt.Println(req.URL.String())

		tmpReq := req

		res, err = cli.c.Do(&tmpReq)
		if err != nil {
			return nil, err
		}

		if res.StatusCode != 429 {
			break
		} else {
			time.Sleep(1 * time.Second)
		}
	}

	if res.StatusCode != 200 {
		return []byte{}, errors.New(fmt.Sprint("HB status code: ", res.StatusCode))
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp interface{}

	json.Unmarshal(bodyBytes, &result)
	json.Unmarshal(bodyBytes, &resp)

	return resp, nil
}

func (cli Client) Post(path string, result interface{}, body interface{}, config ...Configuration) (interface{}, error) {
	var conf Configuration
	endpoint, err := url.ParseRequestURI(cli.baseUrl + path)
	if err != nil {
		return nil, err
	}

	//fmt.Println(endpoint)

	var res *http.Response

	for i := 0; i < 10; i++ {

		req := http.Request{
			Method: "POST",
			URL:    endpoint,
		}

		jsonReq, _ := json.Marshal(body)
		bodyreq := bytes.NewBuffer(jsonReq)

		if len(config) != 0 {
			conf = config[0]
		} else {
			conf = cli.config
		}

		cli.addHeaders(&req, conf)
		cli.addForm(&req, conf)
		cli.addQuery(&req, conf)
		cli.addBody(&req, bodyreq)

		tmpReq := req

		res, err = cli.c.Do(&tmpReq)
		if err != nil {
			return nil, err
		}

		if res.StatusCode != 429 {
			break
		} else {
			time.Sleep(1 * time.Second)
		}
	}

	if res.StatusCode != 200 {
		return []byte{}, errors.New(fmt.Sprint("HB status code: ", res.StatusCode))
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	log.Println(string(bodyBytes))

	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}

	return string(bodyBytes), nil
}

func (c Client) Delete() {

}

func (cli Client) Patch(path string, result interface{}, body interface{}, config ...Configuration) (interface{}, error) {
	var conf Configuration
	endpoint, err := url.ParseRequestURI(cli.baseUrl + path)
	if err != nil {
		return nil, err
	}

	//fmt.Println(endpoint)

	var res *http.Response

	for i := 0; i < 10; i++ {

		req := http.Request{
			Method: "PATCH",
			URL:    endpoint,
		}

		jsonReq, _ := json.Marshal(body)
		bodyreq := bytes.NewBuffer(jsonReq)

		if len(config) != 0 {
			conf = config[0]
		} else {
			conf = cli.config
		}

		cli.addHeaders(&req, conf)
		cli.addForm(&req, conf)
		cli.addQuery(&req, conf)
		cli.addBody(&req, bodyreq)

		tmpReq := req

		res, err = cli.c.Do(&tmpReq)
		if err != nil {
			return nil, err
		}

		if res.StatusCode != 200 {
			return nil, errors.New(fmt.Sprint("Error actualizado deasl: ", res.StatusCode))
		}
	}

	if res.StatusCode != 200 {
		return []byte{}, errors.New(fmt.Sprint("HB status code: ", res.StatusCode))
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	log.Println(string(bodyBytes))

	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}

	return string(bodyBytes), nil

}

func (c Client) addHeaders(req *http.Request, conf Configuration) {
	for k, v := range conf.Header {
		req.Header.Add(k, v)
	}
}

func (c Client) addForm(req *http.Request, conf Configuration) {
	for k, v := range conf.Form {
		req.Form.Add(k, v)
	}
}

func (c Client) addBody(req *http.Request, body *bytes.Buffer) {
	r := io.NopCloser(body)
	req.Body = r
}

func (c Client) addQuery(req *http.Request, conf Configuration) {
	q := req.URL.Query()
	for k, v := range conf.Query {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
}
