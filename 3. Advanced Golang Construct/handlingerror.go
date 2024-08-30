package main

import (
	"errors"
	"fmt"
)


type MyError struct {
	Msg string 
	Code int
}

func (e *MyError) Error ()string {
	return fmt.Sprintf("%d-%s",e.Code,e.Msg)
}
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a/b,nil
}

func main() {
	_,err := divide(4,0)

	if err != nil {
		fmt.Println(err)
	}

	myerr := MyError{
		Msg: "forbidden",
		Code:403,
	}
	fmt.Println(myerr.Error())

}