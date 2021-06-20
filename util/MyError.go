package util

import "os"

func Errors(err error){
	if err!=nil{
		println(err.Error())
		os.Exit(0)
	}
}
