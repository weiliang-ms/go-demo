package main

import "fmt"

func main() {
	test := []int{1, 2, 3}
	fmt.Printf("修改前：%v\n", test)
	modify(test)
	fmt.Printf("修改后：%v\n", test)
	str()
}

func modify(s []int) {
	s[0] = 9
}

// string底层是[]byte，但无法修改内容
func str() {
	str1 := "hello world"
	fmt.Printf("截取str1：%s第6个至结尾字符为：%s\n", str1, str1[6:])

	/**
	[]byte 按字节处理，一个汉字为3个字节
	[]rune 按字符处理，可以处理汉字
	*/

	// 修改英文字符串（字符为字母或字母与数字结合体）
	byte1 := []byte(str1)
	byte1[0] = 'H'
	fmt.Printf("修改字符前：\t%s\t修改字符后:\t%s\n", str1, string(byte1))

	// 修改中文字符
	str2 := "hello 明"
	byte2 := []rune(str2)
	byte2[6] = '红'
	fmt.Printf("修改字符前：\t%s\t修改字符后:\t%s\n", str2, string(byte2))
}
