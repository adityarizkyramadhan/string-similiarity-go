package similiarity

import "strings"

// CompareTwoStrings compares two strings using bigrams and returns a similarity score.
func CompareTwoStrings(first, second string) float64 {
	first = strings.ReplaceAll(first, " ", "")
	second = strings.ReplaceAll(second, " ", "")

	if first == second {
		return 1.0
	}
	if len(first) < 2 || len(second) < 2 {
		return 0.0
	}

	firstBigrams := make(map[string]int)
	for i := 0; i < len(first)-1; i++ {
		bigram := first[i : i+2]
		firstBigrams[bigram]++
	}

	intersectionSize := 0
	for i := 0; i < len(second)-1; i++ {
		bigram := second[i : i+2]
		if count, exists := firstBigrams[bigram]; exists && count > 0 {
			firstBigrams[bigram]-- // decrement the count
			intersectionSize++
		}
	}

	return (2.0 * float64(intersectionSize)) / (float64(len(first) + len(second) - 2))
}

// findBestMatch finds the best match for a main string from a list of target strings.
func FindBestMatch(mainString string, targetStrings []string) (ratings []map[string]interface{}, bestMatch map[string]interface{}) {
	if !areArgsValid(mainString, targetStrings) {
		panic("Bad arguments: First argument should be a string, second should be an array of strings")
	}

	bestMatchIndex := 0
	for i, currentTargetString := range targetStrings {
		currentRating := CompareTwoStrings(mainString, currentTargetString)
		rating := map[string]interface{}{
			"target": currentTargetString,
			"rating": currentRating,
		}
		ratings = append(ratings, rating)

		if currentRating > ratings[bestMatchIndex]["rating"].(float64) {
			bestMatchIndex = i
		}
	}

	bestMatch = ratings[bestMatchIndex]

	return ratings, bestMatch
}

// areArgsValid checks if the arguments are valid.
func areArgsValid(mainString string, targetStrings []string) bool {
	if len(mainString) == 0 {
		return false
	}
	if len(targetStrings) == 0 {
		return false
	}
	for _, s := range targetStrings {
		if len(s) == 0 {
			return false
		}
	}
	return true
}
