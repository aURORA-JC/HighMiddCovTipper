package model

// Structs Area

// Body The request body of POST
type Body struct {
	AppID           string `json:"appId"`
	PaasHeader      string `json:"paasHeader"`
	TimestampHeader string `json:"timestampHeader"`
	NonceHeader     string `json:"nonceHeader"`
	SignatureHeader string `json:"signatureHeader"`
	Key             string `json:"key"`
}

type Covid19Data struct {
	Data struct {
		EndUpdateTime string `json:"end_update_time"`
		Hcount        int    `json:"hcount"`
		Mcount        int    `json:"mcount"`
		HighList      []struct {
			Type       string   `json:"type"`
			Province   string   `json:"province"`
			City       string   `json:"city"`
			County     string   `json:"county"`
			AreaName   string   `json:"area_name"`
			Communities []string `json:"communitys"`
		} `json:"highlist"`
		MiddleList []struct {
			Type        string   `json:"type"`
			Province    string   `json:"province"`
			City        string   `json:"city"`
			County      string   `json:"county"`
			AreaName    string   `json:"area_name"`
			Communities []string `json:"communitys"`
		} `json:"middlelist"`
	} `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
