package electionday

import "fmt"

// NewVoteCounter returns a new vote counter with
// a given number of inital votes.
func NewVoteCounter(initialVotes int) *int {
	//var pointerToVotes *int
	//pointerToVotes = &initialVotes
	//return pointerToVotes
	return &initialVotes

}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}
	return *counter // dereferencing
}

// IncrementVoteCount increments the value in a vote counter
func IncrementVoteCount(counter *int, increment int) {
	numberInCounter := *counter
	fmt.Println("counter as it comes", counter)
	newVotesInCounter := numberInCounter + increment
	fmt.Println("new votes in counter", &newVotesInCounter)
	*counter = newVotesInCounter
	fmt.Println("counter at the end", counter)
}

// NewElectionResult creates a new election result
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	return &ElectionResult{
		Name:  candidateName,
		Votes: votes,
	}
}

// DisplayResult creates a message with the result to be displayed
func DisplayResult(result *ElectionResult) string {
	// dereferencing is explicit in structs, see below the same result
	return fmt.Sprintf("%s (%d)", (*result).Name, result.Votes)
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	//value, exists := results[candidate]
	//if !exists {
	//	return
	//}
	results[candidate] = results[candidate] - 1

}
