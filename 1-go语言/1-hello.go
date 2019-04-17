package main    // 定义包名，

import "fmt"  // 引入fmt包，fmt包实现了格式化IO(输入/输出)的函数。

func main() {   // 程序开始执行的函数。main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）

	// fmt.Println("Hello ,World!")

	// var i int
	// var f float64
	// var b bool
	// var s string
	// fmt.Printf("%v %v %v %q\n", i, f, b, s) 

	// var intVal int
	// // intVal :=1 // 这时候会产生编译错误
	// intVal,intVal1 := 1,2   // 此时不会产生编译错误，因为有声明新的变量，因为:=是一个声明语句
	// fmt.Print( intVal, intVal1)

	var x,y int
	var (  //这种因式分解关键字的写法一般用于声明全局变量
		a int
		b bool
	)

	var c, d int = 1, 2
	var e, f = 123, "hello"

	// 这种不带声明格式的只能在函数体中出现
	g,h := 123,"hello"

	fmt.Println(x, y, a, b, c, d, e, f, g, h)
}