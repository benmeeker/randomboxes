package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	var randomSuccess []int
	var loopSuccess []int
	randomComplete := 0
	loopComplete := 0
	randomFailure := 0
	loopFailure := 0
	boxCount := 100
	prisonerCount := 100
	allowedTries := 50

	for i := 0; i < 1000; i++ {
		boxMap := makeRandomMap(boxCount)
		prisonerArray := makeRandomArray(prisonerCount)

		randomTrial := randomCheck(boxMap, prisonerArray, allowedTries, boxCount)
		if randomTrial == 100 {
			log.Println("All prisoners found their boxes! - Random Trial")
			randomComplete++
		}
		if randomTrial == 0 {
			log.Println("No prisoners found their boxes... Awkward - Random Trial")
			randomFailure++
		}
		randomSuccess = append(randomSuccess, randomTrial)

		loopTrial := loopCheck(boxMap, prisonerArray, allowedTries, boxCount)
		if loopTrial == 100 {
			log.Println("All prisoners found their boxes! -Loop Trial")
			loopComplete++
		}
		if loopTrial == 0 {
			log.Println("No prisoners found their boxes... Awkward - Loop Trial")
			loopFailure++
		}
		loopSuccess = append(loopSuccess, loopTrial)
	}
	randomAverage := (sumArray(randomSuccess) / len(randomSuccess))
	loopAverage := (sumArray(loopSuccess) / len(loopSuccess))

	log.Println("Random average success: " + fmt.Sprint(randomAverage) + "\nRandom full completions: " + fmt.Sprint(randomComplete) + "\nRandom complete failures: " + fmt.Sprint(randomFailure) + "\n\n")
	log.Println("Loop average success: " + fmt.Sprint(loopAverage) + "\nLoop full completions: " + fmt.Sprint(loopComplete) + "\nLoop complete failures: " + fmt.Sprint(loopFailure))
}

func makeRandomMap(count int) map[int]int {
	boxMap := make(map[int]int)
	boxNumbers := makeRandomArray(count)
	boxContents := makeRandomArray(count)

	for i := 0; i < count; i++ {
		boxMap[boxNumbers[i]] = boxContents[i]
	}

	return boxMap
}

func makeRandomArray(count int) []int {
	var randomArray = make([]int, count)
	for i := 0; i < count; i++ {
		randomArray[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(randomArray), func(i, j int) { randomArray[i], randomArray[j] = randomArray[j], randomArray[i] })

	return randomArray
}

func contains(arr []int, num int) bool {
	for _, i := range arr {
		if i == num {
			return true
		}
	}
	return false
}

func sumArray(arr []int) int {
	result := 0
	for _, i := range arr {
		result += i
	}
	return result
}

func randomCheck(boxMap map[int]int, prisonerArray []int, allowedTries int, boxCount int) int {
	var success int
	for i := 0; i < len(prisonerArray); i++ {
		prisoner := prisonerArray[i]
		var openBoxes []int
		for i := 0; i < allowedTries; i++ {
			rand.Seed(time.Now().UnixNano())
			box := rand.Intn(boxCount-0) + 0

			if contains(openBoxes, box) {
				i = i - 1
				continue
			} else if prisoner == boxMap[i] {
				openBoxes = append(openBoxes, box)
				success++
				continue
			}
			openBoxes = append(openBoxes, box)
		}
	}
	return success
}

func loopCheck(boxMap map[int]int, prisonerArray []int, allowedTries int, boxCount int) int {
	var success int
	for i := 0; i < len(prisonerArray); i++ {
		prisoner := prisonerArray[i]
		var openBoxes []int
		box := prisoner
		for i := 0; i < allowedTries; i++ {
			if contains(openBoxes, box) {
				continue
			} else if prisoner == boxMap[box] {
				openBoxes = append(openBoxes, box)
				success++
				continue
			}
			box = boxMap[box]
		}
	}
	return success
}
