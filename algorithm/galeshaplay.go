package gs

import (
	"fmt"
	"github.com/Idiotmann/edge/model"
)

type GS struct {
}

type People struct {
	Name        string   // 名字
	LikePeople  []string // 喜好列表
	CurrentLike int      // 后面算法记录当前表白对象时使用
	Friend      string   // 当前匹配对象
}

func SortLike(graph model.Graph) (boyArr []*People, girlArr []*People) {
	e := len(graph)
	m := e - len(graph[0])
	// n := e - m
	//  var boyArr1, girlArr1 []*People
	boyArr1 := make([]*People, 0)
	girlArr1 := make([]*People, 0)
	for i, k := range graph {
		if i < m {
			// boyArr[i].Name = append(boyArr[i].Name,k[i].Src)
			// boyArr1[i].Name = k[i].Src
			// 升序排列
			likepeople := make([]string, 0)
			for _, v := range k {
				people := v.Des
				likepeople = append(likepeople, people)
				// boyArr1[i].LikePeople = append(boyArr1[i].LikePeople, people)
			}
			// boyArr1[i] = &People{
			boyArr1 = append(boyArr1, &People{
				Name:        k[0].Src,
				LikePeople:  likepeople,
				CurrentLike: 0,
				Friend:      "",
			})
			// 要对邻接矩阵的按照cost标准进行升序排列

		} else {
			likepeople := make([]string, 0)
			for _, v := range k {
				people := v.Des
				likepeople = append(likepeople, people)
				// boyArr1[i].LikePeople = append(boyArr1[i].LikePeople, people)
			}
			girlArr1 = append(girlArr1, &People{
				Name:        k[0].Src,
				LikePeople:  likepeople,
				CurrentLike: 0,
				Friend:      "",
			})
		}

	}
	return boyArr1, girlArr1
}
func (gs *GS) GaleShapley(boyArr []*People, girlArr []*People) {
	// func (gs *GS) GaleShapley(graph Graph) {
	// E := len(graph) //一维数组的元素个数
	// boyArr := make([]*People, E)
	// girlArr := make([]*People, E)

	for {
		// 找到一个没有对象, 且未全部表白的男生
		var searchBoy *People
		for _, boy := range boyArr {
			if boy.Friend != "" { // 当前男孩已经有对象了
				continue
			}
			// 男孩向所有女生表白过了
			if boy.CurrentLike >= len(boy.LikePeople) {
				fmt.Printf("%v\n", boy.CurrentLike)
				continue
			}
			searchBoy = boy
			break
		}
		if searchBoy == nil { // 已经全部有对象了, 结束
			break
		}
		// 男生向女生依次表白
		var i int
		for i := searchBoy.CurrentLike; i < len(searchBoy.LikePeople); i++ {
			girlName := searchBoy.LikePeople[i]
			// 找到这个女孩
			girl := searchPeople(girlArr, girlName)
			if girl == nil { // 习惯了, 判下空
				continue
			}
			if girl.Friend == "" { // 若女孩没有对象, 则直接配对
				girl.Friend = searchBoy.Name
				searchBoy.Friend = girl.Name
				break
			} else { // 若女孩有对象, 看下 girl 更喜欢谁
				searchBoyIdx := searchNameIndex(girl.LikePeople, searchBoy.Name)
				girlFriendIdx := searchNameIndex(girl.LikePeople, girl.Friend)
				if girlFriendIdx < searchBoyIdx { // 保持当前
					continue
				} else { // 重新组队
					girlFriend := searchPeople(boyArr, girl.Friend)
					if girlFriend != nil { // 分手了
						girlFriend.Friend = ""
						girlFriend.CurrentLike++
					}
					girl.Friend = searchBoy.Name
					searchBoy.Friend = girl.Name
					break
				}
			}

		}
		searchBoy.CurrentLike = i
	}
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
