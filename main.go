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

JAVA SPECIFIC: Scanner function may show some error/warning in vscode, ignore it. Do NOT use space/tab while adding
new keywords in Theory.Ques-and_keys function they might add inaccuracies/errors.  

GENERAL:
WRITE ANY KEYWORDS YOU WANT TO ADD IN PROGRAM IN LOWER CASE OTHERWISE FUNCITON WILL BREAK

Answer given by user is passed into format_ans  to remove any duplicate keywords which may result in inaccuracies.
Evaulate_ans contains questions and keywords in form of a map(we can replace this with a database later)

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

			f_ans := theory.Format_ans(ans)
			score := theory.Evaluate_ans(f_ans, keys)
			fmt.Println(score)
		}
		case 2:
			var opt int
			for q_num := 1; q_num <= 2; q_num++{
			ques,corr_ans:=mcq.Ques_keys(q_num)
			fmt.Println(ques)
			fmt.Scanln(&opt)
			if mcq.Evaluate(opt,corr_ans){
				fmt.Println("Correct answer!")
			}else{
				fmt.Println("Wrong\nAns:", corr_ans)
			}
		}

	}
}