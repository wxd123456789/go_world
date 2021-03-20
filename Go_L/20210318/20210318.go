package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	//Print系列函数会将内容输出到系统的标准输出，区别在于Print函数直接输出内容，Printf函数支持格式化输出字符串，Println函数会在输出内容的结尾添加一个换行符。
	/*	fmt.Print("在终端打印该信息。")
		name := "枯藤"
		fmt.Printf("我是：%s\n", name)
		fmt.Println("在终端打印单独一行显示")
		//Sprint系列函数会把传入的数据生成并返回一个字符串。
		s2 := fmt.Sprintf("name:%s,age:%d", name, 12)
		fmt.Printf(s2)
		//Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。
		err := fmt.Errorf("这是一个错误")
		_ = err

		now := time.Now()
		fmt.Printf("current time:%v\n", now)
		year := now.Year()     //年
		month := now.Month()   //月
		day := now.Day()       //日
		hour := now.Hour()     //小时
		minute := now.Minute() //分钟
		second := now.Second() //秒
		fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
		timestamp1 := now.Unix()     //时间戳
		timestamp2 := now.UnixNano() //纳秒时间戳
		fmt.Printf("current timestamp1:%v\n", timestamp1)
		fmt.Printf("current timestamp2:%v\n", timestamp2)
		later := now.Add(time.Hour)
		fmt.Println(later)*/
	/*	ticker := time.Tick(time.Second)
		for i := range ticker {
			fmt.Println(i)//每秒都会执行的任务
		}*/

	//os.Args是一个[]string
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
	//定义命令行参数方式1
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	//解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
	// $ ./flag_demo -name pprof --age 28 -married=false -d=1h30m

}
func testJson() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}
func init() {
	/*	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("open log file failed, err:", err)
			return
		}
		log.SetOutput(logFile)
		log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
		log.Println("这是自定义的logger记录的日志。")*/
}

var BASEURL = "http://127.0.0.1:8080"

func httpServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}

func httpGetReq() {
	apiUrl := fmt.Sprintf("%s/get", BASEURL)
	// URL param
	data := url.Values{}
	data.Set("name", "枯藤")
	data.Set("age", "18")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed,err:%v\n", err)
	}
	u.RawQuery = data.Encode() // URL encode
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}

func httpPostReq() {
	url := fmt.Sprintf("%s/post", BASEURL)
	contentType := "application/json"
	data := `{"name":"枯藤","age":18}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
