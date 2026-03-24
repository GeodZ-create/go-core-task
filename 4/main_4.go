package main

func main() {
}

func NewSlice(first []string, second []string) []string {
	result := make([]string, 0)
	sliceTwoMap := make(map[string]bool)

	for _, v := range second {
		sliceTwoMap[v] = true
	}

	for _, v := range first {
		_, ok := sliceTwoMap[v]
		if !ok {
			result = append(result, v)
		}
	}

	return result
}
