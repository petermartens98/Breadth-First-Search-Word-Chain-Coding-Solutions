# Breadth First Search Solution in Python

''' 
Description:
Breadth-first search (BFS) is a graph traversal algorithm that explores all the vertices
of a graph in breadth-first order, i.e., it visits all the vertices at a given distance 
from the starting vertex before moving on to the vertices at the next distance. It starts
at the specified starting vertex and explores all the vertices that can be reached from that
vertex, then it moves on to explore all the vertices that can be reached from those vertices, 
and so on.

BFS is often used to find the shortest path between two nodes in an unweighted graph, as it
guarantees that it will find the shortest path first. It can also be used to detect cycles 
in a graph and to determine if a graph is connected. The algorithm uses a queue data structure 
to keep track of the nodes to be explored and marks the nodes as visited as they are explored 
to avoid processing them again.

'''

# Worst-Case Time Complexity: O(mk), where k is the average number of neighbors for each word.

# Necessary Imports
from collections import deque
import time

# Load dictionary file into a set for faster lookup
with open('dictionary.txt') as f:
    dictionary = set([word.strip().lower() for word in f])

# Function to generate neighbors of a given word
# Worst-Case Time Complexity: O(26n^2) or O(n^2) 
def get_neighbors(word):
    neighbors = []                                                      # Create list for neighbor words
    for i in range(len(word)):                                          # Loop through word
        for c in 'abcdefghijklmnopqrstuvwxyz':                          # Loop through alphabet
            if c != word[i]:                                            # Check its not the same character as original
                neighbor = word[:i] + c + word[i+1:]                    # Create new word replacing word[i] with c
                if neighbor in dictionary:                              # Check if new word is in dictionary
                    neighbors.append(neighbor)                          # Append to neighbors list
    return neighbors                                                    # Return neighbors list


# Function to find the shortest word chain between two words using Binary Search
# Worst-Case Time Complexity: O(mn^2) 
def find_word_chain(start_word, end_word):
    queue = deque([(start_word, [start_word])])                         # Initialize a queue with the start word
    visited = set([start_word])                                         # Set to keep track of visited words
    while queue:                                                        # While queue is not empty
        current_word, current_chain = queue.popleft()                   # Assign current word/chain by pop left most element
        if current_word == end_word:                                    # Check if end word is reached
            return current_chain                                        # Return current chain if end word reached
        for neighbor in get_neighbors(current_word):                    # Generate neighbors of current word and add them to queue
            if neighbor not in visited:                                 # Check cases neighbor is not visited
                visited.add(neighbor)                                   # Add neighbor to visted
                queue.append((neighbor, current_chain + [neighbor]))    # Append neighbor to the queue
    return None                                                         # If no chain is found return none

# Function to Test Word Chain
def test_word_chain(start_word, end_word, expected_output):
    chain = find_word_chain(start_word, end_word)
    if chain:
        if chain == expected_output:
            print("Test case passed!")
            print(f"Word chain from {start_word} to {end_word}: {' -> '.join(chain)}")
        else:
            print("Test case failed.")
    else:
        print(f"No word chain found from {start_word} to {end_word}")

# Example usage
def main(): 

    # Define test cases
    test_cases = [
        {
            "start_word": "cat",
            "end_word": "dog",
            "expected_output": ["cat", "cot", "dot", "dog"]
        },
        {
            "start_word": "cat",
            "end_word": "pig",
            "expected_output": None
        }
    ]

    # Test word chain for each test case
    for i, test_case in enumerate(test_cases):
        start_time = time.time()
        start_word = test_case["start_word"]
        end_word = test_case["end_word"]
        expected_output = test_case["expected_output"]
        
        chain = find_word_chain(start_word, end_word)
        elapsed_time = (time.time() - start_time) * 1000

        print(f"Test case {i+1}: start={start_word} end={end_word} elapsed={elapsed_time:.3f}ms")

        if chain is not None:
            if chain == expected_output:
                print("Test case passed!")
                print(f"Word chain from {start_word} to {end_word}: {' -> '.join(chain)}")
            else:
                print("Test case failed.")
        else:
            print(f"No word chain found from {start_word} to {end_word}")

        print("-------------------------------------------------------")

main()
