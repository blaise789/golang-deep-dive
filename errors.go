package main

import (
	"errors"
	"fmt"
)

// import (
// 	"errors"
// 	"fmt"
// )
// func f(arg int) (int,error){
// 	if(arg==42){
// 		return -1,errors.New("can't work wth 42")

// 	}
// 	return arg+3,nil
// }
// var ErroOutOfTea=fmt.Errorf("no more tea available")
// var ErrorPower=fmt.Errorf("can't boil water")
// func makeTea(arg int ) error{
// 	if arg==2{
// 		return ErroOutOfTea
// 	}	else if arg==4{
// 		return fmt.Errorf("making tea:%w",ErrorPower)

// 	}
// 	return nil
// }
// func main (){
// 	fmt.Println(f(42))
// 	fmt.Println(f(46))
// 	for _,i:=range []int{7,42}{
// 	if r,e:=f(i);e !=nil{
// 		fmt.Println("f failed",e)
// 	}else{

// 		fmt.Println("worked",r)
// 	}

// 	}
// 	for i:=range 5{
// 	  if err:=makeTea(i); err!=nil{
// 		if errors.Is(err,ErroOutOfTea){
// 			fmt.Println("We shoud buy new tea!")
// 		} else if  errors.Is(err,ErrorPower){
// 			fmt.Println("Wow  it is dark")
// 		} else{
// 			fmt.Printf("unknown error:%s\n",err)
// 		}

// 	  }

// 	  fmt.Print(i)
// 	}

// }
type argError struct {
	arg     int
	status  int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {

		return -1, &argError{arg, 500,"can't work with it"}
	}
	return arg + 3, nil
}

func main() {
  _,err:=f(42)
  var ae *argError
  if errors.As(err,&ae){
	fmt.Println(ae.arg)
        fmt.Println(ae.message)

  }else{
	fmt.Println("err doesn't match argError")
  }
}
