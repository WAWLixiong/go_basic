package learn_math

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func LearnMath() {
	fmt.Println(math.MaxInt)

	fmt.Println(math.Mod(10, 3))   // 取余数
	fmt.Println(math.Modf(10 / 3)) // 分开两部分
	fmt.Println(10 % 3)
	fmt.Println(10 / 3)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func LearnMath2() {

	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Println(a)
	}
	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		fmt.Println(a)
	}
	for i := 0; i < 10; i++ {
		a := rand.Float32()
		fmt.Println(a)
	}
}
