package string_test

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	// s[1] = '3'//string是不可变的btye slice
	t.Log(len(s))
	// s = "\xE4\xBB\xA5"
	s = "\xE4\xBB\xA5\xFF"
	t.Log(s)
	t.Log(len(s))
	s = "中"
	t.Log(len(s)) //是btye数
	c := []rune(s)
	t.Log(len(c))
	t.Logf("中unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "你好啊"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
	}
}
