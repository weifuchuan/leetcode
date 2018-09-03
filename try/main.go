package main

import (
	"reflect"
	"log"
)

func main(){
	for _, r := range "看kk{"{
		log.Println(reflect.TypeOf(r).Name())
	}
}