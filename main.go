package main

import "fmt"

func main() {
	//var ch chan string
	//go func(){
	//	select {
	//	case a := <- ch:
	//		fmt.Printf(a)
	//	}
	//}()

	//fmt.Println(time.Now())
	//str := "天天向上.Day.Day.Up!"
	//
	////for...range遍历字符串，获取后，下标i显示，每个中文字占三个下标
	//a := strings.Replace(str,".","_",-1)
	//fmt.Println(a)
	//fmt.Println("-------------------")
	//
	////for...range遍历所有字节，获取后，每个字节都是以十六进制表示
	//for i, ch := range []byte(str) {
	//	fmt.Printf("%d : %X \n", i, ch)
	//}
	//fmt.Println("-------------------")
	//
	////for...range遍历所有字符，下标i显示，每个中文和英文字符一样，都只占一个下标
	//for i, ch := range []rune(str) {
	//	fmt.Printf("%d : %c \n", i, ch)
	//}
	//fmt.Println("-------------------")
	//
	////获取字符串长度，中文字占3个长度。
	//fmt.Println("字符串长度",len(str))
	////获取字符串长度，使用utf8包里的函数，每个中文字只算一个长度。
	//fmt.Println("字符串长度",utf8.RuneCountInString(str))
	//fmt.Println(time.Now())
	err := errorTest()
	if err != nil {
		fmt.Println(err)
	}
}

func errorTest() interface{} {
	var err interface{}
	defer func() interface{} {
		if err = recover(); err != nil {
			fmt.Println(err, "111111111111111")
			return err
		}
		return err
	}()
	a := 0
	b := 1 / a
	return b
}
