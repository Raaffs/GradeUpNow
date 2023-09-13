package theory

import (
	"strings"
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

func Ques_keys(q_num int) (string, []string) {
	question := map[int]string{
		1: "What is a pointer?",
		2: "What are some concepts used in OOP?",
		3: "What is method overloading?",
	}

	keys := map[int][]string{
		1: {"memory", "address", "location", "reference"},
		2: {"classes", "objects", "abstraction", "encapsulation", "inheritance", "polymorphism", "access",
			"specifiers", "public", "private", "protected", "interface"},
		3: {"parameters", "arguments", "data", "type", "ambiguity"},
	}

	return question[q_num], keys[q_num]
}

//Uses levenshtein distance algorithm to match keywords that are similar, in cases where user uses different for of same 
//word, or makes some spelling error while entering the answer.

func Evaluate_ans(ans string, key []string) int {
	var count int

	 for _, word := range strings.Fields(ans) {
		for i := 0; i < len(key); i++ {
			correct_keyword := key[i]
			distance := levenshtein.DistanceForStrings([]rune(word), []rune(correct_keyword), levenshtein.DefaultOptions)

			threshold := 0.5 * float64(len(correct_keyword))

			if float64(distance) <= threshold {
				count++
				// Removes matched keyword from the list
				key = append(key[:i], key[i+1:]...)
				break
			}
		}
	}

	return count
}
