package main

// Data Structures and Algorithms tests
//func main() {
//	// Data Structures ___________________________________________________________________
//	// Linked List Tests -----------------------------------------------------------------
//	fmt.Println("LINKED LIST TESTS:")
//	ll := LinkedList{}
//	ll.Insert(1)
//	ll.Insert(2)
//	ll.Insert(3)
//	ll.Display()
//	fmt.Println("Length: ", ll.Length())
//	ll.Delete(2)
//	ll.Display()
//	fmt.Printf("Linked List Length: %v\n\n", ll.Length())
//	// End Linked List Tests ---------------------------------------------------------
//
//	// Stack Tests ----------------------------------------------------------------------
//	fmt.Println("STACK TESTS:")
//	s := Stack{}
//	s.Push(1)
//	s.Push(2)
//	s.Push(3)
//	fmt.Println("Stack: ", s)
//	fmt.Println("Stack Length: ", len(s))
//	s.Pop()
//	fmt.Println("Stack: ", s)
//	fmt.Printf("Stack Length: %v\n\n", len(s))
//	// End Stack Tests ------------------------------------------------------------------
//
//	// Queue Tests ----------------------------------------------------------------------
//	fmt.Println("Queue Tests:")
//	q := Queue{}
//	q.Enqueue(1)
//	q.Enqueue(2)
//	q.Enqueue(3)
//	fmt.Println("Queue: ", q)
//	fmt.Println("Queue Length: ", len(q))
//	q.Dequeue()
//	fmt.Println("Queue: ", q)
//	fmt.Printf("Queue Length: %v\n\n", len(q))
//	// End Queue Tests ------------------------------------------------------------------
//
//	// Binary Tree Tests ----------------------------------------------------------------
//	root := &TreeNode{10, nil, nil}
//	root.Insert(5)
//	root.Insert(15)
//	root.Insert(3)
//	root.Insert(20)
//	root.Display()
//	fmt.Println("Binary Tree Length: ", root.Length())
//	root.Delete(15)
//	root.Delete(10)
//	root.Display()
//	fmt.Printf("Binary Tree Length: %v\n\n", root.Length())
//	// End Binary Tree Tests ------------------------------------------------------------
//
//	// Hash Table Tests -----------------------------------------------------------------
//	h := HashTable{}
//	h["one"] = 1
//	h["two"] = 2
//	h["three"] = 3
//	fmt.Println("Hash Table: ", h)
//	fmt.Printf("Hash Table Length: %v\n", len(h))
//	delete(h, "two")
//	fmt.Println("Hash Table: ", h)
//	fmt.Printf("Hash Table Length: %v\n\n", len(h))
//	// End Hash Table Tests -------------------------------------------------------------
//	// End Data Structures ______________________________________________________________
//
//	// Sorting Algorithms _______________________________________________________________
//	// Bubble Sort ----------------------------------------------------------------------
//	fmt.Println("Bubble Sort:")
//	bubbleArr := []int{64, 34, 25, 12, 22, 11, 90}
//	fmt.Println("Original Array: ", bubbleArr)
//	BubbleSort(bubbleArr)
//	fmt.Printf("Bubble Sorted array: %v\n\n", bubbleArr)
//
//	// Selection Sort -------------------------------------------------------------------
//	fmt.Println("Selection Sort:")
//	selectionArr := []int{64, 34, 25, 12, 22, 11, 90}
//	fmt.Println("Original Array: ", selectionArr)
//	SelectionSort(selectionArr)
//	fmt.Printf("Selection Sorted array: %v\n\n", selectionArr)
//
//	// Insertion Sort -------------------------------------------------------------------
//	fmt.Println("Insertion Sort:")
//	insertionArr := []int{64, 34, 25, 12, 22, 11, 90}
//	fmt.Println("Original Array: ", insertionArr)
//	InsertionSort(insertionArr)
//	fmt.Printf("Insertion Sorted array: %v\n\n", insertionArr)
//
//	// Merge Sort ----------------------------------------------------------------------
//	fmt.Println("Merge Sort:")
//	mergeArr := []int{64, 34, 25, 12, 22, 11, 90}
//	fmt.Println("Original Array: ", mergeArr)
//	MergeSort(mergeArr, 0, len(mergeArr)-1)
//	fmt.Printf("Merge Sorted array: %v\n\n", mergeArr)
//
//	// Quick Sort ----------------------------------------------------------------------
//	fmt.Println("Quick Sort:")
//	quickArr := []int{64, 34, 25, 12, 22, 11, 90}
//	fmt.Println("Original Array: ", quickArr)
//	QuickSort(quickArr, 0, len(quickArr)-1)
//	fmt.Printf("Quick Sorted array: %v\n\n", quickArr)
//
//	// Search Algorithms _______________________________________________________________
//	sortedArr := []int{11, 12, 22, 25, 34, 64, 90}
//	target := 25
//	// Binary Search -------------------------------------------------------------------
//	fmt.Println("Binary Search:")
//	fmt.Println("Array: ", sortedArr)
//	result := BinarySearch(sortedArr, target)
//	found(target, result)
//
//	// Linear Search -------------------------------------------------------------------
//	fmt.Println("Linear Search:")
//	fmt.Println("Array: ", sortedArr)
//	result = LinearSearch(sortedArr, target)
//	found(target, result)
//
//	// Jump Search ---------------------------------------------------------------------
//	fmt.Println("Jump Search:")
//	fmt.Println("Array: ", sortedArr)
//	result = JumpSearch(sortedArr, target)
//	found(target, result)
//
//	// Interpolation Search ------------------------------------------------------------
//	fmt.Println("Interpolation Search:")
//	fmt.Println("Array: ", sortedArr)
//	result = InterpolationSearch(sortedArr, target)
//	found(target, result)
//	// End Algorithms __________________________________________________________________
//}

