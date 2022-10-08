package learn_errors

import (
	"errors"
	"fmt"
	"time"
)

func check(s string) (string, error) {
	if s == "" {
		err := errors.New("字符串不能为空")
		return "", err
	} else {
		return s, nil
	}
}

func LearnErrors() {
	s, err := check("")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(s)
	}
}

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("发生时间:%d，错误描述:%s", e.When, e.When)
}

func oops() error {
	return MyError{
		time.Date(2022, 10, 8, 12, 0, 0, 0, time.UTC),
		"the file system has go no way",
	}
}

func LearnErrors2() {
	err := oops()
	if err != nil {
		return
	}

}
