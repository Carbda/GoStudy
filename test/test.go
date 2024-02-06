package main

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"time"
)

func main() {
	//fmt.Println(GetAllCombination(3, 2))
	//testGetPvSameLevel()
	//testLenAndNil(nil)
	//testMapNil()
	testMin()
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

// 组合算法(从nums中取出m个数)
func GetAllCombination(n int, m int) []string {
	if m < 1 || m > n {
		fmt.Println("Illegal argument. Param m must between 1 and len(nums).")
		return []string{}
	}
	//保存最终结果的数组，总数直接通过数学公式计算
	result := make([]string, 0)
	//保存每一个组合的索引的数组，1表示选中，0表示未选中
	indexs := make(intList, n)
	for i := 0; i < n; i++ {
		if i < m {
			indexs[i] = 1
		} else {
			indexs[i] = 0
		}
	}
	//第一个结果
	var str string
	for i := 0; i < len(indexs); i++ {
		str += strconv.Itoa(indexs[i])
	}
	result = append(result, str)
	for {
		find := false
		//每次循环将第一次出现的 1 0 改为 0 1，同时将左侧的1移动到最左侧
		for i := 0; i < n-1; i++ {
			if indexs[i] == 1 && indexs[i+1] == 0 {
				find = true
				indexs.Swap(i, i+1)
				if i > 1 {
					sort.Sort(indexs[:i])
				}
				str = ""
				for i := 0; i < len(indexs); i++ {
					str += strconv.Itoa(indexs[i])
				}
				result = append(result, str)
				break
			}
		}
		//本次循环没有找到 1 0 ，说明已经取到了最后一种情况
		if !find {
			break
		}
	}
	return result
}

type intList []int

func (s intList) Len() int           { return len(s) }
func (s intList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s intList) Less(i, j int) bool { return s[i] > s[j] }

type ParseGroupView struct {
	ParseGroupId int64 `json:"pg_id"`
	ViewId       int32 `json:"view_id"`
}

// 主机组打标
type HostGroupTag struct {
	HostGroupUuid string `json:"host_group_uuid"`
}

type resourcePoolGenerationSvc struct {
	sameLevel map[ParseGroupView]map[string]*HostGroupTag //map[pv]map[hgUuid]*HostGroupTag
}

func testGetPvSameLevel() {
	p1 := ParseGroupView{ParseGroupId: 1, ViewId: 10}
	p2 := ParseGroupView{ParseGroupId: 1, ViewId: 10}
	rps := &resourcePoolGenerationSvc{
		sameLevel: make(map[ParseGroupView]map[string]*HostGroupTag),
	}
	rps.sameLevel[p1] = make(map[string]*HostGroupTag)
	rps.sameLevel[p1]["tmpStr"] = &HostGroupTag{HostGroupUuid: "tmpId"}
	if res, ok := rps.sameLevel[p2]; ok {
		fmt.Println(res)
	}
	if p1 == p2 {
		fmt.Println("same")
	}
}

func testLenAndNil(arrStr []string) {
	if len(arrStr) > 0 {
		fmt.Println("> 0")
	} else {
		fmt.Println("<= 0")
	}
}

func testMapNil() {
	testM := make(map[int64]int64)
	if _, exist := testM[1]; exist {
		fmt.Println("exist!")
	} else {
		fmt.Println("no exist!")
	}
	testM[1] = 1
	if _, exist := testM[1]; exist {
		fmt.Println("exist!")
	} else {
		fmt.Println("no exist!")
	}
}

func testMin() {
	a := int64(70)
	b := int64(12)
	c := min(a, b)
	fmt.Println(c)
}
