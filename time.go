package main

import (
	"fmt"
	"time"
)
func main(){
	p:=fmt.Println
	now:=time.Now()
	
	// p(now)
	// then:=time.Date(
	// 	2007,01,31,23,59,40, 651387237,time.UTC)
    // p(then)
	// p(then.Weekday())
	// p(then.Before(now))
	// p(then.After(now))
	// p(int(now.Sub(then).Hours()/(24*365.5)),"years old")
	// diff:=now.Sub(then)
	
	// p(diff)
	// p("linux is ",now.Unix()/(365*24*60*60),"old")
    // p(time.Unix(now.Unix(),0))
	t:=now.Format(time.RFC3339)
	t1,e:=time.Parse(time.RFC1123,"2012-11-01T22:08:41+00:00")
	if e!=nil{
		fmt.Println(e.Error())
	}
    fmt.Println(t1)

	p(now)
	p(t)
	p(now.Format("3:04PM"))
	ansic:="Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
    p(e)

	// 



}
