package learn_log

import (
	"fmt"
	"log"
	"os"
)

func LearnLog() {
	// print 单纯打印日志
	// panic 打印日志,抛出panic异常, defer会执行
	// fatal 打印日志， os.Exit(1), defer不会执行

	log.Print("my log")
	log.Printf("my log %d", 100)
	name := "tom"
	age := 20
	log.Println(name, " ", age)

	log.Panic("my panic")
	fmt.Println("end...")

	log.Fatal("end")

}

func LearnLog2() {
	i := log.Flags()
	fmt.Println(i)
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func LearnLog3() {
	log.SetPrefix("my log:")
	f, err := os.OpenFile("a.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		return
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println("hello")
}

func LearnLog4() {
	f, err := os.OpenFile("a.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		return
	}
	defer f.Close()
	logger := log.New(f, "my log:", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("hello world i am fine")

}
