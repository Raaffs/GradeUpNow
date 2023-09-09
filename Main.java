


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
                    String[] result = Theory.quesAndKeys(qNum);
                    String ques = result[0];
                    String keysString = result[1];
                    String[] keys = keysString.split(",");

                    System.out.println(ques);
                    Scanner scan = new Scanner(System.in);
                    System.out.println("Enter your answer:");
                    String ans = scan.nextLine();

                    String filteredAns = Theory.removeDuplicate(ans);
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



