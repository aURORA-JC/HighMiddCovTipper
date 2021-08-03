package config

// ServerChan Area

// ServerChanApiUrl ServerChan API's Domain
const ServerChanApiUrl = "https://sctapi.ftqq.com/"

// ServerChanPushToken User's ServerChan Token
// the key will changed by args when script start
var ServerChanPushToken = "SCT49339TdVuPhktIqXCEUrRNIlJWxRur"

// ServerChanPushMisc If the API require, use this
// For example, the ServerChan's API address is 'https://sctapi.ftqq.com/SCT49339TdVuPhktIqXCEUrRNIlJWxRur.send'
const ServerChanPushMisc = ".send"

// API Area

// ApiUrl National Health Commission of PRC 's API address
const ApiUrl = "http://103.66.32.242:8005/zwfwMovePortal/interface/interfaceJson"

// API essential parma from 'http://bmfw.www.gov.cn/yqfxdjcx/risk.html'
const (
	ApiToken     = "23y0ufFl5YxIyGrI8hWRUZmKkvtSjLQA"
	ApiKey       = "3C502C97ABDA40D0A60FBEE50FAAD1DA"
	ApiNonce     = "123456789abcdefg"
	ApiPassid    = "zdww"
	ApiAppid     = "NcApplication"
	XWifNonce    = "QkjjtiLM2dCratiA"
	XWifPassid   = "smt-application"
	SignatureKey = "fTN2pfuisxTavbTuYVSsNJHetwq5bJvCQkjjtiLM2dCratiA"
)

// POST Headers
const (
	ContentType = "application/json; charset=utf-8"
	UserAgent   = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36 Edg/92.0.902.62"
)
