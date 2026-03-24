package main

import (
	"fmt"
	"log"

	"go_learn/greetings"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	// 设置预定义Logger的属性，包括
	// 日志条目前缀和禁用打印的标志
	//  时间、源文件和行号.
	log.SetPrefix("greetings: ") // 这里的前缀"greetings: "将被添加到每条日志消息的开头.
	log.SetFlags(0)              // 这里的0表示禁用所有默认的日志标志.

	// message, err := greetings.Hello("5000")
	message, err := greetings.Hellos([]string{"5000", "xiaoR"})

	// 如果返回错误，则将其打印到控制台并
	// 退出程序.
	if err != nil {
		log.Fatal(err) // log.Fatal()会打印错误消息并调用os.Exit(1)，因此程序会在这里终止.
	}

	// 如果没有返回错误，则打印返回的消息
	// 到控制台.
	fmt.Println(message)
}
