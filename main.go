package main
/*
Brief functionality overview:
GOLANG SPECIFIC: Function/varibales in different packages(e.g in theory folder) must start with capital letter
if you want to access them in other package(e.g main). And access them with package_name.FunctionName().

bufio/os is used to read answer from user as fmt.Scan can only read one word at a time.

RUNNING PROGAM IN GOLANG:
go mod init github.com/Suy56/GradeUpNow
go build
./GradeUpNow

JAVA SPECIFIC:

GENERAL:
WRITE ANY KEYWORDS YOU WANT TO ADD IN PROGRAM IN LOWER CASE OTHERWISE FUNCITON WILL BREAK

Answer given by user is passed into format_ans  to remove any duplicate keywords which may result in inaccuracies.
Evaulate_ans contains questions and keywords in form of a map(we can replace this with a database later)

First read main() function to get general understanding of how code works, after that goto format_ans and finally to Evaluate_ans
*/

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Suy56/GradeUpNow/mcq"
	"github.com/Suy56/GradeUpNow/theory"
)

func main() {
	var choice int
	fmt.Println("1.Practice theory questions\n2.Practice MCQs")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		for q_num := 1; q_num < 4; q_num++ {
			ques, keys := theory.Ques_keys(q_num)
			//Ques_keys has two return value one return question and other returns keywords related to that question.
			// If we only need one of it, use '_' for the value we don't need. Complier will ignore that return value.
			fmt.Println(ques)
			scan := bufio.NewScanner(os.Stdin)
			fmt.Println("Enter your answer:")
			scan.Scan()
			ans := scan.Text()

			f_ans := theory.Remove_duplicate(ans)
			mat_keys := theory.Evaluate_ans(f_ans, keys)
			fmt.Println(mat_keys)
		}
	case 2:
		mcq.AskMCQ(1, "How many days does the earth take to revolve around the sun", "360", "300", "365.25", "365", "C")
	}

}