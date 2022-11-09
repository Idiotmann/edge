package model

import (
	"math/rand"
	"sort"
)

//WBG 是图的邻接矩阵，它存储了地址和消耗
type WBG struct {
	// Type  pb.Basic_TaskType
	Graph Graph
}

type Graph [][]WeightedEdge

// WeightedEdge 包含边的源地址、目标地址和消耗
type WeightedEdge struct {
	Src  string
	Des  string
	Cost float64
}

type PredFunc func(from string, to string, net GirlIP) float64

// var emptyWeightedEdge = WeightedEdge{
// 	Src:  "-1",
// 	Des:  "-1",
// 	Cost: -1,
// }

type GirlIP []string
type BoyIP []string

// Compile 创建矩阵,存储数据，其中每一行按照cost排列
// func Compile(m BoyIP, net GirlIP,pf PredFunc) ([]WBG, error) {
func Compile(m BoyIP, net GirlIP) ([]WBG, error) {
	var wbgList []WBG
	// 先定义Cost用来测试

	var graph Graph
	for _, fromIP := range m {
		var edges []WeightedEdge
		for _, toIP := range net {
			if len(toIP) == 0 {
				continue
			}
			edges = append(edges, WeightedEdge{
				Src:  fromIP,
				Des:  toIP,
				Cost: rand.Float64(),
			})
		}
		graph = append(graph, edges)
	}

	for _, fromIP := range net {
		var edges []WeightedEdge
		for _, toIP := range m {
			edges = append(edges, WeightedEdge{
				Src:  fromIP,
				Des:  toIP,
				Cost: rand.Float64(),
			})

		}
		graph = append(graph, edges)
	}

	for _, v := range graph {
		sort.Sort(LikePeople(v))
	}
	wbgList = append(wbgList, WBG{
		Graph: graph,
	})
	return wbgList, nil
}
