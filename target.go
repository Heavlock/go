package main

//
//import (
//	"bytes"
//	"fmt"
//	"strconv"
//	"strings"
//)
//
//func main() {
//	str := "3[dddf4[ssa2[ffd]]]"
//	str = strings.Replace(str, "]", "", 1000)
//	arr := strings.Split(str, "[")
//	var res string
//	var countToMultiply int64
//	for ind, v := range arr {
//
//		if num, err := strconv.ParseInt(string(v[len(v)-1]), 10, 8); err == nil {
//			v = v[:len(v)-1]
//			addString(&res, v, countToMultiply)
//			countToMultiply = num
//		} else {
//			addString(&res, v, countToMultiply)
//			countToMultiply = 1
//		}
//
//	}
//	fmt.Println(res)
//}
//func addString(str *string, addingString string, count int64) {
//	var buffer bytes.Buffer
//
//	buffer.WriteString(*str)
//	var i int64
//	for ; i < count; i++ {
//		buffer.WriteString(addingString)
//	}
//	*str = buffer.String()
//}
