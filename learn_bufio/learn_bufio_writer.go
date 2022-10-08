package learn_bufio

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func LearnBufIO7() {
	f, err := os.OpenFile("a.txt", os.O_RDWR, 0755) // f及实现了io.reader，又实现了io.writer
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close() // 缺失这句话的话，可能没有写进去
	w := bufio.NewWriter(f)
	_, err = w.WriteString("hello world")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		w.Flush() // 缺失这句话的话，可能没有写进去
	}

}

func LearnBufIO8() {
	b := bytes.NewBuffer(make([]byte, 2))
	w := bufio.NewWriter(b)

	_, err := w.WriteString("hello world")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		w.Flush() // flush
		fmt.Printf("%s", b)
	}
}

func LearnBufIO9() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	bw.WriteString("123456789")
	c := bytes.NewBuffer(make([]byte, 0))
	bw.Reset(c) // 会清空之前的Writer
	bw.WriteString("456")
	bw.Flush()
	fmt.Println(b)
	fmt.Println(c)
}

func LearnBufIO10() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	fmt.Println(bw.Available())
	fmt.Println(bw.Buffered())

	bw.WriteString("ABCDEFGHIJKLMN")
	fmt.Println(bw.Available())
	fmt.Println(bw.Buffered())
	fmt.Printf("%q\n", b)

	bw.Flush()
	fmt.Println(bw.Available())
	fmt.Println(bw.Buffered())
	fmt.Printf("%q\n", b)

}

func LearnBufIO11() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)

	bw.WriteByte('H')
	bw.WriteRune('你')

	bw.Flush()
	fmt.Println(b)
}

func LearnBufIO12() {
	b := bytes.NewBuffer(make([]byte, 0))
	s := strings.NewReader("Hello 你好世界!")
	bw := bufio.NewWriter(b)
	bw.ReadFrom(s)

	fmt.Println(b) // 不需要 Flush
}

func LearnBufIO13() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	s := strings.NewReader("123")
	br := bufio.NewReader(s)

	brw := bufio.NewReadWriter(br, bw)

	p, _ := brw.ReadString('\n')
	fmt.Println(p)
	brw.WriteString("asdf")
	brw.Flush()
	fmt.Println(b)
}
