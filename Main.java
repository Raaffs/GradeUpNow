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

import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int choice;
        System.out.println("1. Practice theory questions");
        System.out.println("2. Practice MCQs");
        choice = scanner.nextInt();

        switch (choice) {
            case 1:
                for (int qNum = 1; qNum < 4; qNum++) {
                    String[] result = Theory.Ques_and_keys(qNum);
                    String ques = result[0];
                    String keysString = result[1];
                    String[] keys = keysString.split(",");

                    System.out.println(ques);
                    Scanner scan = new Scanner(System.in);
                    System.out.println("Enter your answer:");
                    String ans = scan.nextLine();

                    String filteredAns = Theory.Remove_duplicate(ans);
                    int matchingKeys = Theory.Evaluate_ans(filteredAns, keys);
                    System.out.println(matchingKeys);
                }
                break;

            case 2:
                MCQ.AskMCQ(1, "How many days does the earth take to revolve around the sun", "360", "300", "365.25", "365", "C");
                break;
        }
    }
}



