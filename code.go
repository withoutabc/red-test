// 不考虑时间的情况还有点问题，思路就是把品质加在一个数组里，排序以后通过相邻数字的比较判断有几种
// 考虑时间的情况 应该是要对t在24小时以内的进行判断，把在24小时以内的组成数组，然后同上。t和后面的数据可能要map
package main

import "fmt"

func main() {
	var n, t, k, x int
	var s1 []int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		fmt.Scan(&k)
		for j := 0; j < k; j++ {
			fmt.Scan(&x)
			s1 = append(s1, x)
		}
		fmt.Println(s1)
		fmt.Println(Count(s1))
	}
}
func Count(s []int) int {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	var x = 1
	for i := 0; ; i++ {
		if i+2 >= len(s) {
			return x + 1
		}
		if s[i] != s[i+1] {
			x++
		}
	}

}
