package mcq
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func AskMCQ(quesNo int, ques, opt1, opt2, opt3, opt4, correctAns string) {
    scanner := bufio.NewScanner(os.Stdin)

    println(quesNo, ".", ques+"?")
    println("A.", opt1)
    println("B.", opt2)
    println("C.", opt3)
    println("D.", opt4)
    print("Enter Answer [A/B/C/D]: ")

    scanner.Scan()
    ans := strings.ToUpper(scanner.Text())

    if ans == correctAns {
        fmt.Println("Correct Answer!")
    } else {
        fmt.Println("Incorrect Answer, the correct answer is option", correctAns)
    }
}
