package wxpay

import "time"

func timeFormat(sec int64) string {
	return time.Unix(sec, 0).Format("20060102150405")
}

func timeParse(val string) int64 {
	if val == "" {
		return 0
	}
	t, err := time.ParseInLocation("2006-01-02 15:04:05", val, time.Local)
	if err != nil {
		panic(err)
	}
	return t.Unix()
}

func TimeParse2(val string) int64 {

	if val == "" {
		return 0
	}
	t, err := time.ParseInLocation("20060102150405", val, time.Local)
	if err != nil {
		panic(err)
	}

	return t.Unix()
}
