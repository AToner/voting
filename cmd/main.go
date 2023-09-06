package main

/*
 * Process a ballot, and return all candidates sorted in descending order by their total number
 * of points.  Each ballot can have up to 3 candidates; candidates get points based on their
 * place in the ballot: 1st position = 3 points; 2nd = 2 points; 3rd = 1 point
 */
/*List<String> getResults(List<String> ballot)
```
First Call:   ["A", "B", "C"] -> ["A", "B", "C"]
Second Call:  ["A", "D", "B"] -> ["A", "B", "D", "C"]

1) First candidiate to the largest number of votes
First Call:   ["A", "B", "C"] -> ["A", "B", "C"]
Second Call:  ["D", "E", "F"] -> ["A", "D", "B", "E", "C", "F"]

//2) First candidate with most first-place votes wins

*/

import (
	"fmt"
	"reflect"
	"sort"
	"time"
)

var (
	state candidates
)

func main() {
	tests := []struct {
		input  []string
		output []string
	}{
		{[]string{}, []string{}},
		{[]string{"A", "B", "C"}, []string{"A", "B", "C"}},
		{[]string{"A", "D", "B"}, []string{"A", "B", "D", "C"}},
		{[]string{}, []string{}},
		{[]string{"A", "B", "C"}, []string{"A", "B", "C"}},
		{[]string{"D", "E", "F"}, []string{"A", "D", "B", "E", "C", "F"}},
	}

	for _, test := range tests {
		if len(test.input) == 0 {
			state = make(candidates, 0)
			continue
		}

		result := getResults(test.input)

		if !reflect.DeepEqual(test.output, result) {
			fmt.Printf("Error! Input %v Expected %v, got %v\n", test.input, test.output, result)
		}
	}
}

type candidate struct {
	candidate    string
	votes        int
	lastVoteTime int64
}

type candidates []candidate

func (c candidates) vote(vote string, count int) bool {
	for _, aCandidate := range c {
		if aCandidate.candidate == vote {
			aCandidate.votes += count
			aCandidate.lastVoteTime = time.Now().UnixNano()
			return true
		}
	}
	return false
}

func (c candidates) Len() int {
	return len(c)
}

func (c candidates) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c candidates) Less(i, j int) bool {
	if c[i].votes == c[j].votes {
		return c[i].lastVoteTime > c[j].lastVoteTime
	}

	return c[i].votes > c[j].votes
}

func getResults(ballot []string) []string {
	for voteNumber, vote := range ballot {
		voteWeight := 3 - voteNumber
		if !state.vote(vote, voteWeight) {
			state = append(state, candidate{
				candidate: vote,
				votes:     voteWeight,
			})
		}
	}

	sort.Sort(state)

	result := make([]string, len(state))
	for i, k := range state {
		result[i] = k.candidate
	}
	return result
}
