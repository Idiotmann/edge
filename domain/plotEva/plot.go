package plotEva

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

//func Plot(index Index) {
//	p := plot.New()
//	p.Title.Text = "迭代次数随大小变化"
//	p.X.Label.Text = "fromIP"
//	p.Y.Label.Text = "toIP"
//	err := plotutil.AddLinePoints(p, "迭代", randomPoints(15))
//
//}

func Plot(index Index) {
	println("Size:", index.Size[0], index.Size[1])           // [2]int {fromIp, toIp}
	println("Interactions:", index.Interactions)             //迭代次数
	println("WorkingTimeCost:", index.WorkingTimeCost)       //平均执行时间
	println("SchedulingTimeCost:", index.SchedulingTimeCost) //调度时间
	p := plot.New()
	p.Title.Text = "迭代次数随fromip变化次数" //
	p.X.Label.Text = "fromIp"
	p.Y.Label.Text = "international"
	//参数类型 plot string xys 其中xys是 []xy   xy是struct{x,y float64}
	err := plotutil.AddLinePoints(p,
		"迭代次数", Points(index))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func Points(index Index) plotter.XYer {
	pts := make(plotter.XYs, 1)
	pts[0].X = float64(index.Size[0])
	pts[0].Y = float64(index.Interactions)
	return pts
}