// Function for generic algorithm tests above
//func found(t, r int) {
//	if r != -1 {
//		fmt.Printf("%v found at index %d\n\n", t, r)
//	} else {
//		fmt.Printf("%v not found\n\n", t)
//	}
//}

// HashMap most freq scores and letter grade problem
//// The main function is the entry point of the program.
//func main() {
//	// Initialize a slice of integers representing scores.
//	scores := []int{98, 100, 93, 90, 86, 83, 77, 75, 76, 81, 85, 90, 93, 95, 97, 99, 100, 100}
//
//	// Call the function mostFrequentScore to determine the most frequent score and its frequency map.
//	mostFreqScore, scoreMap := mostFrequentScore(scores)
//
//	// Print the most frequent score.
//	fmt.Printf("Most frequent score: %d\n\n", mostFreqScore)
//
//	// Call the function mostFrequentLetterGrade to determine the most frequent letter grade based on the frequency map.
//	mostFreqLetter := mostFrequentLetterGrade(scoreMap)
//
//	// Print the most frequent letter grade.
//	fmt.Printf("Most frequent letter: %s\n\n", mostFreqLetter)
//}
//
//// Function to determine the most frequent score and return it along with a map of score frequencies.
//func mostFrequentScore(scores []int) (int, map[int]int) {
//	// Create a map to keep track of the frequency of each score.
//	scoreMap := make(map[int]int)
//
//	// Initialize variables to keep track of the most frequent score and its count.
//	maxCount := 0
//	mostFrequent := scores[0]
//
//	// Iterate over each score in the slice.
//	for _, score := range scores {
//		// Increment the count for the current score in the map.
//		scoreMap[score]++
//
//		// Update the most frequent score if the current score's count is greater than maxCount.
//		if scoreMap[score] > maxCount {
//			maxCount = scoreMap[score]
//			mostFrequent = score
//		}
//	}
//
//	// Return the most frequent score and the map of score frequencies.
//	return mostFrequent, scoreMap
//}
//
//// Function to determine the most frequent letter grade based on a map of score frequencies.
//func mostFrequentLetterGrade(scoreMap map[int]int) string {
//	// Return a message if no scores are available.
//	if len(scoreMap) == 0 {
//		return "No scores available"
//	}
//
//	// Initialize a map to keep track of the count of each letter grade.
//	letterScoreMap := map[string]int{"A": 0, "B": 0, "C": 0, "D": 0, "F": 0}
//
//	// Iterate over each score and its count in the scoreMap.
//	for score, count := range scoreMap {
//		// Determine the letter grade for the current score and update the letterScoreMap map.
//		switch {
//		case score >= 90:
//			letterScoreMap["A"] += count
//		case score >= 80 && score < 90:
//			letterScoreMap["B"] += count
//		case score >= 70 && score < 80:
//			letterScoreMap["C"] += count
//		case score >= 65 && score < 70:
//			letterScoreMap["D"] += count
//		case score >= 0 && score < 65:
//			letterScoreMap["F"] += count
//		default:
//			log.Panicf("Letter grade Switch defaulted. SCORE : COUNT pair that defulted: %d : %d\n", score, count)
//		}
//	}
//
//	// Initialize variables to keep track of the most frequent letter grade and its count.
//	var mostFrequentLetterGrade string
//	mostFrequentLetterGradeCount := 0
//
//	// Iterate over each letter grade and its count in the letterScoreMap map.
//	for grade, count := range letterScoreMap {
//		// Update the most frequent letter grade if the current grade's count is greater than mostFrequentLetterGradeCount.
//		if count > mostFrequentLetterGradeCount {
//			mostFrequentLetterGradeCount = count
//			mostFrequentLetterGrade = grade
//		}
//	}
//
//	// Return the most frequent letter grade.
//	return mostFrequentLetterGrade
//}
