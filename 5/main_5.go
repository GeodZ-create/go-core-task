package main

func main() {

}

func MatchesSlice(first []int, second []int) (bool, []int) {
	var result []int
	sliceTwoMap := make(map[int]bool)
	sliceResultMap := make(map[int]bool)
	for _, v := range second {
		sliceTwoMap[v] = true
	}
	for _, v := range first {
		_, ok := sliceTwoMap[v]
		if ok {
			_, ok := sliceResultMap[v]
			if !ok {
				result = append(result, v)
				sliceResultMap[v] = true
			}
		}
	}
	if len(result) == 0 {
		return false, []int{}
	}
	return true, result
}
