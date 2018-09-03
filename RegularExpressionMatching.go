package leetcode

import (
	"strings" 
)

func IsMatch(s string, p string) bool  {
	//var fsm [3][3][]func(byte)bool
	return true
}

type state int 
const(
	t state = iota // true
	f  // false
	e  // end 
	
)

type controller = func(reader strings.Reader)()

func build(action string, values ...byte)func(byte)bool   {
	switch action {
	case "E":
		return func (c byte)bool   {
			if c==values[0]{
				return true 
			}
			return false 
		}
	case "N":
		return func (_ byte)bool   {
			return true 
		}
	}
	return func (c byte)bool   {
		return false 
	}
}