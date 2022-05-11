package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	var i int = 0
	key := make([]int, len(input))
	for k := range input {
		key[i] = k
		i++
	}
	sort.Ints(key)
	for _, b := range key {
		result = append(result, input[b])

	}
	return result
}
