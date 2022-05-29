package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	notSorted := make([]int, 0)
	for i := range input {
		notSorted = append(notSorted, i)
	}
	sort.Ints(notSorted)
	for _, v := range notSorted {
		result = append(result, input[v])
	}
	return
}
