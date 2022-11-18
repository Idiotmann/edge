package gs

import (
	"github.com/Idiotmann/edge/model"
	"github.com/Idiotmann/edge/plotEva"
	"time"
)

type GS struct {
}

type People struct {
	Name        string    // 名字
	LikePeople  []string  // 喜好列表
	Cost        []float64 //想要记录
	CurrentLike int       // 后面算法记录当前表白对象时使用
	Friend      friend    // 当前匹配对象
}
type friend struct {
	Name string
	Cost float64
}

// SortLike 找boy(m)的people结构体，girl(w)的people结构体
func SortLike(graph model.Graph) (boyArr []*People, girlArr []*People) {
	e := len(graph)
	// len(boyArr)
	m := e - len(graph[0])
	boyArr1 := make([]*People, 0)
	girlArr1 := make([]*People, 0)
	for i, k := range graph {
		if i < m {
			boyLikePeople := make([]string, 0)
			boyFriend := friend{}
			boyCost := make([]float64, 0)
			for _, v := range k {
				boyLikePeople = append(boyLikePeople, v.Des)
				boyCost = append(boyCost, v.Cost)
			}
			boyArr1 = append(boyArr1, &People{
				Name:        k[0].Src,
				LikePeople:  boyLikePeople,
				Cost:        boyCost,
				CurrentLike: 0,
				Friend:      boyFriend,
			})

		} else {
			girlLikePeople := make([]string, 0)
			girlCost := make([]float64, 0)
			girlFriend := friend{}
			for _, v := range k {

				girlLikePeople = append(girlLikePeople, v.Des)
				girlCost = append(girlCost, v.Cost)
			}
			girlArr1 = append(girlArr1, &People{
				Name:        k[0].Src,
				LikePeople:  girlLikePeople,
				Cost:        girlCost,
				CurrentLike: 0,
				Friend:      girlFriend,
			})
		}

	}
	return boyArr1, girlArr1
}
func (gs *GS) GaleShapley(boyArr []*People, girlArr []*People) plotEva.Index {
	m := len(boyArr)
	n := len(girlArr)
	size := [2]int{m, n}
	//获取当前时间信息
	startTs := time.Now()

	//迭代次数
	var iteration int
	var cost float64
	for {
		// 找到一个没有对象, 且未全部表白的男生
		var searchBoy *People
		for _, boy := range boyArr {
			if boy.Friend.Name != "" || boy.CurrentLike >= len(boy.LikePeople) { //已经有对象了或者已经全部表白过了
				continue
			}
			searchBoy = boy
			break
			//这里直接结束是因为前面还有个for，下面执行结束会再次寻找boy
		}
		// 已经全部有对象了, 结束
		if searchBoy == nil {
			break
		}
		// 男生向女生依次表白

		for i := searchBoy.CurrentLike; i < len(searchBoy.LikePeople); i++ {
			iteration++
			girlName := searchBoy.LikePeople[i]
			boyCost := searchBoy.Cost[i]
			//在girl-arr中 找到这个女孩
			girl := searchPeople(girlArr, girlName)
			if girl == nil {
				continue
			}
			if girl.Friend.Name == "" { // 若女孩没有对象, 则直接配对
				girl.Friend.Name = searchBoy.Name
				girl.Friend.Cost = boyCost
				searchBoy.Friend.Name = girlName
				searchBoy.Friend.Cost = boyCost
				searchBoy.CurrentLike = i
				break
			} else {
				// 若女孩有对象, 看下 girl 更喜欢谁
				//找到girl的searchBoy的偏好下标和当前friend的idx,谁在前就把谁当朋友
				searchBoyIdx := searchNameIndex(girl.LikePeople, searchBoy.Name)
				girlFriendIdx := searchNameIndex(girl.LikePeople, girl.Friend.Name)
				if girlFriendIdx < searchBoyIdx { // 保持当前
					continue
				} else {
					// 通过girl.Friend的名字找到这个男孩的结构体，把他的friend变成空的
					//让他继续按照偏好列表向下一个表白
					girlFriend := searchPeople(boyArr, girl.Friend.Name)
					if girlFriend != nil {
						girlFriend.Friend.Name = "" //把她的friend（boy）的Friend名字变成空的
						girlFriend.Friend.Cost = 0  //把她的friend（boy）的Friend的cost变成空的
						girlFriend.CurrentLike++    //让他继续按照偏好列表向下一个表白
					}
					//把girl把当前男孩变成friend，男孩把当前girl变成friend,然后退出不表白了
					girl.Friend.Name = searchBoy.Name
					girl.Friend.Cost = boyCost
					searchBoy.Friend.Name = girl.Name
					searchBoy.Friend.Cost = boyCost
					searchBoy.CurrentLike = i
					break
				}
			}
		}

	}
	for _, boy := range boyArr {
		cost += boy.Friend.Cost //计算总cost
	}
	//调度执行时间
	d := time.Now().Sub(startTs).Microseconds()                                                                          //
	return plotEva.Index{Size: size, Interactions: iteration, WorkingTimeCost: cost / float64(m), SchedulingTimeCost: d} // TOD
}
func searchPeople(peopleArr []*People, name string) *People {
	for _, people := range peopleArr {
		if people.Name == name {
			return people
		}
	}
	return nil
}

func searchNameIndex(nameArr []string, name string) int {
	for i, tmpName := range nameArr {
		if tmpName == name {
			return i
		}
	}
	return -1
}

//Function to get the max and min values of a list
