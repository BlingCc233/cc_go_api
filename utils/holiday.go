package utils

import (
	"fmt"
	"github.com/6tail/lunar-go/calendar"
	"studyGo/models"
	"time"
)

// 定义节假日（公历和农历）
var holidays = []string{
	"元旦", "除夕", "春节", "元宵节", "清明节", "劳动节", "端午节", "中秋节", "国庆节", "七夕", "圣诞节", "情人节", "母亲节", "父亲节", "愚人节",
}

func GetHolidays(now time.Time) []models.Holiday {
	fmt.Printf("今天是：%s\n", now.Format("2006-01-02"))
	var countdowns []models.Holiday
	firstWeekend := true
	// 365天的循环
	for i := 1; i <= 365; i++ {
		// 计算日期
		date := now.AddDate(0, 0, i)
		// 判断是否为周六
		if date.Weekday() == time.Saturday && firstWeekend {
			firstWeekend = false
			countdowns = append(countdowns, models.Holiday{Name: "周末", Days: i})
		}
		solar := calendar.NewSolarFromDate(date)
		var thisDayHoliday []string
		fesList := solar.GetFestivals()

		for o := fesList.Front(); o != nil; o = o.Next() {
			thisDayHoliday = append(thisDayHoliday, o.Value.(string))
		}

		// 转阴历
		lunar := solar.GetLunar()
		fesList = lunar.GetFestivals()
		for o := fesList.Front(); o != nil; o = o.Next() {
			thisDayHoliday = append(thisDayHoliday, o.Value.(string))
		}
		for _, h := range thisDayHoliday {
			for _, hd := range holidays {
				if h == hd {
					countdowns = append(countdowns, models.Holiday{Name: h, Days: i})
				}
			}
		}
	}

	// 排序countdowns
	sortCountdowns(&countdowns)

	// 遍历countdowns列表
	for _, cd := range countdowns {
		fmt.Printf("距离%s还有：%d天\n", cd.Name, cd.Days)
	}
	return countdowns
}

func sortCountdowns(countdowns *[]models.Holiday) {
	for i := 0; i < len(*countdowns); i++ {
		for j := i + 1; j < len(*countdowns); j++ {
			if (*countdowns)[i].Days > (*countdowns)[j].Days {
				(*countdowns)[i], (*countdowns)[j] = (*countdowns)[j], (*countdowns)[i]
			}
		}
	}
}
