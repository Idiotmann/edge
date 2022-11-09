package model

import (
	"math/rand"
	"sort"
)

//WBG 是图的邻接矩阵，

type WBG struct {
	// Type  pb.Basic_TaskType
	Graph Graph
}

type Graph [][]WeightedEdge

// WeightedEdge contains the source, destination and Cost of edges.WeightedEdge 包含边的源、目标和成本。
type WeightedEdge struct {
	Src  string
	Des  string
	Cost float64
}
type LikePeople []WeightedEdge

func (array LikePeople) Len() int {
	return len(array)
}

func (array LikePeople) Less(i, j int) bool {
	return array[i].Cost < array[j].Cost //从小到大， 若为大于号，则从大到小
}

func (array LikePeople) Swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}

type PredFunc func(from string, to string, net GirlIP) float64

// var emptyWeightedEdge = WeightedEdge{
// 	Src:  "-1",
// 	Des:  "-1",
// 	Cost: -1,
// }

type GirlIP []string
type BoyIP []string

// var boyip, girlip []string

// func Compile(m ExpandedTaskMap, net comnet.NetV4, pf PredFunc) ([]WBG, error) {
// func Compile(m BoyIP, net GirlIP,pf PredFunc) ([]WBG, error) {
func Compile(m BoyIP, net GirlIP) ([]WBG, error) {
	var wbgList []WBG
	// 先定义cost用来测试
	var cost float64
	var graph Graph
	for _, fromIP := range m {
		var edges []WeightedEdge
		for _, toIP := range net {
			edges = append(edges, WeightedEdge{
				Src:  fromIP,
				Des:  toIP,
				Cost: rand.Float64(),
			})
			// graph = append(graph, edges)
			cost++
		}
		graph = append(graph, edges)
	}
	// } else {
	for _, fromIP := range net {
		var edges []WeightedEdge
		// for j := 0; j < N; j++ {
		for _, toIP := range m {

			edges = append(edges, WeightedEdge{
				Src:  fromIP,
				Des:  toIP,
				Cost: rand.Float64(),
			})
			// graph = append(graph, edges)

		}
		// }
		graph = append(graph, edges)
	}

	for _, v := range graph {
		sort.Sort(LikePeople(v))
	}
	wbgList = append(wbgList, WBG{
		// Type:  key,
		Graph: graph,
	})
	// }
	return wbgList, nil
}
