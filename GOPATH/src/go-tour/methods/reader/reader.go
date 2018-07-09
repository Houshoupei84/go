package main

import (
	"strings"
	"fmt"
	"io"
)

func main(){

	r := strings.NewReader("Hello, Reader")

	//每次8个字节的slice
	b := make([]byte, 8)

	for{

		//每次读8个字节，直到结束
		n, err := r.Read(b)

		//%v 打印的是默认格式 这里是整形值
		fmt.Printf("n = %v err= %v b = %v\n",
			n , err, b)

		//%q 打印字符值
		fmt.Printf("b[:n] = %q\n", b[:n])

		if err == io.EOF {
			break
		}
	}
}
