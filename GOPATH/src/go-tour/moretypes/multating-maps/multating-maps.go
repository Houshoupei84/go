package main

import "fmt"

func main(){

	m := make( map[string]int )

	m["Answer" ] = 42
	fmt.Println("The Value : ", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value : ", m["Answer"])

	//删除元素：
	delete(m , "Answer")
	fmt.Println("The value : ", m["Answer"])

	//如果 key 在 m 中，`ok` 为 true 。
	// 否则， ok 为 `false`，并且 elem 是 map 的元素类型的零值。
	v, ok := m["Answer"]
	fmt.Println("The value : ", v, "Present?", ok)

}
