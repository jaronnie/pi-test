package main

import (
	"bufio"
	"embed"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

//go:embed pi.txt
var f embed.FS

func main() {
	//FormatFile()
	//格式化网上下载的文件

	var input string
	fmt.Println("请输入生日，(格式为20000619)")
	_, _ = fmt.Scan(&input, "%s")
	bytes, _ := f.ReadFile("pi.txt")
	b, index := ContainsSubstring(string(bytes), input)
	if b {
		fmt.Println(input, "存在")
		fmt.Println("位置", index)
	} else {
		fmt.Println("不存在")
	}
	time.Sleep(time.Second *15)
}

func FormatFile() {
	f, _ := os.Open("pi-download.txt")
	fs, _ := os.Create("pi.txt")
	var str string
	defer f.Close()
	defer fs.Close()
	reader := bufio.NewReader(f)
	writer := bufio.NewWriter(fs)
	var n int
	for {
		n += 1
		bytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if len(bytes) < 50 {
			continue
		}
		str = strings.Replace(string(bytes[:54]), "\r\n", "", -1)
		str = strings.Replace(str, " ", "", -1)
		_, _ = writer.Write([]byte(str))
		fmt.Println("正在处理第", n, "行")

	}
	_ = writer.Flush()
}

func ContainsSubstring(s string, date string) (bool, int) {
	var index int
	index = strings.Index(s, date)
	if index > 0 {
		return true, index
	}
	return false, -1
}