//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("Hello, Go!")
//}

//package main

//import "fmt"
//func main() {
//	i := 1
//	for { // 相当于c中的while循环
//		fmt.Println("loop")
//		break // 跳出循环
//	}
//	// 打印7、8
//	for j := 7; j < 9; j++ {
//		fmt.Println(j)
//	}
//	for n := 0; n < 5; n++ {
//		if n%2 == 0 {
//			continue
//			// 当n模2为0时不打印，进到下一次的循环
//		}
//		fmt.Println(n)
//	}
//	// 直到i>3
//	for i <= 3 {
//		fmt.Println(i)
//		i = i + 1
//	}
//	// for 循环嵌套
//	for i := 0; i < 5; i++ {
//		for j := 0; j < 5; j++ {
//			fmt.Printf("i = %d, j = %d\n", i, j)
//		}
//	}
//}
//
//package main
//
//import "fmt"
//
//func main() {
//	// 条件表达式为false，打印出"7 is odd"
//	if 7%2 == 0 {
//		fmt.Println("7 is even")
//	} else {
//		fmt.Println("7 is odd")
//	}
//	// 条件表达式为ture，打印出"8 is divisible by 4"
//	if 8%4 == 0 {
//		fmt.Println("8 is divisible by 4")
//	}
//	// 短声明，效果等效于
//	//num := 9
//	//if num < 0{
//	// ...
//	//}
//	if num := 9; num < 0 {
//		fmt.Println(num, "is negative")
//	} else if num < 10 {
//		fmt.Println(num, "has 1 digit")
//	} else {
//		fmt.Println(num, "has multiple digits")
//	}
//}

//package main

//func main() {
//	a := 2
//	switch a {
//	case 1:
//		fmt.Println("one")
//	case 2:
//		// 在此打印"two"并跳出
//		fmt.Println("two")
//	case 3:
//		fmt.Println("three")
//	case 4, 5:
//		fmt.Println("four or five")
//	default:
//		fmt.Println("other")
//	}
//	t := time.Now()
//	switch {
//	// 根据现在的时间判断是上午还是下午
//	case t.Hour() < 12:
//		fmt.Println("It's before noon")
//	default:
//		fmt.Println("It's after noon")
//	}
//}

// package main
//
// import "fmt"
//
//	func main() {
//		// 声明了长度为5的数组，数组中的每一个元素都是int类型
//		var a [5]int
//		// 给数组a的第4位元素赋值为100
//		a[4] = 100
//		fmt.Println("get:", a[2])
//		fmt.Println("len:", len(a))
//		// 在给数组声明的同时赋值
//		b := [5]int{1, 2, 3, 4, 5}
//		fmt.Println(b)
//		// 声明二维数组
//		var twoD [2][3]int
//		for i := 0; i < 2; i++ {
//			for j := 0; j < 3; j++ {
//				twoD[i][j] = i + j
//			}
//		}
//		fmt.Println("2d: ", twoD)
//	}
//package main

//func main() {
//	s := make([]string, 3)
//	s[0] = "a"
//	s[1] = "b"
//	s[2] = "c"
//	fmt.Println("get:", s[2])   // c
//	fmt.Println("len:", len(s)) // 3
//	// 使用append在尾部添加元素
//	s = append(s, "d")
//	s = append(s, "e", "fg")
//	fmt.Println(s) // [a b c d e f]
//	c := make([]string, len(s))
//	// 将s复制给c
//	copy(c, s)
//	fmt.Println(c)      // [a b c d e f]
//	fmt.Println(s[2:5]) // [c d e]
//	fmt.Println(s[:5])  // [a b c d e]
//	fmt.Println(s[2:])  // [c d e f]
//	good := []string{"g", "o", "o", "d"}
//	fmt.Println(good) // [g o o d]
//}

