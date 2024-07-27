package main

import (
	"fmt"
	"log"
)

// HashMap most freq scores and letter grade problem
// The main function is the entry point of the program.
func main() {
	// Initialize a slice of integers representing scores.
	scores := []int{98, 100, 93, 90, 86, 83, 77, 75, 76, 81, 85, 90, 93, 95, 97, 99, 100, 100}

	// Call the function mostFrequentScore to determine the most frequent score and its frequency map.
	mostFreqScore, scoreMap := mostFrequentScore(scores)

	// Print the most frequent score.
	fmt.Printf("Most frequent score: %d\n\n", mostFreqScore)

	// Call the function mostFrequentLetterGrade to determine the most frequent letter grade based on the frequency map.
	mostFreqLetter := mostFrequentLetterGrade(scoreMap)

	// Print the most frequent letter grade.
	fmt.Printf("Most frequent letter: %s\n\n", mostFreqLetter)
}

// Function to determine the most frequent score and return it along with a map of score frequencies.
func mostFrequentScore(scores []int) (int, map[int]int) {
	// Create a map to keep track of the frequency of each score.
	scoreMap := make(map[int]int)

	// Initialize variables to keep track of the most frequent score and its count.
	maxCount := 0
	mostFrequent := scores[0]

	// Iterate over each score in the slice.
	for _, score := range scores {
		// Increment the count for the current score in the map.
		scoreMap[score]++

		// Update the most frequent score if the current score's count is greater than maxCount.
		if scoreMap[score] > maxCount {
			maxCount = scoreMap[score]
			mostFrequent = score
		}
	}

	// Return the most frequent score and the map of score frequencies.
	return mostFrequent, scoreMap
}

// Function to determine the most frequent letter grade based on a map of score frequencies.
func mostFrequentLetterGrade(scoreMap map[int]int) string {
	// Return a message if no scores are available.
	if len(scoreMap) == 0 {
		return "No scores available"
	}

	// Initialize a map to keep track of the count of each letter grade.
	letterScoreMap := map[string]int{"A": 0, "B": 0, "C": 0, "D": 0, "F": 0}

	// Iterate over each score and its count in the scoreMap.
	for score, count := range scoreMap {
		// Determine the letter grade for the current score and update the letterScoreMap map.
		switch {
		case score >= 90:
			letterScoreMap["A"] += count
		case score >= 80 && score < 90:
			letterScoreMap["B"] += count
		case score >= 70 && score < 80:
			letterScoreMap["C"] += count
		case score >= 65 && score < 70:
			letterScoreMap["D"] += count
		case score >= 0 && score < 65:
			letterScoreMap["F"] += count
		default:
			log.Panicf("Letter grade Switch defaulted. SCORE : COUNT pair that defulted: %d : %d\n", score, count)
		}
	}

	// Initialize variables to keep track of the most frequent letter grade and its count.
	var mostFrequentLetterGrade string
	mostFrequentLetterGradeCount := 0

	// Iterate over each letter grade and its count in the letterScoreMap map.
	for grade, count := range letterScoreMap {
		// Update the most frequent letter grade if the current grade's count is greater than mostFrequentLetterGradeCount.
		if count > mostFrequentLetterGradeCount {
			mostFrequentLetterGradeCount = count
			mostFrequentLetterGrade = grade
		}
	}

	// Return the most frequent letter grade.
	return mostFrequentLetterGrade
}
