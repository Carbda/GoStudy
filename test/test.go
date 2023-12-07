package main

import (
	"context"
	"fmt"
	"time"
)

func main0() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}

func longestPalindrome(s string) string {
	maxLen := 1
	dp := make([][]int, len(s))
	begin := 0
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	for L := 2; L <= len(s); L++ {
		for i := 0; i < len(s); i++ {
			j := i + L - 1
			if j >= len(s) {
				break
			}
			if s[i] != s[j] {
				dp[i][j] = 0
			} else {
				if j-i <= 2 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] == 1 && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

func convert(s string, numRows int) string {
	strArr := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		strArr[i] = make([]byte, 0)
	}
	p := 0
	for p < len(s) {
		for i := 0; i < numRows && p < len(s); i++ {
			strArr[i] = append(strArr[i], s[p])
			p++
		}
		for i := numRows - 2; i > 0 && p < len(s); i-- {
			strArr[i] = append(strArr[i], s[p])
			p++
		}
	}
	// 拼接字符串
	ansStr := ""
	for i := 0; i < numRows; i++ {
		tmpStr := string(strArr[i])
		ansStr = ansStr + tmpStr
	}
	return ansStr
}

func tmpFun() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	go tmpFun2(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func tmpFun2(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
