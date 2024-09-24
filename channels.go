package main

/*import (
	"fmt"
	// "time"
)
func main(){
	// time.Sleep(time.Minute)
	 
	messages:=make(chan string)
	// unbuffered channel
	// num_ch:=make(chan  int,3)
	
	// num_ch<-4
	// num_ch<-5
	// num_ch<-6
	// for num:=range num_ch{
	// fmt.Println(num)

	// }
	// sending the value to the channel

	// num:=<-num_ch
	// recieving values from the channel

	// fmt.Println(num)
   go func(){

		messages<-"pinging data from the first go routine"

	}()
	close(messages)
	fmt.Println("hello")
	msg,ok:=<-messages
	if !ok{
		fmt.Println("channel is empty")
	}

	fmt.Println(msg)
}
	*/