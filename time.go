package main

import (
	"fmt"
	"time"
)

func main() {
	/*fmt.Println("Please input the day")
	var a int
	fmt.Scan(&a)

	fmt.Println("when is Saturday ?")
	var today = time.Now().Weekday()
	fmt.Println(today + 1)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}*/
	testTime()
	//getTime()
	/*a := DateToTimestamp("2018-06-30 15:04:05")
	fmt.Println(a)

	b := TimestampToDate(1530371045)
	fmt.Println(b)*/

}

func testTime() {
	fmt.Println(111, time.Now().Year())
	yearA := time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.Local)
	fmt.Println(222, yearA)
	yearB := yearA.AddDate(5, 0, 0)

	t1 := time.Date(yearA.Year(), 0, 0, 0, 0, 0, 0, time.Local)
	t2 := time.Date(yearB.Year(), 0, 0, 0, 0, 0, 0, time.Local)

	fmt.Println(yearA.Year())
	fmt.Println(yearB.Year())

	fmt.Println(t1.Year())
	fmt.Println(t2.Year())
}

func getTime() {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(timeStr)

	//使用Parse 默认获取为UTC时区 需要获取本地时区 所以使用ParseInLocation
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 23:59:59", time.Local)
	t2, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)

	fmt.Println(t, t2)

	fmt.Println(t.Unix() + 1)
	fmt.Println(t2.AddDate(0, 0, 1).Unix())

}

/*输入一个时间戳，得到ymdhis的格式*/
func TimestampToDate(timestamp int64) string {
	str_time := time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
	return str_time
}

/*输入一个日期格式，得到时间戳*/
func DateToTimestamp(timestamp string) int64 {
	var layout string
	if len(timestamp) > 10 {
		layout = "2006-01-02 15:04:05"
	} else {
		layout = "2006-01-02"
	}
	the_time, _ := time.Parse(layout, timestamp)
	return the_time.Unix()
}
