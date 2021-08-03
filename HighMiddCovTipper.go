package main

import (
	"HighMiddCovTipper/config"
	"HighMiddCovTipper/service"
	"HighMiddCovTipper/util"
	"fmt"
	"os"
)

func main() {
	// Check start args
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./HighMiddCovTipper <YOUR_SERVERCHAN_TOKEN>")
		return
	}

	// Set ServerChen token
	config.ServerChanPushToken = os.Args[1]

	// Get Data from NHC's API
	data := service.GetCovData()
	if data.Code == 114514 {
		return
	}

	// Check & Send to ServerChan's API
	if util.CompareTime(data.Data.EndUpdateTime, util.ReadLatestTime()) {
		content := util.FormatContent(data)
		service.SendCovInfo2Users(content)
		util.WriteLatestTime(data.Data.EndUpdateTime)
	}
}
