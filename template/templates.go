package template

// PushContentTemplate A Markdown Format Text Template
const PushContentTemplate = "## 全国中高风险疫情地区\n\n" +
	"> **数据来源：国家卫健委 - %s 更新**\n\n" +
	"**免责声明**：此脚本数据来自国家卫健委，详情查看下方链接：\n\n" +
	"[中高风险等级地区-便民服务-中国政府网](http://bmfw.www.gov.cn/yqfxdjcx/risk.html)\n\n" +
	"-----\n\n" +
	"### **高风险地区**：**%d** 个、**中风险地区**：**%d** 个\n\n" +
	"### 高风险地区列表\n\n" +
	"|  省份  |  城市  |  城区  | 街道、社区 |\n" +
	"| :----: | :----: | :----: | :--------------: |\n" +
	"%s\n\n" +
	"### 中风险地区列表\n\n" +
	"|  省份  |  城市  |  城区  | 街道、社区 |\n" +
	"| :----: | :----: | :----: | :--------------: |\n" +
	"%s\n\n\n" +
	"注：其余未列出地区均为低风险\n\n" +
	"疫情尚未结束，请做好必要的个人防护！\n\n" +
	"更多疫情防控信息：请关注各地卫健委相关发布。\n\n" +
	"[新型冠状病毒肺炎疫情防控 - 国家卫健委](http://www.nhc.gov.cn/xcs/xxgzbd/gzbd_index.shtml)"

// TableTemplate A Markdown Table's Template
const TableTemplate = "| %s | %s | %s | %s |\n"