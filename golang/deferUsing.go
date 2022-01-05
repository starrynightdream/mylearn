package main
// testing defer use

import (
	"fmt"
)
// 匿名返回情况
func a() int{
	var i int = 0
	defer func() {
		i++
		fmt.Println("a func first defer \t call i = ", i)
	}()
	defer func() {
		i++
		fmt.Println("a func second defer \t call i = ", i)
	}()
	return i // 匿名返回值在 return 执行时才创建
	// 返回的不是i，是将i赋给了返回值
}

// 具名返回情况
func b() (i int) { // 具名返回值在函数开始便创建
	defer func() {
		i++
		fmt.Println("b func first defer \t call i = ", i)
	}()
	defer func() {
		i++
		fmt.Println("b func second defer \t call i = ", i)
	}()
	return i // 返回i
	// i可以不写 直接写 return
}

// 证明
func a_to_pos() *int {
	var i int = 0
	defer func() {
		i++
		fmt.Println("a func first defer \t call i = ", i)
	}()
	defer func() {
		i++
		fmt.Println("a func second defer \t call i = ", i)
	}()
	return &i
}
// 可以发现原来的变量确实被修改为了2，但直接返回值则是0

/* 
结论
1. defer是堆栈的，后进先执行
2. return 会将创建返回值，但并未结束函数，会检查是否有defer语句。执行defer之后才根据 RET指令返回值
 */
func main() {
	// fmt.Println("a func return num \t call i = ", a())
	// fmt.Println("b func return num \t call i = ", b())
	fmt.Printf("a func retrn num \t call i =  %d\n", *a_to_pos())
}
// 结果
// a func second defer      call i =  1
// a func first defer       call i =  2
// a func return num        call i =  0
// b func second defer      call i =  1
// b func first defer       call i =  2
// b func return num        call i =  2