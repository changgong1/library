package util

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	WeekDay    = 7
	DayHour    = 24
	MinSecond  = 60
	HourSecond = 3600
	DaySecond  = 86400
	WeekSecond = 604800
	IsoLength  = 25 // 2022-10-25T20:09:16-06:00
	DateLength = 10 // 2022-10-25
	PaserImp   = "2006-01-02 15:04:05"
)

var zoneJp int32 = 9 * 3600

func UnixTime() int32 {
	return int32(time.Now().Unix())
}
func UnixTimeStr() string {
	s := strconv.Itoa(int((time.Now().Unix())))
	return s
}

func UnixMillTime() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func UnixMillTimeStr() string {
	s := strconv.FormatInt(UnixMillTime(), 10)
	return s
}

// 2006-01
func GetNowMonth() string {
	return time.Now().Format("2006-01")
}

// 2006-01-10
func GetNowDate() string {
	return time.Now().Format("2006-01-02")
}

// 2006-01-10
func GetNowTime() string {
	return time.Now().Format(PaserImp)
}

// 2006-12
// 获取前n个月 0-表示当前月 <0 表示往前n个月 >0 往后n个月
// n最大12个月
func GetNumYearMonth(yearMonth string, n int) string {
	if len(yearMonth) != 7 {
		return ""
	}
	if n == 0 {
		return yearMonth
	}
	mStr := yearMonth[6:]
	y, _ := strconv.Atoi(yearMonth[:4])
	m := 0
	if mStr[:1] == "0" {
		m, _ = strconv.Atoi(mStr[1:])
	} else {
		m, _ = strconv.Atoi(mStr)
	}

	dm := m + n
	mo := dm
	if dm <= 0 {
		// 上一年
		y--
		mo = 12 + dm

	} else if dm > 12 {
		// 下一年
		y++
		mo = dm - 12
	}
	yeStr := strconv.Itoa(y)
	moStr := strconv.Itoa(mo)
	if mo < 10 {
		moStr = "0" + moStr
	}
	return yeStr + "-" + moStr
}

// 2006-01-02
func GetUnixTimeDate(timestramp int32) string {
	fix := time.FixedZone("UTC", 0)
	timeObj := time.Unix(int64(timestramp), 0)
	return timeObj.In(fix).Format("2006-01-02")
}

// 2006-01-02 15:04:05
func GetUnixTimeDateTime(timestramp int32) string {
	fix := time.FixedZone("UTC", 0)
	timeObj := time.Unix(int64(timestramp), 0)
	return timeObj.In(fix).Format("2006-01-02 15:04:05")
}

// 2006-01-02 15:04:05 北京时间日期
func GetCstTimeDateTime(timestramp int32) string {
	timestramp += 8 * 3600
	fix := time.FixedZone("UTC", 0)
	timeObj := time.Unix(int64(timestramp), 0)
	return timeObj.In(fix).Format("2006-01-02 15:04:05")
}

// 20060102150405
func GetUnixTimeDateTimeHH(timestramp int32) string {
	fix := time.FixedZone("UTC", 0)
	timeObj := time.Unix(int64(timestramp), 0)
	return timeObj.In(fix).Format("20060102150405")
}

// 2006-01-02 15:04:05
func ParseTimeUnix(dateTime string) (int32, error) {
	t, err := time.Parse("2006-01-02 15:04:05", dateTime)
	if err != nil {
		return 0, err
	}
	return int32(t.Unix()), nil
}

// dateStr = 2022-01-02
func GetPreDay(dateStr string, day int) string {
	if day == 0 {
		return dateStr
	}
	dateStr += " 00:00:00"
	ti, _ := ParseTimeUnix(dateStr)
	sDay := int(ti) - day*DaySecond
	return GetUnixTimeDate(int32(sDay))
}

// 返回当前时间距离指定时间过了几天
func GetOverDay(preStamp int32, hour, minutes, second int) (day int32) {
	day = GetOverDay2(preStamp, UnixTime(), hour, minutes, second)
	return
}

// 返回两个时间相隔多少天 同一天返回0
func GetOverDayTwo(preStamp, eStamp int32) (day int32) {
	return GetOverDay2(preStamp, eStamp, 0, 0, 0)
}

