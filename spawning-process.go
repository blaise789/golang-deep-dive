package main

import (
	"fmt"
	"io"
	"os/exec"
)
func main(){
	dateCmd:=exec.Command("date")
	dateOutput,err:=dateCmd.Output()
	if err!=nil{
		panic(err)
	}

	fmt.Println(">date")
	fmt.Println(string(dateOutput))
    _,err=exec.Command("date","-x").Output()
	if err !=nil{
		 switch e:=err.(type){
		 case *exec.Error:
			fmt.Println("failed executing",err)
		 case *exec.ExitError:
			fmt.Println("command exit rc=",e.ExitCode())
	     default:
			panic(err)		
		 }
		 
	}
 grepCmd:=exec.Command("grep","hello")
 grepIn,_:=grepCmd.StdinPipe()
 grepOut,_:=grepCmd.StdoutPipe()
 grepCmd.Start()
 grepIn.Write([]byte("hello people \n goodbye grep"))
 grepIn.Close()
 grepBytes,_:=io.ReadAll(grepOut)
 grepCmd.Wait()
 fmt.Println(string(grepBytes))
lsCmd:=exec.Command("bash","-c","ls -a -l -h")
lsOut,err:=lsCmd.Output()
if err !=nil{
	panic(err)
}
fmt.Println(string(lsOut))
 
	

}