package modules

import "time"

type Dates struct {
}

func (date Dates) GetDate() string {
	temp := "2006-01-02 15:04:05"
	return time.Now().Format(temp)
}

func (date Dates) GetDay() string {
	temp := "20060102"
	return time.Now().Format(temp)
}

func (date Dates) GetUnix() int64 {
	return time.Now().Unix()
}
