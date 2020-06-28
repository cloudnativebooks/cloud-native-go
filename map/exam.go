package _map

/*
下面代码存在什么问题？
*/
var FruitColor map[string]string

func AddFruit(name, color string) {
	FruitColor[name] = color
}

/*
panic: assignment to entry in nil map
*/

var StudentScore map[string]int

func GetScore(name string) int {
	score := StudentScore[name]
	return score
}

func GetScoreImproved(name string) int {
	score, ok := StudentScore[name]
	if ok {
		return score
	}

	return -1
}
