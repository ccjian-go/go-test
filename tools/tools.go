package tools

import (
	"fmt"
	"time"
)

func main(){

}

func Test(){
	fmt.Print("test")
}

// 获取两个日期相差多少天
func GetDayGap(now time.Time, before time.Time) int {
	now = now.UTC().Truncate(24 * time.Hour)
	before = before.UTC().Truncate(24 * time.Hour)
	sub := now.Sub(before)
	return int(sub.Hours() / 24)
}