// 返回两个时间间隔多少天 同一天返回0
func GetOverDay2(preStamp, nowStamp int32, hour, minutes, second int) (day int32) {
	PrevTm := time.Unix(int64(preStamp), 0)
	d := nowStamp - preStamp

	seconds := 0
	if PrevTm.Hour() < hour {
		seconds = (hour - PrevTm.Hour()) * HourSecond
	} else {
		seconds = ((DayHour - PrevTm.Hour()) + hour) * HourSecond
	}

	seconds = (seconds + minutes*MinSecond + second) - (PrevTm.Minute()*MinSecond + PrevTm.Second())
	if d >= int32(seconds) {
		day += 1
		day += (d - int32(seconds)) / DaySecond
	}

	return
}

// 获取今天N点N时N分的时间戳
func GetDayTimeStamp(h, m, s int) (timeStamp int32) {
	currentTime := time.Now()
	timeStamp = int32(time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), h, m, s, 0, currentTime.Location()).Unix())
	return
}

// 获取让个月的年和月 2023-01
func GetLastYearMonth() string {
	nowTime := time.Now()
	year := nowTime.Year()
	month := nowTime.Month()
	month -= 1
	if month == 1 {
		year -= 1
		month = 12
	}
	yearStr := strconv.Itoa(year)
	monStr := strconv.Itoa(int(month))
	if len(monStr) < 2 {
		monStr = "0" + monStr
	}

	return yearStr + "-" + monStr
}

// 获取上个月 yearMonth:2023-08
func GetPreMonth(yearMonth string) string {
	ym := strings.Split(yearMonth, "-")
	y, _ := strconv.Atoi(ym[0])
	m, _ := strconv.Atoi(ym[1])
	if m == 1 {
		m = 12
		y--
	} else {
		m--
	}
	mStr := strconv.Itoa(m)
	if m < 10 {
		mStr = "0" + mStr
	}

	return strconv.Itoa(y) + "-" + mStr
}

// 获取上个月 yearMonth:2023-08
func GetNextMonth(yearMonth string) string {
	ym := strings.Split(yearMonth, "-")
	y, _ := strconv.Atoi(ym[0])
	m, _ := strconv.Atoi(ym[1])
	if m == 12 {
		m = 1
		y++
	} else {
		m++
	}
	mStr := strconv.Itoa(m)
	if m < 10 {
		mStr = "0" + mStr
	}

	return strconv.Itoa(y) + "-" + mStr
}

// 获取本月最后一天 yearMonth:2023-08
func GetMonthEndDay(yearMonth string) string {
	nM := GetNextMonth(yearMonth)
	nt, _ := ParseTimeUnix(nM + "-01 00:00:01")
	return GetUnixTimeDate(nt - 120)
}

// 获取上n个月的年和月 [2023-07 2023-06]
func GetLastTowYearMonth(n int) []string {
	monthList := make([]string, 0)
	nowTime := time.Now()
	year := nowTime.Year()
	month := nowTime.Month()
	for i := 0; i < n; i++ {
		month -= 1
		if month == 1 {
			year -= 1
			month = 12
		}
		yearStr := strconv.Itoa(year)
		monStr := strconv.Itoa(int(month))
		if len(monStr) < 2 {
			monStr = "0" + monStr
		}
		monthList = append(monthList, yearStr+"-"+monStr)
	}

	return monthList
}

// 两个时间戳是否同一天
func IsSameDay(timeStampA, timeStampB int32) bool {
	tmA := time.Unix(int64(timeStampA), 0)
	tmB := time.Unix(int64(timeStampB), 0)
	return tmA.Year() == tmB.Year() && tmA.YearDay() == tmB.YearDay()
}

// 返回某个时刻那天的24点的时间戳
func Get24TimeStamp(timeStamp int32) (result int32) {
	tm := time.Unix(int64(timeStamp), 0)
	result = int32(time.Date(tm.Year(), tm.Month(), tm.Day(), 23, 59, 59, 0, tm.Location()).Unix())
	return
}

// 返回指定时间的起始时间
func Get0TimeStamp(timeStamp int32) (result int32) {
	result = Get24TimeStamp(timeStamp) - 86400 + 1
	return
}

func GetHour() int32 {
	return int32(time.Now().Hour())
}

func ToStr(d int32) string {
	return strconv.Itoa(int(d))
}

func GetMinute() int32 {
	return int32(time.Now().Minute())
}

func GetSecond() int32 {
	return int32(time.Now().Second())
}

func GetWeek() time.Weekday {
	return time.Now().Weekday()
}

func GetZone() int32 {
	_, zone := time.Now().Zone()
	return int32(zone / 3600)
}

