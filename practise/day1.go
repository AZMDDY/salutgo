package practise

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"unsafe"
)

type Student struct {
	Name string
	Age  uint32
}

const MAX_NUM uint32 = 6666

func init() {
	name := "tom" // 只能函数内部声明，短变量声明
	fmt.Printf("name: %v\n", name)
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
	var (
		age uint32 = 7
	)
	fmt.Printf("age: %v\n", age)
	name, age = getName()
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	const (
		a1 = iota
		a2 = iota
	)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)

	const (
		a3 = iota
		a5 = 100
		a4 = iota
	)
	fmt.Printf("a3: %v\n", a3)
	fmt.Printf("a4: %v\n", a4)
	fmt.Printf("a3 type: %T\n", a3)
	ptr := &age
	fmt.Printf("ptr type: %T\n", ptr)
	b := [2]int{1, 2}
	fmt.Printf("b: %v\n", b)
	fmt.Printf("b type: %T\n", b)
	fmt.Printf("func: %T\n", getName)

	var b1 bool = false
	fmt.Printf("b1: %v\n", b1)
	testStr()
	testSwitch()
	testFor()
	testArray()
	testSlice()
	testMap()
}

func testStr() {
	s1 := `rtert
	ertert
	e4tertert`
	fmt.Printf("s1: %v\n", s1)
	s1 = "tom"
	s2 := "20"
	msg := s1 + s2
	fmt.Printf("msg: %v\n", msg)
	// 拼接
	msg = fmt.Sprintf("%s, %s\n", s1, s2)
	fmt.Printf("msg: %v\n", msg)

	// json拼接
	fmt.Printf("strings.Join([]string{s1, s2}, \",\"): %v\n", strings.Join([]string{s1, s2}, ","))

	var buffer bytes.Buffer
	buffer.WriteString("test")
	buffer.WriteString(": tttt")
	fmt.Printf("buffer: %v\n", buffer.String())
	str := "hello day1"
	fmt.Printf("str[:2]: %v\n", str[:2])

	tom := Student{Name: "tom", Age: 12}
	fmt.Printf("tom: %v\n", tom)
	fmt.Println("%%")
}

func testSwitch() {
	g := 10
	switch g {
	case 10:
		fmt.Printf("g: %v\n", g)
	case 20:
		fmt.Printf("g: %v\n", g)
	}
}

func testFor() {
	i := 1
	for i <= 10 {
		i++
	}

	var a = []int{1, 2, 3, 4}
	for idx, v := range a {
		fmt.Printf("idx: %v, v: %v\n", idx, v)
	}
	m := make(map[string]string, 0)
	m["tom"] = "111"
	for key, val := range m {
		fmt.Printf("key: %v\n", key)
		fmt.Printf("val: %v\n", val)
		for idx, v := range a {
			fmt.Printf("idx: %v, v: %v\n", idx, v)
			goto M
		}
	}
M:
	fmt.Printf("m: %v\n", m)
}

func testArray() {
	var a1 = [...]int{1, 2, 4}
	fmt.Printf("len(a1): %v\n", len(a1))
}

func testSlice() {
	var a []string
	var b []int
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	var c = make([]int, 10)
	fmt.Printf("c: %v\n", c)

	var s1 = []int{1, 2, 3, 4, 5, 6} // 切片初始化
	fmt.Printf("s1: %v\n", s1)
	s2 := s1[0:3] // [0, 3)
	fmt.Printf("s2: %v\n", s2)

	// 切片的添加
	s1 = append(s1, 10)
	fmt.Printf("s1: %v\n", s1)
	s1 = append(s1[:2], s1[3:]...)
	s2 = s1[0:3]
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("&s1: %p\n", &s1)
	fmt.Printf("&s2: %p\n", &s2)
}

func testMap() {
	var t map[int32]string
	t = make(map[int32]string)
	fmt.Printf("t: %v\n", t)
	var m = map[string]string{"1": "2", "2": "3"}
	fmt.Printf("m: %v\n", m)
	val, ok := m["3"]
	if ok {
		fmt.Printf("val: %v\n", val)
	}
}

func getName() (string, uint32) {
	fmt.Printf("math.MaxInt32: %v\n", math.MaxInt32)
	fmt.Printf("unsafe.Sizeof(math.MaxFloat32): %v\n", unsafe.Sizeof(math.MaxFloat32))
	return "tom", 15
}
