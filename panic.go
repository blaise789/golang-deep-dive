 package main

 /* import (
  	"fmt"
  	"os"
  )

  

  func main() {
  	//  panic("a problem")
  	 _, err := os.Create("/tmp/file")
  	 if err != nil {
  	 	panic(err.Error())
  	 }
  	  fmt.Println("hello")
	

  	 defer closeFile(res)
	 
  	defer func(){
  		if r:=recover();r !=nil{
  		    fmt.Println("main go routine is recovered",r)
  		}

  	}()
  	res:=createFile("./hello.txt")

  	writeFile(res,"hello brother","hey man")
	


  }
  func createFile(p string) *os.File{
  	f,err:=os.Create(p)
  	if err !=nil { 
  		panic(err)
  	}
  	return f
  }

  func closeFile(f *os.File){
  	err:=f.Close()
  	if err !=nil{
  		fmt.Fprintf(os.Stderr,"error %v\n",err)
  		panic( err)
  	}


  }
  func writeFile(f *os.File,data ...any){
	
 
	// fmt.Println(data...)
 	fmt.Fprintln(f,data...)
  }
*/