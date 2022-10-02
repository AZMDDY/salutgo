package practise

import (
	"fmt"
)

// 函数：普通函数，匿名函数
// struct中的是方法，方法可以重载

// 函数参数是传值的
// map, slice, interface, channel本身是指针类型
// 数组不是引用或指针
func func2(a []int) {
	a[0] = 0
}

func func3(a [2]int) {
	a[0] = 0
}

func func1(a int) (int, int) {
	return a, a + 1
}

func func4(args ...int) {
	fmt.Printf("args: %v\n", args)
}

// 函数类型
// 类型定义
type fun func(int, int) (int, int)

// 类型别名
type ffun = fun

func add(int) func(int) int {
	var x int
	f := func(y int) int {
		x += y
		return x
	}
	return f
}

// defer 延迟处理，逆序处理
// 在return之前才被执行
// 永远资源释放

func func5() {
	fmt.Printf("%s\n", "first")
	defer fmt.Printf("%s\n", "1")
	defer fmt.Printf("%s\n", "2")
	defer fmt.Printf("%s\n", "3")
	defer fmt.Printf("%s\n", "4")
}

// 先于main函数，每个包可以有多个init，同一个包的init函数执行顺序没有规定，不能依赖这个顺序
// 不用的包的init的依赖导入包的顺序

func testPtr() {
	// go 中不能对地址做偏移
	var ptr *int = nil
	fmt.Printf("ptr: %v\n", ptr)

}

type Person struct {
	Name string
	Age  uint8
	id   uint32
}

// 值类型会拷贝
func (p Person) GetName() string {
	return p.Name
}

func testStruct() {
	var tom Person
	tom.Age = 12
	tom.Name = "tom"
	tom.id = 10
	fmt.Printf("tom: %v\n", tom)
	jake := Person{Name: "jake", Age: 12, id: 333}
	fmt.Printf("jake: %v\n", jake)

	type Class struct {
		id uint32
	}

}

type Inter interface {
	read()
	write()
}

type Cpr struct {
}

func (c *Cpr) read()  {}
func (c *Cpr) write() {}

type Crr struct {
}

func (c *Crr) read() {}

func testInter(i Inter) {
	switch t := i.(type) {
	case *Cpr:
		fmt.Printf("cpr: %T\n", t)
	default:
	}
}

// 变量初始化 > init > main
func init() {
	fmt.Println(func1(3))
	a := []int{2, 2, 2}
	func2(a)
	fmt.Printf("a: %v\n", a)
	b := [2]int{1, 2}
	func3(b)
	fmt.Printf("b: %v\n", b)
	func4(1, 2, 3)
	ad := add(10)
	fmt.Printf("ad(10): %v\n", ad(10))
	fmt.Printf("ad(20): %v\n", ad(20))
	func5()
	testPtr()
	var c Cpr
	c.read()
	fmt.Printf("c: %T\n", c)
	testStruct()
}
