package main
import (

)
func mcq_Ques_keys(q_num int)(string,int){
    ques:=map[int]string{
        1: "What is the function of a pointer?\n1.Store data\n2.Store memory address of another variable\n3.Reference to variable\n4.A special function.",
        2: "Which of the following is a non linear data structure?\n1.Tree\n2.Array\n3.Linked list\n4.Queue",
    }
    keys:=map[int]int{
        1:2,
        2:1,
    }
    return ques[q_num],keys[q_num]
}
func Evaluate(opt int, corr_ans int) bool{
    return opt==corr_ans
}
