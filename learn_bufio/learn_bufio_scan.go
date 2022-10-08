package learn_bufio

import (
	"bufio"
	"fmt"
	"strings"
)

func LearnBufIO14() {
	s := strings.NewReader("ABC DEF GHI JKL")
	bs := bufio.NewScanner(s)
	// bs.Split(bufio.ScanBytes)
	bs.Split(bufio.ScanWords)

	for bs.Scan() {
		fmt.Println(bs.Text())
	}
}
