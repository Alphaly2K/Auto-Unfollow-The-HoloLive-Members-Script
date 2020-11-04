package main

import (
	"Unfollow/Unfo"
	"fmt"
	"os"
	"os/exec"
)
var sessdata string
func main() {
	fmt.Println("请从cookies中获取sessdata和bili_jct信息")
	fmt.Println("输入回车以继续")
	fmt.Scanln()
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println("欢迎来到hololive一键取关脚本")
	fmt.Println("Version:1.1 preview * Unfollow MinatoAqua ver.")
	fmt.Println("Yagoo去死，COVER倒闭！")
	Unfo.Unfo()
	for {
		if Unfo.Unfo() == "OK" {
			os.Exit(0)
		}
	}
}