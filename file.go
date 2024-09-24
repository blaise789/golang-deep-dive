package main


/*import (
	"bufio"
	"fmt"
	"io"
	"os"
)
func check(e error){
	if e !=nil{
		panic(e)
	}
}
func main(){
	p:=fmt.Println
data,err:=os.ReadFile("./hello.txt")
check(err)
p(string(data))
f,err:=os.Open("./hello.txt")
check(err)
b1:=make([]byte,5)
p(b1)
// n1,err:=f.Read(b1)
check(err)
// fmt.Printf("%d bytes:%s\n",n1,string(b1))
o2,err:=f.Seek(6,io.SeekStart)
check(err)
b2:=make([]byte,2)
n2,err:=f.Read(b2)
fmt.Printf("%d bytes @ %d\n",n2,o2)
fmt.Printf("%v\n",string(b2[:n2]))
_,err=f.Seek(4,io.SeekCurrent)
check(err)
o3,err:=f.Seek(6,io.SeekStart)
b3:=make([]byte,2)
n3,err:=io.ReadAtLeast(f,b3,2)
check(err)
fmt.Printf("%d bytes @ %d:%s\n",n3,o3,string(b3))

_,err=f.Seek(0,io.SeekStart)


check(err)
r4:=bufio.NewReader(f)
b4,err:=r4.Peek(5)
check(err)
fmt.Printf("5 bytes:%s\n",string(b4))
f.Close()
}*/