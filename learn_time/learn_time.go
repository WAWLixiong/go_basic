package learn_time

import (
	"fmt"
	"reflect"
	"time"
)

func LearnTime() {
	// time.Location{}
	// time.Duration()
	// time.Time{}
	// time.Month()
	// time.Weekday()

	now := time.Now()
	fmt.Println(reflect.TypeOf(now))
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d-%02d-%02d-%02d\n", year, month, day, hour, minute, second)
	fmt.Printf("%T-%T-%T-%T-%T-%T\n", year, month, day, hour, minute, second)

}

func LearnTime2() {
	// 时间戳
	now := time.Now()
	fmt.Println(now.Unix())
}

func LearnTime3() {
	// 时间戳转为时间
	timestamp := time.Now().Unix()
	timeObj := time.Unix(timestamp, 0)
	fmt.Println(timeObj)
	fmt.Println(timeObj.UTC())
}

func LearnTime4() {
	// 操作时间
	now := time.Now()
	newTime := now.Add(time.Hour * 1)
	sub := now.Sub(newTime)
	fmt.Println(sub.Hours())

	// 比较
	fmt.Println(now.Before(newTime))
	fmt.Println(now.After(newTime))
	fmt.Println(now.Equal(newTime))
}

func LearnTime5() {
	// 时间格式化

	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan")) // 24小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))

}

func LearnTime6() {
	// 解析时间格式
	now := time.Now()
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	timeObj, err := time.ParseInLocation("2006/01/02 15:04", "2022/10/08 14:07", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))
}
