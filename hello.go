package main

import (
	"fmt"
	"log"
)

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func connectDB() {
	panic("unconnected DB")
}

type IntType struct {
	X, Y int
}

func (v IntType) Area() int {
	return v.X * v.Y
}

func (v *IntType) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

func save() {
	defer func() {
		s := recover() //panicの内容を受け取り強制終了しないようにする
		fmt.Println(s)
	}()
	connectDB()
}

func one(x *int) {
	*x = 1 //引数のアドレスの中身（メモリ）に1を渡している

}

type Man interface {
	Say()
}

type Person struct {
	Name string
}

func (p *Person) Say() {
	p.Name = p.Name + "san"
	fmt.Println(p.Name)
}

func main() {
	// c := make([]int, 0, 5)//メモリが5個確保されていて、0が一つ入っている配列
	c := make([]int, 5) //0が５個入っている配列で、メモリは特に宣言されていない
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)

	// var m map[string]int //varで宣言するとnilになり、メモリーが確保できないためエラーになる
	// m["px"] = 500
	// fmt.Println(m)

	m2 := make(map[string]int)
	m2["px"] = 500
	fmt.Println(m2)

	m3 := map[string]int{"apple": 500, "banana": 300}
	fmt.Println(m3)

	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	c1 := circleArea(3.14)
	fmt.Println(c1(2))

	c2 := circleArea(3)
	fmt.Println(c2(2))

	len := []string{"js", "ts", "go"}

	for i, v := range len {
		fmt.Println(i, v)
	}
	for _, v := range len {
		fmt.Println(v)
	}

	m4 := map[string]int{"apple": 500, "orange": 200}

	for k, v := range m4 {
		fmt.Println(k, v)
	}

	for k := range m4 {
		fmt.Println(k)
	}

	for _, v := range m4 {
		fmt.Println(v)
	}

	log.Println("logging")
	// log.Fatalln("error") //この時点でプログラムが止まる

	save()
	fmt.Println("0K?")

	var n int = 100 //メモリの中身
	fmt.Println(n)
	fmt.Println(&n) //100を入れたメモリのアドレスを表示する

	var p *int = &n //intergerのpoint型に100を入れたメモリのアドレスを代入
	fmt.Println(p)
	fmt.Println(*p) //アドレスが指しているメモリの中身が表示される

	var n1 int = 100
	one(&n1)
	println(n1)

	var n3 *int = new(int) //値はないがメモリの領域を確保している、ポインターが返ってくる返り値に関してはnewを使う
	fmt.Println(n3)

	v := IntType{2, 3}
	v.Scale(10)
	fmt.Println(v.Area())

	var max Man = &Person{"max"}
	max.Say()
}
