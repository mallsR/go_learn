package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName 用一个名字调用 greetings.Hello,
// 检查有效的返回值.
func TestHelloName(t *testing.T) { // 测试函数会获取指针指向testing包的testing.T 类型作为参数。
	name := "Gladys"
	// 正则中的单词边界符。拼接后形成 \bGladys\b，确保匹配完整的单词 Gladys，避免匹配到包含该字符串的其他单词（例如防止匹配到 "NotGladys"）。
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty 使用空字符串调用 greetings.Hello,
// 检查错误.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
