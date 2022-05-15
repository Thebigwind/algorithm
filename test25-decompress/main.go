package main

import (
	"fmt"
	"strconv"
)

/*
字符串解压缩
*/

func main(){
decompressString("a4v5gj2ld3n7")
}
//a4v5gj2ld3n7
func decompressString(S string) string {
	result := ""
	previous := ""
	for _,v := range S{
		temp := string(v)
		num,err := strconv.Atoi(temp)
		if err != nil{
			previous = temp
			result = result +previous
		}else{
			tempStr := ""
			for i:= 0; i<num;i++{
				tempStr = tempStr + previous
			}
			result = result + tempStr
		}
	}
	fmt.Println(result)
	return result
}
