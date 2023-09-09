import java.util.Scanner;
public class MCQ {
    public static void AskMCQ(int quesNo, String ques, String opt1, String opt2, String opt3, String opt4, String correctAns) {
        Scanner scanner = new Scanner(System.in);

        System.out.println(quesNo + ". " + ques + "?");
        System.out.println("A. " + opt1);
        System.out.println("B. " + opt2);
        System.out.println("C. " + opt3);
        System.out.println("D. " + opt4);
        System.out.print("Enter Answer [A/B/C/D]: ");

        String ans = scanner.nextLine();

        if (ans.equals(correctAns)) {
            System.out.println("Correct Answer!");
        } else {
            System.out.println("Incorrect Answer, the correct answer is option " + correctAns);
        }
    }
}