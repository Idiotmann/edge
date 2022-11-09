package model

//LikePeople根据cost进行排序

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
