package theory
import(
	"strings"
)
func Ques_keys(q_num int)(string,[]string){
	question:=map[int]string{
		1:"What is a pointer?",
		2:"What are some concepts used in OOP?",
		3:"What is method overloading?",
	}
	//Will've change functionality in evaluate_ans so that if 85-90% characters in keywords matches it will evaluate to true
	//in case user enter different grammatical form of same sentence or makes some dumb spelling mistake. 
	keys:=map[int][]string{
		1:{"memory","address","location","reference","refrence","&","*",},
		2:{"classes","objects","class","oject","abstraction","encapsulation","inheritance","polymorphism","access",
			"specifiers","public","private","protected","interface","interfaces",},
		3:{"parameters", "number of arguements","arguments","data type","ambiguity"},
	}
	return question[q_num],keys[q_num]
}
func  Evaluate_ans(ans string, key []string)int{
	var count int

	for _,words:=range strings.Fields(ans){
		for _,correct_keywords:=range key {
			
			if words==correct_keywords{
				count++
			}
		}
	}
	return count
}