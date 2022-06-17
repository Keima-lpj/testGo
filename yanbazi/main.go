package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	fmt.Println("请输入：1 从C盘copy到当前目录 2 从当前目录copy到C盘")
	var flag int
	fmt.Scanf("%d", &flag)
	if flag == 1 {
		CCopyToCur()
	} else {
		curCopyToC()
	}
	//time.Sleep(time.Second * 200)
}

func curCopyToC() {
	pwd, _ := os.Getwd()
	pwd = pwd + "/Save_1room"
	fmt.Println(pwd)
	var user string
	user = os.Getenv("USERPROFILE")
	m := user + "/AppData/LocalLow/cultparthia/1room/Utage"
	m = "C:/Users/Zt_01/AppData/LocalLow/cultparthia/1room/Utage/Save_1room"
	copyDir(pwd, m)
}

func CCopyToCur() {
	var user string
	user = os.Getenv("USERPROFILE")
	m := user + "/AppData/LocalLow/cultparthia/1room/Utage"
	m = "C:/Users/Zt_01/AppData/LocalLow/cultparthia/1room/Utage"
	//当前目录
	pwd, _ := os.Getwd()
	//进入存档系统目录下
	copyDir(m, pwd)
}

func FormatPath(s string) string {
	switch runtime.GOOS {
	case "windows":
		return strings.Replace(s, "/", "\\", -1)
	case "darwin", "linux":
		return strings.Replace(s, "\\", "/", -1)
	default:
		return s
	}
}

func copyDir(src string, dest string) {
	src = FormatPath(src)
	dest = FormatPath(dest)
	log.Println(src)
	log.Println(dest)

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("xcopy", src, dest, "/I", "/E")
	case "darwin", "linux":
		cmd = exec.Command("cp", "-R", src, dest)
	}

	outPut, e := cmd.Output()
	if e != nil {
		return
	}
	fmt.Println(string(outPut))
}
