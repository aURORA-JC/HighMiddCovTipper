package util

import (
	"HighMiddCovTipper/model"
	"HighMiddCovTipper/template"
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"
)

// GetTimestamp Get the timestamp
func GetTimestamp() int64 {
	return time.Now().Unix()
}

// GetSHA256String Get the string's key of sha256
func GetSHA256String(key string) string {
	byteSHA := sha256.Sum256([]byte(key))
	return strings.ToUpper(hex.EncodeToString(byteSHA[:]))
}

// FormatContent A creator of Markdown Content
func FormatContent(info model.Covid19Data) string {
	content := fmt.Sprintf(template.PushContentTemplate, info.Data.EndUpdateTime, info.Data.Hcount, info.Data.Mcount, tableCreator(info.Data.HighList), tableCreator(info.Data.MiddleList))
	return content
}

// tableCreator A creator of Markdown Table
func tableCreator(data []struct {
	Type        string   `json:"type"`
	Province    string   `json:"province"`
	City        string   `json:"city"`
	County      string   `json:"county"`
	AreaName    string   `json:"area_name"`
	Communities []string `json:"communitys"`
}) string {
	var table string
	for _, address := range data {
		var communities string
		for i, community := range address.Communities {
			if i != 0 {
				communities += "、"
			}
			communities += community
		}
		table += fmt.Sprintf(template.PushTableTemplate, address.Province, address.City, address.County, communities)
	}

	return table
}

// ReadLatestTime Read time from file
func ReadLatestTime() string {
	file, openError := os.Open("./latest.dat")
	if openError != nil {
		return ""
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\n')

	return line
}

// WriteLatestTime Write the latest time to file
func WriteLatestTime(time string) {
	file, openError := os.OpenFile("./latest.dat", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0662)
	if openError != nil {
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	_, writeError := file.WriteString(time)
	if writeError != nil {
		return
	}
}

// CompareTime Compare now & store to decide push or not
func CompareTime(nowTime string, storeTime string) bool {
	if nowTime == "" {
		return true
	}
	result := strings.Compare(nowTime, storeTime)
	if result == 0 {
		return false
	}

	return true
}
