package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"unicode"
)

func spite(str string) (ret []string, err error) {
	if len(str) != 4 {
		return nil, errors.New("请输出4个数字")
	}

	ret = make([]string, 4)
	for index, s := range str {

		if !unicode.IsNumber(s) {
			return nil, errors.New("请输出正确的数字")
		}

		pos := getNumberPos(ret, string(s))

		if pos == index {
			return nil, errors.New("请勿输入重复数字")
		}

		ret[index] = string(s)
	}

	return
}

func getNumberPos(numList []string, v string) int {
	for index, n := range numList {
		if n == v {
			return index
		}

	}

	return -1
}

func genNumList() (ret []string) {

	for len(ret) < 4 {
		n := strconv.Itoa(int(rand.Int31n(10)))

		if pos := getNumberPos(ret, n); pos != -1 {
			continue
		}

		ret = append(ret, n)
	}

	return
}

func compare(answer, user []string) (ret string) {

	var numA int
	var numB int
	for index, u := range user {
		a := answer[index]

		if a == u {
			numA++
		} else if pos := getNumberPos(answer, u); pos != -1 {
			numB++
		}
	}

	return fmt.Sprintf("%dA%dB", numA, numB)

}

func main() {

	rand.Seed(time.Now().Unix())

	answer := genNumList()

	fmt.Println("请输入4位连续数字答案")

	for i := 1; i <= 10; i++ {

	Reinput:
		var input string
		fmt.Scan(&input)

		userList, err := spite(input)
		if err != nil {
			fmt.Println(err)
			goto Reinput
		}

		result := compare(answer, userList)
		if result == "4A0B" {
			fmt.Println("正确！")
			return
		}
		fmt.Printf("%d: %s\n", i, result)

	}

	fmt.Println("答案:", answer)
}
