package additional

import (
	"fmt"
	"testing"
)

// 给定一个数字类型的字符串，想办法截取连续在一起的和正好等于字符串长度的字符串，如123222，可获得123，222这两个字符串
func TestAAA(t *testing.T) {
	//var s = "123222"
	var s = "345651"
	left := 0
	//reght := 0
	result := make([]string, 0)
	sum := 0
	length := len(s)
	for i := 0; i < length; i++ {
		c := s[i]
		val := int(c - '0') //
		sum += val
		if sum > len(s) { //爆掉
			sum -= int(s[left] - '0') //减去最开头的值
			left++                    //左坐标要右移
		} else if sum == length { //正好等于
			re := s[left : i+1]         //要包含一下当前i
			result = append(result, re) //left开始，到i结束，这段都是
			sum -= int(s[left] - '0')   //减去最开头的值,
			left++
		}
	}
	fmt.Println(result)
}
