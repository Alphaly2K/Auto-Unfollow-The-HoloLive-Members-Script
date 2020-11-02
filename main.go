package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	fmt.Println("欢迎来到hololive一键取关脚本")
	fmt.Println("Version:1.0-final")
	fmt.Println("Yagoo去死，COVER倒闭！")
	var sessdata string
loop1:
	for
	{
		fmt.Println("输入您的SESSDATA")
		fmt.Scanln(&sessdata)
	loop:
		fmt.Println("您的sessdata为：", sessdata, "确定吗（yes/no）")
		var confirm string
		fmt.Scanln(&confirm)
		switch confirm {
		case "yes":
			break loop1
		case "no":
			continue
		default:
			goto loop
		}
	}
	sessdatalen := strings.Count(sessdata, "") - 1
	for {
		if sessdatalen != 32 {
			fmt.Println("SESSDATA输入错误")
			goto loop1
		} else {
			fmt.Println("查询中....")
			break
		}
	}
	client := &http.Client{}
	request, _ := http.NewRequest("GET", "https://api.bilibili.com/x/space/myinfo", nil)
	request.Header.Set("User-Agent", "FUCKCOCO/321.23(fuckcoco@fuckholo.com)")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Cookie", "SESSDATA="+sessdata)
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		namevalue := gjson.Get(bodystr, "data.name")
		codevalue := gjson.Get(bodystr, "code")
		if codevalue.String() != "0" {
			fmt.Println("未知错误,请重新检查您的SESSDATA")
			goto loop1
		}
	loop2:
		for
		{
			fmt.Println("您的昵称为：", namevalue, "对吗?（yes/no）")
			var confirm string
			fmt.Scanln(&confirm)
			switch confirm {
			case "yes":
				break loop2
			case "no":
				goto loop1
			default:
				continue
			}
		}
		var csrf string
	loop3:
		for
		{
			fmt.Println("请输入您的bili_jct信息：")
			fmt.Scanln(&csrf)
			var csrflen = strings.Count(csrf, "") - 1
			if csrflen != 32 {
				fmt.Println("输入错误")
				goto loop3
			} else {
				fmt.Println("工作中...")
				break loop3
			}
		}
		var uids = []string{"286700005", "286179206", "20813493", "366690056", "9034870", "3327041171", "336731767", "389856447", "339567211", "389857131", "375504219", "389857640", "389858027", "389858754", "389859190", "389862071", "412135222", "412135619", "443305053", "443300418", "454737600", "454733056", "454955503", "491474048", "491474049", "491474050", "491474051", "491474052", "624252710", "624252706", "624252712", "624252709", "389056211", "674600645", "674600646", "674600647", "674600648", "674600649", "624252711"}
		for i := 0; i < len(uids); i++ {
			uid := uids[i]
			body1 := "fid=" + uid + ";act=2;re_src=11;csrf=" + csrf
			body2 := strings.NewReader(body1)
			client2 := &http.Client{}
			request2, _ := http.NewRequest("POST", "http://api.bilibili.com/x/relation/modify", io.Reader(body2))
			request2.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			request2.Header.Set("user-agent", "FUCKCOCO/123.123(FUCKCOCO@ANTIHOLO.ORG)")
			request2.Header.Set("Cookie", "SESSDATA="+sessdata)
			response2, _ := client2.Do(request2)
			request2.Body.Close()
			if response2.StatusCode == 200 {
				body1, _ := ioutil.ReadAll(response2.Body)
				bodystr1 := string(body1)
				codevalue1 := gjson.Get(bodystr1, "code")
				if codevalue1.String() != "0" {
					fmt.Println("请检查您的bili_jct信息，然后重试")
					goto loop3
				}
				fmt.Println(bodystr1)
			}
		}
		fmt.Println("取关完毕！")
		fmt.Println("HOLOLIVE,FUCK U AND NEVER COME BACK!")
		fmt.Scanln()
	} else {
		goto loop1
	}
}
