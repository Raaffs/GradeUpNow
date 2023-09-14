package main
/*
Brief functionality overview:
GOLANG SPECIFIC: Function/varibales in different packages(e.g in theory folder) must start with capital letter
if you want to access them in other package(e.g main). And access them with package_name.FunctionName().

bufio/os is used to read answer from user as fmt.Scan can only read one word at a time.

Setup: 
go get github.com/texttheater/golang-levenshtein/levenshtein	//for pattern matching
go get github.com/gorilla/mux									//to handle http request
go mod init github.com/Suy56/GradeUpNow
go build
./GradeUpNow
  

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
	choice:=0	
	for choice!=3{
		fmt.Println("1.Practice theory questions\n2.Practice MCQs\n3.Exit")
		fmt.Scanln(&choice)
		switch choice {
			case 1:
			for q_num := 1; q_num <= 2; q_num++ {
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
			fmt.Println("You scored",score,"points on this question\n")
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
		if(choice==3){
			break;
		}
	
	
	}

	}
}