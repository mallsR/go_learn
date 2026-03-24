package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello 为指定的人返回问候语.
func Hello(name string) (string, error) {
	// 如果没有给出名字，返回一个带有消息的错误.
	fmt.Println("hello: ", name)
	if name == "" {
		return "", errors.New("empty name")
	}

	// 返回在消息中嵌入名称的问候语.
	message := fmt.Sprintf(randomFormat(), name) // Sprintf会将格式化的问候文本返回给调用者.
	return message, nil                          // 这里的nil表示没有错误发生.
}

// Hellos 返回一个map，该地图将每个已命名的人员
// 与问候消息相关联.
func Hellos(names []string) (map[string]string, error) { // 以下语法初始化map：make(map[key-type]value-type
	// 将名称与消息关联的map.
	messages := make(map[string]string)
	// 遍历接收到的名称切片，调用
	// Hello 函数为每个名字获取一条消息.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// 在map中，将检索到的消息与
		// 名称相关联.
		messages[name] = message
	}
	return messages, nil
}

// init 为函数中使用的变量设置初始值.
// 添加 init 函数以使用当前时间为rand包设定种子。
// Go 在初始化全局变量后，在程序启动时自动执行 init 函数。
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat 返回一组问候消息中的一个。返回的
// 消息是随机选择的.
func randomFormat() string {
	// 消息格式的切片.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// 通过为格式切片指定随机索引
	// 来返回随机选择的消息格式.
	return formats[rand.Intn(len(formats))]
}
