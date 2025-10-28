// Result - Paseed
// Runtime - Beats 79.93%
// Memory - Beats 95.36%

package main

import (
	"fmt"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	// Sort strings lexicographically and add to map
	anagramMap := make(map[string][]string)
	for _, str := range strs {
		runes := []rune(str)
		slices.Sort(runes)
		sortedRunes := string(runes)
		anagramMap[sortedRunes] = append(anagramMap[sortedRunes], str)
	}

	// Put map results in result list
	result := [][]string{}
	for _, v := range anagramMap {
		result = append(result, v)
	}

	return result
}

func main() {
	var strs []string

	strs = []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Printf("%v\n", groupAnagrams(strs))

	strs = []string{""}
	fmt.Printf("%v\n", groupAnagrams(strs))

	strs = []string{"a"}
	fmt.Printf("%v\n", groupAnagrams(strs))
}
