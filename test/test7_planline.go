package main

import "fmt"

func main7() {
	view5MinsBws := []int64{2000, 3000, 3000, 4000, 1000}
	tuihuo(1500, view5MinsBws)
}

func tuihuo(oriPlanLine int64, view5MinsBws []int64) bool {

	planLine := oriPlanLine

	planLine += 1000

	//算提线后这段的面积，生成tmpView5MinsBws
	tmpView5MinsBws := make([]int64, len(view5MinsBws))
	for i, bw := range view5MinsBws {
		if bw <= oriPlanLine {
			tmpView5MinsBws[i] = 0
		} else if bw > planLine {
			tmpView5MinsBws[i] = planLine - oriPlanLine
		} else if bw < planLine && bw > oriPlanLine {
			tmpView5MinsBws[i] = bw - oriPlanLine
		}
	}

	fmt.Println(tmpView5MinsBws)

	//拿tmpView5MinsBws去算非限速区的planLine，看有没有降1g以上
	calPlanLine()

	return false
}

func calPlanLine() int64 {
	return 2000
}
