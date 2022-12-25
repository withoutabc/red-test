// 又臭又长
package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func IfSum(i int, sum int) (a int) {
	if sum == i {
		fmt.Println(i)
	}
	return
}
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
func Count(a int, b int) (result int) {
	if b == 0 {
		result = 1
		return
	}
	result = a * Count(a, b-1)
	return
}
func Sleep() {
	time.Sleep(1e9)
}
func main() {
	wg.Add(6)
	go func() {
		for i := 1; i < 10; i++ {
			num0 := strconv.Itoa(i)
			num1, err := strconv.Atoi(num0[0:1])
			CheckErr(err)
			sum := Count(num1, 1)
			IfSum(i, sum)
		}
		wg.Done()
	}()

	go func() {
		for i := 11; i < 100; i++ {
			num0 := strconv.Itoa(i)
			num1, err := strconv.Atoi(num0[0:1])
			CheckErr(err)
			num2, err1 := strconv.Atoi(num0[1:2])
			CheckErr(err1)
			sum := Count(num1, 2) + Count(num2, 2)
			IfSum(i, sum)
		}
		wg.Done()
	}()
	go func() {
		for i := 101; i < 1000; i++ {
			num0 := strconv.Itoa(i)
			num1, err := strconv.Atoi(num0[0:1])
			CheckErr(err)
			num2, err1 := strconv.Atoi(num0[1:2])
			CheckErr(err1)
			num3, err2 := strconv.Atoi(num0[2:3])
			CheckErr(err2)
			sum := Count(num1, 3) + Count(num2, 3) + Count(num3, 3)
			IfSum(i, sum)
		}
		wg.Done()
	}()
	go func() {
		for i := 1001; i < 10000; i++ {
			num0 := strconv.Itoa(i)
			num1, err := strconv.Atoi(num0[0:1])
			CheckErr(err)
			num2, err1 := strconv.Atoi(num0[1:2])
			CheckErr(err1)
			num3, err2 := strconv.Atoi(num0[2:3])
			CheckErr(err2)
			num4, err3 := strconv.Atoi(num0[3:4])
			CheckErr(err3)
			sum := Count(num1, 4) + Count(num2, 4) + Count(num3, 4) + Count(num4, 4)
			IfSum(i, sum)
		}
		wg.Done()
	}()
	go func() {
		for i := 10001; i < 100000; i++ {
			num0 := strconv.Itoa(i)
			num1, err := strconv.Atoi(num0[0:1])
			CheckErr(err)
			num2, err1 := strconv.Atoi(num0[1:2])
			CheckErr(err1)
			num3, err2 := strconv.Atoi(num0[2:3])
			CheckErr(err2)
			num4, err3 := strconv.Atoi(num0[3:4])
			CheckErr(err3)
			num5, err4 := strconv.Atoi(num0[4:5])
			CheckErr(err4)
			sum := Count(num1, 5) + Count(num2, 5) + Count(num3, 5) + Count(num4, 5) + Count(num5, 5)
			IfSum(i, sum)
		}
		wg.Done()
	}()
	go func() {
		for i := 100001; i < 1000000; i++ {
			num0 := strconv.Itoa(i)
			num1, err := strconv.Atoi(num0[0:1])
			CheckErr(err)
			num2, err1 := strconv.Atoi(num0[1:2])
			CheckErr(err1)
			num3, err2 := strconv.Atoi(num0[2:3])
			CheckErr(err2)
			num4, err3 := strconv.Atoi(num0[3:4])
			CheckErr(err3)
			num5, err4 := strconv.Atoi(num0[4:5])
			CheckErr(err4)
			num6, err5 := strconv.Atoi(num0[5:6])
			CheckErr(err5)
			sum := Count(num1, 6) + Count(num2, 6) + Count(num3, 6) + Count(num4, 6) + Count(num5, 6) + Count(num6, 6)
			IfSum(i, sum)
		}
		wg.Done()
	}()

	wg.Wait()
}
