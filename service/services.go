package service

import (
	"HighMiddCovTipper/config"
	"HighMiddCovTipper/model"
	"HighMiddCovTipper/template"
	"HighMiddCovTipper/util"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

//generateParma Generate the POST parma
func generateParma() (interface{}, string) {
	timestamp := strconv.FormatInt(util.GetTimestamp(), 10)
	return &model.Body{
		AppID:           config.ApiAppid,
		PaasHeader:      config.ApiPassid,
		TimestampHeader: timestamp,
		NonceHeader:     config.ApiNonce,
		SignatureHeader: util.GetSHA256String(fmt.Sprint(timestamp, config.ApiToken, config.ApiNonce, timestamp)),
		Key:             config.ApiKey,
	}, timestamp
}

// GetCovData Get Data from API
func GetCovData() model.Covid19Data {
	// Generate POST Body
	parma, timestamp := generateParma()
	// Calculate signature with function getSHA256String
	signature := util.GetSHA256String(fmt.Sprint(timestamp, config.SignatureKey, timestamp))
	// Convert to JSON Data
	jsonData, _ := json.Marshal(parma)

	// Customer's Client
	client := &http.Client{}
	// Set request's methode & url & body
	req, _ := http.NewRequest("POST", config.ApiUrl, bytes.NewReader(jsonData))
	// Set Header's keys & values
	req.Header.Set("Content-Type", config.ContentType)
	req.Header.Set("User-Agent", config.UserAgent)
	req.Header.Set("x-wif-nonce", config.XWifNonce)
	req.Header.Set("x-wif-paasid", config.XWifPassid)
	req.Header.Set("x-wif-signature", signature)
	req.Header.Set("x-wif-timestamp", timestamp)

	// Send request to target server
	resp, postError := client.Do(req)

	if postError != nil {
		return model.Covid19Data{Code: 114514}
	}

	// Close the stream at the end
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	// Decode with JSON & struct
	dec := json.NewDecoder(resp.Body)
	var info model.Covid19Data
	decodeError := dec.Decode(&info)
	if decodeError != nil {
		return model.Covid19Data{Code: 114514}
	}

	return info
}

// send2Users Send anything to users
func send2Users(title string, content string) {
	resp, postError := http.PostForm(fmt.Sprint(config.ServerChanApiUrl, config.ServerChanPushToken, config.ServerChanPushMisc),
		url.Values{
			"title": {title},
			"desp":  {content},
		})
	if postError != nil {
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	// Read string from io
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// SendCovInfo2Users Send information to users
func SendCovInfo2Users(content string) {
	send2Users(template.PushTitleTemplate, content)
}
