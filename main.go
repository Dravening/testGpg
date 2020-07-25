package main

import (
	"context"
	"errors"
	"log"
	"os/exec"
	"time"
)

//func main() {
//	//var ch chan string
//	//go func(){
//	//	select {
//	//	case a := <- ch:
//	//		fmt.Printf(a)
//	//	}
//	//}()
//
//	//fmt.Println(time.Now())
//	//str := "天天向上.Day.Day.Up!"
//	//
//	////for...range遍历字符串，获取后，下标i显示，每个中文字占三个下标
//	//a := strings.Replace(str,".","_",-1)
//	//fmt.Println(a)
//	//fmt.Println("-------------------")
//	//
//	////for...range遍历所有字节，获取后，每个字节都是以十六进制表示
//	//for i, ch := range []byte(str) {
//	//	fmt.Printf("%d : %X \n", i, ch)
//	//}
//	//fmt.Println("-------------------")
//	//
//	////for...range遍历所有字符，下标i显示，每个中文和英文字符一样，都只占一个下标
//	//for i, ch := range []rune(str) {
//	//	fmt.Printf("%d : %c \n", i, ch)
//	//}
//	//fmt.Println("-------------------")
//	//
//	////获取字符串长度，中文字占3个长度。
//	//fmt.Println("字符串长度",len(str))
//	////获取字符串长度，使用utf8包里的函数，每个中文字只算一个长度。
//	//fmt.Println("字符串长度",utf8.RuneCountInString(str))
//	//fmt.Println(time.Now())
//	err := errorTest()
//	if err != nil {
//		fmt.Println(err)
//	}
//}

//设置超时时间为 5秒
var Timeout = 5 * time.Second

//执行命令并添加超时检测
func Command(name string, arg ...string) (string, error) {
	ctxt, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel() //releases resources if slowOperation completes before timeout elapses
	cmd := exec.CommandContext(ctxt, name, arg...)
	//当经过Timeout时间后，程序依然没有运行完，则会杀掉进程，ctxt也会有err信息
	if out, err := cmd.Output(); err != nil {
		//检测报错是否是因为超时引起的
		if ctxt.Err() != nil && ctxt.Err() == context.DeadlineExceeded {
			return "", errors.New("command timeout")
		}
		return string(out), err
	} else {
		return string(out), nil
	}
}
func main() {
	//设置超时为5秒 sleep 6秒肯定会超时
	out, err := Command("sleep", "6")
	if err != nil {
		//经验证，没毛病
		log.Println("错误返回:", err.Error())
	} else {
		log.Println("正常返回:", out)
	}
}
