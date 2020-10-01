package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	time "time"
)

// Page 分页
func Page(Limit, Page int) (limit, offset int) {
	if Limit > 0 {
		limit = Limit
	} else {
		limit = 10
	}
	if Page > 0 {
		offset = (Page - 1) * limit
	} else {
		offset = -1
	}
	return limit, offset
}

// Sort 排序
// 默认 created_at desc
func Sort(Sort string) (sort string) {
	if Sort != "" {
		sort = Sort
	} else {
		sort = "create_at desc"
	}
	return sort
}

const TimeLayout = "2006-01-02 15:04:05"

var (
	Local = time.FixedZone("CST", 8*3600)
)

func GetNow() string {
	now := time.Now().In(Local).Format(TimeLayout)
	return now
}

func TimeFormat(s string) string {
	result, err := time.ParseInLocation(TimeLayout, s, time.Local)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	return result.In(Local).Format(TimeLayout)

}

func Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}
