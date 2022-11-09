package main

import (
	"fmt"
	gs "github.com/Idiotmann/edge/algorithm"
	"github.com/Idiotmann/edge/model"
)

var GS gs.GS

func main() {

	girlIp := model.GirlIP{"w1", "w2", "w3", "w4"}
	boyIp := model.BoyIP{"m1", "m2", "m3"}
	// 调用pf 函数

	wbglist, _ := model.Compile(boyIp, girlIp)

	for _, v := range wbglist {
		fmt.Printf("Graph长度是=%v \n", len(v.Graph))
		graph := v.Graph
		boyArr1, girlArr1 := gs.SortLike(graph)
		GS.GaleShapley(boyArr1, girlArr1)
		fmt.Println("名字:\t最终匹配对象:")
		for _, v := range boyArr1 {
			fmt.Printf("%v\t%v\n", v.Name, v.Friend)
		}

		fmt.Println("")
		for _, v := range girlArr1 {
			fmt.Printf("%v\t%v\n", v.Name, v.Friend)
		}

		for i, k := range v.Graph {
			fmt.Printf("Graph[i]宽度是=%v \n", len(k))

			for j, m := range k {
				fmt.Printf("Graph[%v][%v]=%v \n", i, j, m)

			}
		}
		fmt.Println(v)
	}
	fmt.Println(wbglist)
	fmt.Printf("长度是%v", len(wbglist))

}