// 获取某年某月的天数
func GetMonthDay(year, month int) int {
	yStr := strconv.Itoa(year)
	mStr := strconv.Itoa(month)
	if len(mStr) == 1 {
		mStr = "0" + mStr
	}
	startDay := yStr + "-" + mStr + "-01 00:00:00"
	startInt, _ := ParseTimeUnix(startDay)

	endyStr := ""
	eMstr := ""
	if month == 12 {
		endyStr = strconv.Itoa(year + 1)
		eMstr = "01"
	} else {
		endyStr = strconv.Itoa(year)
		em := month + 1
		eMstr = strconv.Itoa(em)
		if len(eMstr) == 1 {
			eMstr = "0" + eMstr
		}
	}
	endDay := endyStr + "-" + eMstr + "-01 00:00:00"
	endInt, _ := ParseTimeUnix(endDay)
	day := GetOverDayTwo(startInt, endInt)
	return int(day)
}

// 2006-01-02T15:04:05Z
func GetUnixTimeDateTimeTZ(timestramp int32) string {
	fix := time.FixedZone("UTC", 0)
	timeObj := time.Unix(int64(timestramp), 0)
	return timeObj.In(fix).Format("2006-01-02T15:04:05Z")
}

// return 2006-01-02T15:04:05Z"
func GetUnixPreMonthTZ(day int) string {
	nowTime := int(UnixTime())
	startTime := nowTime - day*DaySecond
	startStr := GetUnixTimeDate(int32(startTime))

	date := startStr + "T15:04:05Z"
	return date
}

// dataTime 2006-01-02T15:04:05Z return 2006-01-02 15:04:05
func TZTimeTOTime(dateTime string) string {
	dateStr := strings.Replace(dateTime, "T", " ", -1)
	return dateStr[:19]
}

// dataTime 2006-01-02 15:04:05 return 2006-01-02T15:04:05Z
func TimeTOTzTime(dateTime string) string {
	dateStr := strings.Replace(dateTime, " ", "T", 1)
	dateStr += "Z"
	return dateStr[:20]
}

// dataTime 2022-10-25T20:09:16-06:00 return 2006-01-02T15:04:05Z
func TZoneToUtcTZ(dateTime string) string {
	t, err := DateIsoToUTCTime(dateTime)
	if err != nil {
		return ""
	}
	return t.Format("2006-01-02T15:04:05Z")
}

// dataTime 2022-10-25T20:09:16-06:00 return 秒时间戳
func DateIsoToTimeUnix(dateIso string) (int, error) {
	ti, err := DateIsoToUTCTime(dateIso)
	if err != nil {
		return 0, err
	}
	return int(ti.Unix()), err
}

// dateIso 2022-10-25T20:09:16-06:00
func DateIsoToUTCTime(dateIso string) (*time.Time, error) {
	ti, err := time.Parse(time.RFC3339, dateIso)
	if err != nil {
		return nil, err
	}
	fix := time.FixedZone("UTC", 0)
	tti := ti.In(fix)
	return &tti, nil
}

// dateIso 2022-10-25T20:09:16-0600 return 2022-10-25T20:09:16-06:00
func DateIsoToT(dateIso string) string {
	oTime := dateIso[:22] + ":00"
	return oTime
}

// dateIso 2022-10-25T20:09:16-06:00 return 2006-01-02
func GetIsoDate(dateIso string) string {
	uc, _ := DateIsoToTimeUnix(dateIso)
	return GetUnixTimeDate(int32(uc))
}

// dateIso 2022-10-25T20:09:16-06:00 return 2006-01-02 15:04:05
func GetIsoDateTime(dateIso string) string {
	uc, _ := DateIsoToTimeUnix(dateIso)
	return GetUnixTimeDateTime(int32(uc))
}

// date 2022-10-25T20:09:16-0600 return 2022-10-25T20:09:16-06:00
func DatetoDateIso(date string) string {
	if len(date) != IsoLength-1 {
		return ""
	} else {
		str := date[:IsoLength-3] + ":00"
		return str
	}
}

func LocalZone() int {
	_, offset := time.Now().Zone()
	return offset
}

func RateChar(value float32) string {
	s := fmt.Sprintf("%.2f", value*100)
	return s + "%"
}

func TwoDecimal(value float32) string {
	s := fmt.Sprintf("%.2f", value)
	return s
}

