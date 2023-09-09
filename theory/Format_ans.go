package theory
import(
	"strings"
)

func Remove_duplicate(ans string) string {

	lower_ans:=strings.ToLower(ans)
    // Split the ans string into words
    words:=strings.Fields(lower_ans)

	
    
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
