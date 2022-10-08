package learn_bytes

import (
	"bytes"
	"fmt"
)

func LearnBytes() {
	// 类型转换
	var i int = 100
	var b byte = 10
	b = byte(i)
	i = int(b)

	var s string = "hello world"
	b1 := []byte{1, 2, 3}

	s = string(b1)
	b1 = []byte(s)

}

func LearnBytes2() {
	s := "hello world"
	b := []byte(s)
	b1 := []byte("hello world")
	b2 := []byte("Hello world")

	r1 := bytes.Contains(b1, b)
	fmt.Println(r1)

	r2 := bytes.Contains(b2, b1)
	fmt.Println(r2)
}

func LearnBytes3() {
	s := "hello world"
	b := []byte(s)
	c := bytes.Count(b, []byte{'o'})
	fmt.Println(c)
}

func LearnBytes4() {
	s := "hello world"
	b := []byte(s)
	fmt.Println(string(bytes.Repeat(b, 2)))
}

func LearnBytes5() {
	// 0, 1, -1 , 不替换，替换1次，不限制次数
	s := "hello world"
	b := []byte(s)
	replace := bytes.Replace(b, []byte{'o'}, []byte("AA"), 2)
	fmt.Println(string(replace))
}

func LearnBytes6() {
	// 一个中文字符三字节
	s := []byte("你好世界")
	r := bytes.Runes(s)
	fmt.Println(len(s))
	fmt.Println(len(r))
}

func LearnBytes7() {
	s2 := [][]byte{[]byte("你好"), []byte("世界")}
	fmt.Println(string(bytes.Join(s2, []byte(","))))
}

func LearnBytes8() {
	data := "123456789"
	re := bytes.NewReader([]byte(data))
	fmt.Println(re.Len())  // 未读部分的长度
	fmt.Println(re.Size()) // 底层数据总长度

	buf := make([]byte, 2)

	for {
		// 读取数据
		n, err := re.Read(buf)
		if err != nil {
			break
		}
		fmt.Println(string(buf[0:n]))
	}

	fmt.Println("--------------")

	// 设置偏移量，因为上面的操作已经修改了读取位置等信息
	re.Seek(0, 0)
	for {
		b, err := re.ReadByte()
		if err != nil {
			break
		}
		fmt.Println(string(b))
	}

	re.Seek(0, 0)
	off := int64(0)

	for {
		// 指定偏移读取
		n, err := re.ReadAt(buf, off)
		if err != nil && n == 0 { // n == 0 才能判断没有结果
			fmt.Println(err)
			break
		}
		off += int64(n)
		fmt.Println(off, string(buf[:n]))
	}

}

func LearnBytes9() {
	var b bytes.Buffer
	fmt.Printf("b: %v\n", b)
	// fmt.Println(len(b), cap(b))

	var b1 = bytes.NewBufferString("hello")
	fmt.Printf("b1:%v\n", b1)

	var b2 = bytes.NewBuffer([]byte("hello"))
	fmt.Printf("b1:%v\n", b2)

}

type Hello struct {
	name []byte
	age  int
}

func LearnBytes10() {
	var b bytes.Buffer
	b.WriteString("hello") // caution: 只要是类型，就可以调用其类型方法
	fmt.Printf("%v\n", string(b.Bytes()))
	fmt.Println()

	var c Hello
	fmt.Println(c.name, len(c.name), cap(c.name))
}

func LearnBytes11() {
	var b = bytes.NewBufferString("hello world")
	b1 := make([]byte, 2)
	for {
		n, err := b.Read(b1)
		if err != nil {
			break
		}
		fmt.Println(n)
		fmt.Println(string(b1)) // caution: 不使用切片导致最后输出为dl
		fmt.Println(string(b1[:n]))
	}

}
