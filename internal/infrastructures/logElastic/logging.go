package logElastic

import "time"

func getDateTime() string {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	dt := time.Now().In(loc)
	return dt.Format("2006-01-02 15:04:05.00000000")
}
