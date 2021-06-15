package util

import (
	"fmt"
	time "time"
)

// ObtainWorkDayToCurrentWeek 获取当前星期的工作日
func ObtainWorkDayToCurrentWeek() []string {
	// 计算当前周工作日期
	dates := getWorkingDays()
	// weekNum := whichWeek()
	for i, s := range dates {
		println(i, s)
	}
	return dates
	// 返回周一周五具体时间yyyyMMdd 返回周一到周五时间yyyy.MM.dd 月份,第几周
}
func whichWeek() int {
	// 获取这个月的第一天,并知晓这一天星期几 假设为星期6 则x = 6
	month := time.Now().Month()
	year := time.Now().Year()
	day := getYearMonthToDay(year, int(month))
	weekNum := 0
	//nowDay := time.Now().Day()
	nowDay := 22
	if nowDay<3 {
		return 4
	}

	for i := 0; i < day; i++ {
		if i%7 == 0 {
			if weekNum<4 {
				weekNum++
			}
		}
		if nowDay == i {
			break
		}

	}
	return weekNum
}

// 获取当前周第一天日期 考虑跨月,跨年
func getWorkingDays() []string {
	nowDay := time.Now().Day()
	nowWeekDay := int(time.Now().Weekday())
	month := int(time.Now().Month())
	years := time.Now().Year()
	// 星期一 今天, 今天星期几
	timeDifference := (nowDay - nowWeekDay)+1
	// 确认当前周开始日期
	startDate:=0
	maxDay:=0
	// 创建日期容器
	workDates := make([]string, 0, 5)
	if timeDifference < 0 {
		month--
		if month < 0 {
			years,month = years-1,12
		}
		maxDay = getYearMonthToDay(years, month)
		startDate = maxDay+timeDifference
	}else {
		maxDay = getYearMonthToDay(years, month)
		startDate = timeDifference
	}
	for i := 0; i < 5; i++ {
		p:= startDate+i
		if p>maxDay {
			p=1
			if month==12 {
				years,month = years+1,1
			}
		}
		workDates = append(workDates, fmt.Sprintf("%d.%d.%d",years,month,p))
	}
	return workDates
}

func getYearMonthToDay(year int, month int) int {
	// 有31天的月份
	day31 := map[int]bool{
		1:  true,
		3:  true,
		5:  true,
		7:  true,
		8:  true,
		10: true,
		12: true,
	}
	if day31[month] == true {
		return 31
	}
	// 有30天的月份
	day30 := map[int]bool{
		4:  true,
		6:  true,
		9:  true,
		11: true,
	}
	if day30[month] == true {
		return 30
	}
	// 计算是平年还是闰年
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		// 得出2月的天数
		return 29
	}
	// 得出2月的天数
	return 28
}
