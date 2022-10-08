package learn_bufio

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func TestBufIO1() {
	f, _ := os.Open("a.txt")
	defer f.Close()

	r := bufio.NewReader(f)
	s, _ := r.ReadString('\n')
	fmt.Printf("s:%v\n", s)

}

func TestBufIO2() {
	s := strings.NewReader("ABCEFG")
	str := strings.NewReader("123455")
	br := bufio.NewReader(s)
	b, _ := br.ReadString('\n')
	fmt.Println(b)
	br.Reset(str) // 清空之前读到的内容
	b, _ = br.ReadString('\n')
	fmt.Println(b)
}

func TestBufIO3() {
	s := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	br := bufio.NewReader(s)
	p := make([]byte, 10)

	for {
		_, err := br.Read(p)
		if err == io.EOF {
			break
		} else {
			fmt.Println(string(p))
		}
	}
}

func TestBufIO4() {
	s := strings.NewReader("hello\nworld\ni am fine\n")
	br := bufio.NewReader(s)

	for {
		line, prefix, err := br.ReadLine()
		if err == io.EOF {
			return
		}
		fmt.Printf("%q %v\n", line, prefix)
	}

}

func TestBufIO5() {
	s := strings.NewReader("ABC,DEF,GHI")
	br := bufio.NewReader(s)
	for {
		readString, err := br.ReadString(',')
		if err != nil {
			return
		}
		fmt.Printf("%s\n", readString)
	}
}

func TestBufIO6() {
	s := strings.NewReader("ABC,DEF,GHI")
	br := bufio.NewReader(s)

	b := make([]byte, 2)
	w := bytes.NewBuffer(b)
	_, err := br.WriteTo(w)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("%s\n", w)
	}

}
