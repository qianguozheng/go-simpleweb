package airdisk

import (
	"time"
)

func startTimer(f func())  {
	go func() {
		for{
			f()
			now:= time.Now()
			//计算下一个零点
			next := now.Add(time.Hour *2)
			next = time.Date(next.Year(),next.Month(),next.Day(), next.Hour(), next.Minute(), next.Second(), 0,
				next.Location())
			t := time.NewTicker(next.Sub(now))
			<- t.C
		}
	}()
}


