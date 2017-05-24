package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"strconv"
)

type Vertex struct {
	X int
	Y int
}
var (
	v1 = Vertex{1, 2}  // 类型为 Vertex
	v2 = Vertex{X: 1}  // Y:0 被省略
	v3 = Vertex{}      // X:0 和 Y:0
	p3  = &Vertex{1, 2} // 类型为 *Vertex
)

func main() {

	fmt.Println(v1, p3, v2, v3)

	v := Vertex{1, 2}
	p1 := &v
	p1.X = 1e9
	fmt.Println(v)

	fmt.Println("hello,world!你好，世界！")
	fmt.Println(math.Pi)
	fmt.Println(rand.Intn(99))
	fmt.Println(add(11, 2))
	fmt.Println(swap("aa", "bb"))
	fmt.Println(split(2))
	fmt.Println(isb())
	fmt.Println(kzl(10))
	fmt.Println(strconv.Itoa(100))
	convs()

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
	i, j := 42, 2701
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

//:= 在明确类型的情况下可代替 var 定义  但不能在函数外使用
func add(x int, y int) int {
	r := x + y
	return r
}

//多个返回值
func swap(x, y string) (string, string) {
	x = "this is " + x
	y = "\nthis is " + y
	return x, y
}

//返回值命名 return 没有参数  称作 裸返回
func split(sum int) (x, y int) {
	x = sum * 10
	y = sum * 100
	return
}

//声明变量
var i, str, bo = 1, "string", false
var (
	j int = 2
	s     = "str"
	b     = true
)

func isb() (int, string, bool) {
	sum := i + j
	return sum, str + s, bo || b
}

//控制流
func kzl(k int) string {
	str := "count down :"
	for i := k; i >= 0; i-- {
		str = str + strconv.Itoa(i) + "\n"
	}
	return str
}

//类型转换
func convs() {
	str1 := "1234"
	a1, error := strconv.Atoi(str1)
	if error != nil {
		fmt.Println("error")
	} else {
		fmt.Println(a1)
	}
	a2, err := strconv.ParseInt("123", 10, 64)
	if err == nil {
		fmt.Println(a2)
	}
	str2 := strconv.FormatInt(32456, 10)
	fmt.Println(str2)

	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	//while
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)
}
