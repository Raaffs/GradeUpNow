package main
import (
    "strings"
	"regexp"
    "github.com/texttheater/golang-levenshtein/levenshtein"
)



func Format_ans(ans string) string {
    // Convert the input to lowercase
    lowerAns := strings.ToLower(ans)

    // Use a regular expression to match and replace symbols with empty strings
    re := regexp.MustCompile(`[^a-zA-Z0-9\s]`)
    cleanAns := re.ReplaceAllString(lowerAns, "")

    // Split the cleaned string into words
    words := strings.Fields(cleanAns)

    // Create a map to store unique words
    uniqueWords := make(map[string]bool)

    for _, word := range words {
        uniqueWords[word] = true
    }

    // Convert unique words back into a string
    var result string
    for word := range uniqueWords {
        result += word + " "
    }

    // Trim any leading or trailing spaces
    result = strings.TrimSpace(result)

    return result
}
// Uses levenshtein distance algorithm to match keywords that are similar, in cases where the user uses a different form of the same
// word or makes some spelling error while entering the answer.

func Evaluate_ans(ans string, key string) float64 {
    var count int
    totalKeys := len(key)
    key_arr:=strings.Split(key,",")
    for _, word := range strings.Split(ans,",") {

        for i, correctKeyword := range key_arr {
            distance := levenshtein.DistanceForStrings([]rune(word), []rune(correctKeyword), levenshtein.DefaultOptions)
            threshold := 0.5 * float64(len(correctKeyword))
            if float64(distance) <= threshold {
                count++
                // Removes matched keyword from the list
                key_arr = append(key_arr[:i], key_arr[i+1:]...)
                break
            }
        }

        
    }

    // Calculate the percentage of matched words
    percent := (float64(count) / float64(totalKeys)) * 100
	if percent>75{
		return 10
	}else if percent>70{
		return 9
	}else if percent>60{
		return 8
	}else if percent>50{
		return 7
	}else if percent>40{
		return 6
	}else if percent>25{
		return 5
	}else{
		return 2
	}
}
