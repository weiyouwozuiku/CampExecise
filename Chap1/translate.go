package chap1

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
}

type DictResp struct {
	Rc   int `json:"rc"`
	Wiki struct {
		KnownInLaguages int `json:"known_in_laguages"`
		Description     struct {
			Source string      `json:"source"`
			Target interface{} `json:"target"`
		} `json:"description"`
		ID   string `json:"id"`
		Item struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"item"`
		ImageURL  string `json:"image_url"`
		IsSubject string `json:"is_subject"`
		Sitelink  string `json:"sitelink"`
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

func Transfer(r *DictRequest) error {
	url := "https://api.interpreter.caiyunai.com/v1/dict"
	method := "POST"
	buf, err := json.Marshal(r)
	if err != nil {
		return err
	}
	payload := bytes.NewReader(buf)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err
	}
	req.Header.Add("authority", "api.interpreter.caiyunai.com")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("app-name", "xy")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/json;charset=UTF-8")
	req.Header.Add("device-id", "")
	req.Header.Add("dnt", "1")
	req.Header.Add("origin", "https://fanyi.caiyunapp.com")
	req.Header.Add("os-type", "web")
	req.Header.Add("os-version", "")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("referer", "https://fanyi.caiyunapp.com/")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"104\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"104\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "cross-site")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
	req.Header.Add("x-authorization", "token:qgemv4jr1y38jyq6vhvi")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if res.StatusCode != 200 {
		return errors.New(string(res.StatusCode))
	}
	dictResp := DictResp{}
	err = json.Unmarshal(body, &dictResp)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", dictResp)
	return nil
}
