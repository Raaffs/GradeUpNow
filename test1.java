import java.util.Scanner;

class MCQ {
    Scanner sc = new Scanner(System.in);

    MCQ(int quesno, String ques, String opt1, String opt2, String opt3, String opt4, String correctans) {
        System.out.println(quesno + ". " + ques + "?");
        System.out.println("A. " + opt1);
        System.out.println("B. " + opt2);
        System.out.println("C. " + opt3);
        System.out.println("D. " + opt4);
        System.out.print("Enter Answer [A/B/C/D]: ");
        String ans = sc.nextLine();
        if (ans.equals(correctans)) {
            System.out.println("Correct Answer!");
        } else {
            System.out.println("Incorrect Answer, the correct answer is option " + correctans);
        }
    }
}

public class test2 {
    public static void main(String[] args) {
        MCQ m1 = new MCQ(1, "How many days does the earth take to revolve around the sun",
                "360", "300", "365.25", "365", "C");
    }
}
