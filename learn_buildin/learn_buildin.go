package learn_buildin

import "fmt"

func LearnBuildin() {
	s := []int{}
	i := append(s, 100)
	fmt.Println(i)
	fmt.Println(s)
	fmt.Printf("%p, %p\n", &s, &i)

	s1 := []int{4, 5, 6}
	i2 := append(s, s1...)
	fmt.Println(i2)
	fmt.Println(s1)
	fmt.Printf("%p, %p\n", &s1, &i2)
}

func LearnBuildin2() {
	// new 返回指针，make返回引用, make分配内存，new清除内存, 初始化的都是 0 值

	b := new(bool)
	fmt.Printf("b:%T\n", b)
	fmt.Printf("b:%v\n", *b)

	i := new(int)
	fmt.Printf("i:%T\n", i)
	fmt.Printf("i:%v\n", *i)

	s := new(string)
	fmt.Printf("s:%T\n", s)
	fmt.Printf("s:%v\n", *s)
}

func LearnBuildin3() {
	var p *[]int = new([]int)
	var v []int = make([]int, 10)
	fmt.Println(p)
	fmt.Println(v)

	// 蛋疼的做法
	p = new([]int)
	*p = make([]int, 10, 10)

	// 习惯做法
	v = make([]int, 10)

}
