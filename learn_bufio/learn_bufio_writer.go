package learn_bufio

import (
	"bufio"
	"os"
)

func LearnBufIO7() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR, 0755) // f及实现了io.reader，又实现了io.writer
	bufio.NewWriter(f)

}
