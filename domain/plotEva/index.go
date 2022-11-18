package plotEva

// Index 评估指标
type Index struct {
	//平均执行时间
	WorkingTimeCost float64
	//调度时间
	SchedulingTimeCost int64
	//数据大小[2]int {fromIp, toIp}
	Size [2]int
	//迭代次数
	Interactions int
}
