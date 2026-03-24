package greetings

import (
	"errors"
	"fmt"
)

// Hello 为指定的人返回问候语.
func Hello(name string) (string, error) {
	// 如果没有给出名字，返回一个带有消息的错误.
	if name == "" {
		return "", errors.New("empty name")
	}

	// 返回在消息中嵌入名称的问候语.
	message := fmt.Sprintf("Hi, %v. Welcome!", name) // Sprintf会将格式化的问候文本返回给调用者.
	return message, nil                              // 这里的nil表示没有错误发生.
}