//package main
//
//import "fmt"
//
//func main() {
//	nums := []int{2, 3, 4}
//	sum := 0
//	for i, num := range nums {
//		sum += num
//		if num == 2 {
//			fmt.Println("index:", i, "num:", num) // index: 0 num: 2
//		}
//	}
//	fmt.Println(sum) // 9
//	m := map[string]string{"a": "A", "b": "B", "c": "C"}
//	for k, v := range m {
//		fmt.Println(k, v)
//	}
//	for k := range m {
//		fmt.Println("key", k)
//	}
//}

//package main
//
//import "fmt"
//
//func add(a int, b int) int {
//	// 返回a+b的和
//	return a + b
//}
//
//// 若类型相同，允许这样写
//func add2(a, b int) int {
//	return a + b
//}
//func exists(m map[string]string, k string) (v string, ok bool) {
//	v, ok = m[k]
//	return v, ok
//}
//func main() {
//	res := add(1, 2)
//	fmt.Println(res) // 3
//	v, ok := exists(map[string]string{"a": "A"}, "a")
//	fmt.Println(v, ok) // A True
//	v, ok = exists(map[string]string{"a": "A"}, "b")
//	fmt.Println(v, ok) // false
//}

//package main
//
//import "fmt"
//
//func add2(n int) {
//	n += 2
//}
//func add2pt(n *int) {
//	*n += 2
//}
//func swap(a int, b int) {
//	a, b = b, a
//}
//
//// 函数参数为指针类型
//func swapWithPt(a *int, b *int) {
//	*a, *b = *b, *a
//}
//func main() {
//	n := 5
//	add2(n)
//	fmt.Println(n) // 5
//	add2pt(&n)
//	fmt.Println(n) // 7
//	a, b := 2, 3
//	swap(a, b)
//	fmt.Println(a, b)
//	swapWithPt(&a, &b)
//	fmt.Println(a, b)
//}

//package main
//
//import "fmt"
//
//type point struct {
//	x, y int
//}
//
//func main() {
//	s := "hello"
//	n := 123
//	p := point{1, 2}
//	fmt.Println(s, n)        // hello 123
//	fmt.Println(p)           // {1 2}
//	fmt.Printf("s=%v\n", s)  // s=hello
//	fmt.Printf("n=%v\n", n)  // n=123
//	fmt.Printf("p=%v\n", p)  // p={1 2}
//	fmt.Printf("p=%+v\n", p) // p={x:1 y:2}
//	fmt.Printf("p=%#v\n", p) // p=main.point{x:1, y:2}
//	f := 3.141592653
//	fmt.Println(f)          // 3.141592653
//	fmt.Printf("%.2f\n", f) // 3.14
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	now := time.Now()
//	fmt.Println(now) // 2022-03-27 18:04:59.433297 +0800 CST m=+0.000087933
//	t := time.Date(2022, 3, 27, 1, 25, 36, 0, time.UTC)
//	t2 := time.Date(2022, 3, 27, 2, 30, 36, 0, time.UTC)
//	fmt.Println(t)                                                  // 2022-03-27 01:25:36 +
//	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute()) // 2022 March 27 1 25
//	fmt.Println(t.Format("2006-01-02 15:04:05"))                    // 2022-03-27 01:25:36
//	diff := t2.Sub(t)
//	fmt.Println(diff)                           // 1h5m0s
//	fmt.Println(diff.Minutes(), diff.Seconds()) // 65 3900
//	t3, err := time.Parse("2006-01-02 15:04:05", "2022-03-27 01:25:36")
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(t3 == t)    // true
//	fmt.Println(now.Unix()) // 1648738080
//}

package lesson

//import (
//	"fmt"
//	"strconv"
//)
//
//func main() {
//	f, _ := strconv.ParseFloat("1.234", 64)
//	fmt.Println(f) // 1.234
//	n, _ := strconv.ParseInt("111", 10, 64)
//	fmt.Println(n) // 111
//	n, _ = strconv.ParseInt("0x1000", 0, 64)
//	fmt.Println(n) // 4096
//	n2, _ := strconv.Atoi("123")
//	fmt.Println(n2) // 123
//	n2, err := strconv.Atoi("AAA")
//	fmt.Println(n2, err) // 0 strconv.Atoi: parsing "AAA": invalid syntax
//}
