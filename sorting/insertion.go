package sorting

// rate of growth
// Θ(n2)
func insertion(listToSort []int) []int {

	for i := 1; i < len(listToSort); i++ {
		for j := 0; j < i; j++ {
			if listToSort[j] > listToSort[i] {
				temp := listToSort[i]
				copy(listToSort[j+1:i+1], listToSort[j:i])
				listToSort[j] = temp
				break
			}
		}
	}

	return listToSort
}
