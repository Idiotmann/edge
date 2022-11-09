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
	// 将数据存储按cost排序存储
	wbgList, _ := model.Compile(boyIp, girlIp)
	fmt.Println(wbgList)
	fmt.Printf("长度是%v", len(wbgList))

	for _, v := range wbgList {
		fmt.Printf("Graph长度是=%v \n", len(v.Graph))
		graph := v.Graph
		//创建people数据
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

}
