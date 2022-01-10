package main

func mainProgram(input []int) int {

    // input := []int{7,1,2,9,7,2,1}

	var reverseInput = reverseInput(input)

	var sequenceArr = sequenceVal(input)
	var sequenceArrReverse = sequenceVal(reverseInput)

	var newArray []int

	for i := 0; i < len(sequenceArr); i++ {
		for j := 0; j < len(sequenceArrReverse); j++ {
			if sequenceArr[i] == sequenceArrReverse[j] {
				newArray = append(newArray, sequenceArr[i])
			}
		}
	}

	var findMax = findMax(newArray)

	return findMax
	
}

func reverseInput(input []int) []int {
    if len(input) == 0 {
        return input
    }
    return append(reverseInput(input[1:]), input[0]) 
}

func sequenceVal(input []int) []int{
	var newArray []int //nill
	for i := 0; i < len(input)-1; i++ {
		if(input[i+1] - input[i] == 1) {
			newArray = append(newArray, input[i], input[i+1])
		} 
	}
	return newArray
}

func findMax(a []int) (max int) {
	max = a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}
