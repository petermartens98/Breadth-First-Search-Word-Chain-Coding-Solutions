package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Load dictionary file into a map for faster lookup
func loadDictionary(filename string) map[string]bool {
	dictionary := make(map[string]bool)
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(strings.ToLower(scanner.Text()))
		dictionary[word] = true
	}
	return dictionary
}

// Function to generate neighbors of a given word
// Worst-Case Time Complexity: O(26n^2) or O(n^2)
func getNeighbors(word string, dictionary map[string]bool) []string {
	var neighbors []string           							// Create list for neighbor words
	for i := 0; i < len(word); i++ { 							// Loop through word
		for c := 'a'; c <= 'z'; c++ { 							// Loop through alphabet
			if rune(word[i]) != c { 						// Check its not the same character as original
				neighbor := word[:i] + string(c) + word[i+1:] 			// Create new word replacing word[i] with c
				if dictionary[neighbor] {                     			// Check if new word is in dictionary
					neighbors = append(neighbors, neighbor) 		// Append to neighbors list
				}
			}
		}
	}
	return neighbors // Return neighbors list
}

// Function to find the shortest word chain between two words using Binary Search
// Worst-Case Time Complexity: O(mn^2)
func findWordChain(startWord string, endWord string, dictionary map[string]bool) []string {
	queue := [][]string{{startWord, startWord}} 						// Initialize a queue with the start word
	visited := map[string]bool{startWord: true} 						// Set to keep track of visited words
	for len(queue) > 0 {                        					        // While queue is not empty
		currentWord, currentChain := queue[0][0], queue[0][1] 				// Assign current word/chain by dequeuing left most element
		queue = queue[1:]                                     				// Remove the dequeued element from queue
		if currentWord == endWord {                           				// Check if end word is reached
			return strings.Split(currentChain, " ") 				// Return current chain if end word reached
		}
		for _, neighbor := range getNeighbors(currentWord, dictionary) { 		// Generate neighbors of current word and add them to queue
			if _, ok := visited[neighbor]; !ok { 					// Check cases neighbor is not visited
				visited[neighbor] = true                                        // Add neighbor to visted
				queue = append(queue, []string{neighbor, currentChain + " " + neighbor}) // Append neighbor to the queue
			}
		}
	}
	return nil 	// If no chain is found return nil
}

// Function to test word chain
func testWordChain(startWord string, endWord string, expectedOutput []string, dictionary map[string]bool) {
	chain := findWordChain(startWord, endWord, dictionary)
	if chain != nil {
		if strings.Join(chain, ", ") == strings.Join(expectedOutput, ", ") {
			fmt.Println("Test case passed!")
			fmt.Printf("Word chain from %s to %s: %s\n", startWord, endWord, strings.Join(chain, " -> "))
		} else {
			fmt.Println("Test case failed.")
		}
	} else {
		fmt.Printf("No word chain found from %s to %s\n", startWord, endWord)
	}
}

// Main Function
func main() {
	// Load dictionary file into a map for faster lookup
	dictFile, err := os.Open("dictionary.txt")
	if err != nil {
		fmt.Println("Error opening dictionary file:", err)
		return
	}
	defer dictFile.Close()

	dictionary := make(map[string]bool)
	scanner := bufio.NewScanner(dictFile)
	for scanner.Scan() {
		dictionary[strings.ToLower(scanner.Text())] = true
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading dictionary file:", err)
		return
	}

	// Define test cases
	testCases := []struct {
		startWord      string
		endWord        string
		expectedOutput []string
	}{
		{"cat", "dog", []string{"cat", "cot", "dot", "dog"}},
		{"cat", "pig", nil},
	}

	// Test word chain for each test case
	for i, testCase := range testCases {
		start := time.Now()
		chain := findWordChain(testCase.startWord, testCase.endWord, dictionary)
		elapsed := time.Since(start)

		fmt.Printf("Test case %d: start=%s end=%s elapsed=%v\n", i+1, testCase.startWord, testCase.endWord, elapsed)

		if chain != nil {
			if fmt.Sprintf("%v", chain) == fmt.Sprintf("%v", testCase.expectedOutput) {
				fmt.Println("Test case passed!")
				fmt.Printf("Word chain from %s to %s: %s\n", testCase.startWord, testCase.endWord, strings.Join(chain, " -> "))
			} else {
				fmt.Println("Test case failed.")
			}
		} else {
			fmt.Printf("No word chain found from %s to %s\n", testCase.startWord, testCase.endWord)
		}
		fmt.Println("-------------------------------------------------------")
	}
}
