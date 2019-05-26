package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// 上证综指结构体
type sz struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// 上证数据列表
var szData []sz

// 获取数据上证列表
func getSZ() {
	data, err := ioutil.ReadFile("./resource/sz.json")
	// 文件读取失败提醒
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	// 转化为对象
	err = json.Unmarshal(data, &szData)
	// 转换失败提醒
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
}

func main() {
	// 获取数据上证列表
	getSZ()
	fmt.Println(szData[0].Name)
}