// 夏令时转utc时间
// date:2023-07-01
// offset 时区
// summerOffsert 夏日调整量 -表示美国 +表示中欧
func SummerTimeToUtc(date string, offset, summerOffsert int) (string, error) {
	utcDate := date + "T00:00:00Z"
	ti, err := time.Parse("2006-01-02T15:04:05Z", utcDate)
	if err != nil {
		return "", err
	}
	utcTime := ti.Unix()

	if summerOffsert == 0 {
		utcTime -= int64(offset) * HourSecond
		return GetUnixTimeDateTimeTZ(int32(utcTime)), nil
	}

	yearStr := strconv.Itoa(ti.Year())
	sumDateStart := ""
	sumDateEnd := ""

	// 美国夏令时
	if offset < 0 {
		// PDT 开始日期: 每年的三月的第二个星期日，凌晨2点整。结束日期: 每年的十一月的第一个星期日，凌晨2点整。
		sumDateStart = GetWeekst(yearStr+"-03", time.Sunday, 2)
		sumDateEnd = GetWeekst(yearStr+"-11", time.Sunday, 1)
	} else {
		// 欧洲夏令时 欧洲的夏令时从3月的最后一个星期天的凌晨2点开始，到10月份的最后一个星期天的凌晨1点59结束。
		sumDateStart = GetWeekLast(yearStr+"-03", time.Sunday)
		sumDateEnd = GetWeekLast(yearStr+"-10", time.Sunday)
	}
	if date > sumDateStart && date < sumDateEnd {
		utcTime -= int64(offset+summerOffsert) * HourSecond
	} else {
		utcTime -= int64(offset) * HourSecond
	}
	return GetUnixTimeDateTimeTZ(int32(utcTime)), nil
}

// 获取某年某月的第几个星期日
// yearMonth: 2023-01 ordinal 第几个 weekDay-星期x
// 返回 yyyy-mm-dd
func GetWeekst(yearMonth string, weekDay time.Weekday, ordinal int) string {
	monthOneDay := yearMonth + "-01 00:00:00"
	ti, err := time.Parse("2006-01-02 15:04:05", monthOneDay)
	if err != nil {
		return ""
	}
	wDay := ti.Weekday()
	dis := 0
	if wDay > weekDay {
		dis = 7 - int(wDay) + int(weekDay)
	} else {
		dis = int(weekDay - wDay)
	}
	day := 1 + int(math.Abs(float64(dis))) + (ordinal-1)*7

	dayStr := strconv.Itoa(day)
	if day < 10 {
		dayStr = "0" + dayStr
	}
	return yearMonth + "-" + dayStr
}

// 获取某年某月的最后一个个星期x
// yearMonth: 2023-01 weekDay-星期x
// 返回 yyyy-mm-dd
func GetWeekLast(yearMonth string, weekDay time.Weekday) string {
	// 计算下一个月的第一天
	yearStr := yearMonth[:4]
	monthStr := yearMonth[5:]
	sl := strings.Split(monthStr, "0")
	if len(sl) > 1 {
		monthStr = sl[1]
	}
	year, _ := strconv.Atoi(yearStr)
	month, _ := strconv.Atoi(monthStr)
	if month == 12 {
		monthStr = "01"
		year += 1
	} else {
		month += 1
		monthStr = strconv.Itoa(month)
	}
	addYearMonth := strconv.Itoa(year) + monthStr

	// 下一个月的第一天
	monthOneDay := addYearMonth + "01 00:00:00"
	ti, err := time.Parse("2006-01-02 15:04:05", monthOneDay)
	if err != nil {
		return ""
	}

	// 最后一天是星期x
	wDay := ti.Weekday() - 1
	if wDay < 0 {
		wDay = 1
	}

	dis := 0
	if wDay >= weekDay {
		dis = int(wDay - weekDay)
	} else {
		dis = int(wDay) + 7 - WeekDay
	}
	ut := ti.Unix()
	ut -= int64(dis * DaySecond)
	return GetUnixTimeDate(int32(ut))
}

// 获取每个西部时间 timestamp-毫秒时间戳
func GetUsLosDate(timestamp int) (string, error) {
	timestam := int64(timestamp) // 替换为您的时间戳

	// 使用time.Unix将时间戳转换为time.Time对象
	tm := time.Unix(timestam, 0)

	// 获取美国的本地时间
	// 可以根据需要将时区更改为其他美国时区，例如time.LoadLocation("America/Los_Angeles")
	usEastern, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		return "", err
	}
	localTime := tm.In(usEastern)
	// localTime.Format(util.PaserImp)
	// 打印美国的本地时间
	return localTime.Format(PaserImp), nil
}

// 前一天
func GetPreDate() string {
	t := time.Now()
	t = t.AddDate(0, 0, -1)
	return t.Format("2006-01-02")
}
