package main

import (
	"fmt"
	"sort"
	"strconv"
)

type intList []int

func (s intList) Len() int           { return len(s) }
func (s intList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s intList) Less(i, j int) bool { return s[i] > s[j] }

// 阶乘
func factorial(n int) int64 {
	var fc, i int64
	fc = 1 // 0! = 1
	for i = 1; i <= (int64)(n); i++ {
		fc = fc * i
	}
	return fc
}

// 组合数
func combo(total, pick int) int64 {
	//up := factorial(total)
	//down := factorial(pick) * factorial(total - pick)
	//result := up/down
	return factorial(total) / factorial(pick) * factorial(total-pick)
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
	fmt.Println(str)
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
				fmt.Printf("i:%d %s\n", i, str)
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

// 将ele复制后添加到arr中，返回新的数组
func addTo(arr [][]int, ele []int) [][]int {
	newEle := make([]int, len(ele))
	copy(newEle, ele)
	arr = append(arr, newEle)
	return arr
}
func moveOneToLeft(leftNums []int) {
	//计算有几个1
	sum := 0
	for i := 0; i < len(leftNums); i++ {
		if leftNums[i] == 1 {
			sum++
		}
	}
	//将前sum个改为1，之后的改为0
	for i := 0; i < len(leftNums); i++ {
		if i < sum {
			leftNums[i] = 1
		} else {
			leftNums[i] = 0
		}
	}
}

// 根据索引号数组得到元素数组
func findNumsByIndexs(nums []int, indexs [][]int) [][]int {
	if len(indexs) == 0 {
		return [][]int{}
	}
	result := make([][]int, len(indexs))
	for i, v := range indexs {
		line := make([]int, 0)
		for j, v2 := range v {
			if v2 == 1 {
				line = append(line, nums[j])
			}
		}
		result[i] = line
	}
	return result
}

func main() {
	GetAllCombination(4, 2)
	//for _, str := range strArr {
	//	fmt.Println(str)
	//}
	indexs := make(intList, 3)
	indexs[0] = 1
	indexs[1] = 2
	indexs[2] = 3
	sort.Sort(indexs)
	for _, item := range indexs {
		fmt.Println(item)
	}
}
