package main

import (
	"fmt"
	"strings"
)

func main() {

	name := "John Doe"
	for i := 0; i < len(name); i++ {
		fmt.Printf("%x ", name[i])
	}
	fmt.Println()

	h := "Hello, "
	w := "World!"

	var sb strings.Builder
	sb.WriteString(h)
	sb.WriteString(w)
	println(sb.String())

	println(name[0:4]) // John
	println(name[5:8]) // Doe

	var mySlice []string
	mySlice = append(mySlice, "Hello")
	mySlice = append(mySlice, "World")
	mySlice = append(mySlice, "Go")
	mySlice = append(mySlice, "is")
	mySlice = append(mySlice, "awesome")

	println(mySlice[1]) 
	println(mySlice[2]) 

	email := "user@example.com"
    url := "https://golang.org"
    filename := "document.pdf"
    
    // 使用 HasPrefix 检查字符串前缀
    fmt.Printf("Does email start with 'user'? %t\n", strings.HasPrefix(email, "user"))
    fmt.Printf("Does URL start with 'https'? %t\n", strings.HasPrefix(url, "https"))
    fmt.Printf("Does URL start with 'http'? %t\n", strings.HasPrefix(url, "http"))
    
    fmt.Printf("Does filename end with '.pdf'? %t\n", strings.HasSuffix(filename, ".pdf"))
    
    // 空字符串是任何字符串的前缀
    fmt.Printf("Does email start with ''? %t\n", strings.HasPrefix(email, ""))

	fmt.Println(strings.Contains(email, "example")) 
	fmt.Println(strings.Count(email, "@"))    
	fmt.Println(strings.Index(email, "@"))     // 4
	fmt.Println(strings.LastIndex(email, "m")) 

	text := "Go is good. Go is fast. Go is productive."

	// ReplaceAll 示例 - 总是替换所有匹配项
    fmt.Println("\n=== ReplaceAll 示例 ===")
    result4 := strings.ReplaceAll(text, "Go", "Golang")
    fmt.Println("替换所有:", result4)

	// string comparison
	fmt.Println(strings.Compare("abc", "abc")) // 0
	fmt.Println(strings.Compare("abc", "ab"))  // 1
	if "abc" < "abd" {
		fmt.Println("c < d")
	} else {
		fmt.Println("c >= d")
	}

	lowStr := "hello"
	upperStr := strings.ToUpper(lowStr)
	fmt.Println(upperStr) // HELLO
	upperStrings := "HELLO"
	lowerStr := strings.ToLower(upperStrings)
	fmt.Println(lowerStr) // hello

	


}