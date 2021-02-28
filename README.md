# Π中是否一定存在你的生日？

有这个一个有趣的问题：Π中是否一定会存在你的生日？

如我的生日是2000年6月19号，那么Π的小数点中一定会存在连续的20000619。

为了去证实这个事清，我去查找了一些资料。

> 2019年3月14日，也是“圆周率日”，Google 宣布将π的计算精确到了小数点后31.4万亿位。
>
> 将圆周率计算到数十万亿位，需要数百 TB 的储存空间。

由于机器有限，时间和精力有限，所以我只选取了一亿位数字进行测试。存储空间需要100Mb左右。

在网上找到的资料长这个样子。

![image-20210228145122806](http://picture.nj-jay.com/image-20210228145122806.png)

所以我必须对其进行处理。

* 去掉末尾的 :50等字样
* 去掉所有的空格和换行符

我使用Go语言对其进行处理，由于文件很大，所以使用缓存的方式（一行行读取）进行读取，处理之后写入新的文件。

经计算，处理一亿位大概需要几分钟的时间。

```go
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
		if len(bytes) < 3 {
			continue
		}
		str = strings.Replace(string(bytes[:54]), "\r\n", "", -1)
		str = strings.Replace(str, " ", "", -1)
		_, _ = writer.Write([]byte(str))
		fmt.Println("正在处理第", n, "行")
	}
	_ = writer.Flush()
    fmt.Println("处理完毕")
}
```

处理完之后长这个样子。

![image-20210228145939611](http://picture.nj-jay.com/image-20210228145939611.png)

接下来我们读取这个文件，然后判断是否存在字符串"20000619"

```go
func ContainsSubstring(s string, date string) (bool, int) {
	var index int
	index = strings.Index(s, date)
	if index > 0 {
		return true, index
	}
	return false, -1
}
```

```go
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
```

![image-20210228162043030](http://picture.nj-jay.com/image-20210228162043030.png)

卧槽，我惊呆了，竟然真的存在。

当然我只是测试了这一个数据，如果你想测试你的生日的话，可以在我的github主页上下载已经编译好的程序进行测试。

github链接：https://github.com/gocloudcoder/pi-test

> 我使用了go1.16的新特性将静态文件打包到二进制程序当中
>
> 所以下载完之后直接运行就可以了
>
> 下载链接：https://picture.nj-jay.com/pi-test.exe

如果不存在，可以按照我的方法选取更多位数进行测试。

我相信这个事清是真的存在，确实很神奇。

Π还有很多奥秘等待我们去发现。

