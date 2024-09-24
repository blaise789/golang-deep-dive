package main

/*import (
	"fmt"
	"log"
	"net/http"
	"time"
)
func helloWorld(w http.ResponseWriter,req *http.Request){
	ctx:=req.Context()
	log.Print(ctx)
	fmt.Println("server:hello handler started")
	defer  fmt.Println("server:hello handler ended")
  select{
  case <-time.After(10*time.Second):
	fmt.Fprintf(w,"hello\n")
  case <-ctx.Done():
	err:=ctx.Err()
	fmt.Println("server:",err)
	internalError:=http.StatusInternalServerError
	http.Error(w,err.Error(),internalError)
  }
	// fmt.Fprintf(w,"hello world")
	
}
func headers(w http.ResponseWriter,req *http.Request){
	for name,headers:=range req.Header{
	 for _,h:=range headers{
		fmt.Fprintf(w,"%v:%v\n",name,h)
	 }
	}
}

func main(){
	http.HandleFunc("/",helloWorld)
	http.HandleFunc("/headers",headers)
	http.ListenAndServe(":8000",nil)


}*/