package main

import (
	"HighMiddConvTipper/service"
	"HighMiddConvTipper/util"
)

func main() {
	data := service.GetConvData()
	if data.Code == 114514 {
		return
	}

	if util.CompareTime(data.Data.EndUpdateTime, util.ReadLatestTime()) {
		content := util.FormatContent(data)
		service.SendConvInfo2Users(content)
		util.WriteLatestTime(data.Data.EndUpdateTime)
	}
}



