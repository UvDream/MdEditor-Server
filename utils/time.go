package utils

import "time"

// GetDateList 求两个时间之间的所有日期集合
func GetDateList(start, end string, isYear string) (dateList []time.Time, err error) {
	startTime, err := time.Parse("2006-01-02 15:04:05", start)
	if err != nil {
		return nil, err
	}
	endTime, err := time.Parse("2006-01-02 15:04:05", end)
	if err != nil {
		return nil, err
	}
	if isYear == "0" {
		for startTime.Before(endTime) || startTime.Equal(endTime) {
			dateList = append(dateList, startTime)
			startTime = startTime.AddDate(0, 0, 1)
		}
	} else {
		year := startTime.Year()
		for i := 0; i < 12; i++ {
			dateList = append(dateList, time.Date(year, time.Month(i+1), 1, 0, 0, 0, 0, time.Local))
		}
	}
	return dateList, nil
}
