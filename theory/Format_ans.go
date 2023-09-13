package theory

import (
    "strings"
    "regexp"
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
