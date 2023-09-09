import java.util.HashSet;
import java.util.Set;
public class Theory {
    public static String[] Ques_and_keys(int qNum) {
        String[] questions = {
                "What is a pointer?",
                "What are some concepts used in OOP?",
                "Explain method overloading in java"
        };
    //Will've change functionality in evaluate_ans so that if 85-90% characters in keywords matches it will evaluate to true
	//in case user enter different grammatical form of same word or makes some dumb spelling mistake. 
        String[] keys = {
                "memory,address,location,reference,refrence,&,*",
                "classes,objects,class,oject,abstraction,encapsulation,inheritance,polymorphism,access,specifiers,public,private,protected,interface,interfaces",
                "multiple,parameters,arguments,ambiguity,data,type"
        };

        return new String[]{questions[qNum - 1], keys[qNum - 1]};
    }

    public static int Evaluate_ans(String ans, String[] key) {
        int count = 0;
        String[] words = ans.split("\\s+");

        for (String word : words) {
            for (String correctKeyword : key) {
                if (word.equals(correctKeyword)) {
                    count++;
                }
            }
        }
        return count;
    }

    public static String Remove_duplicate(String ans) {
        ans = ans.toLowerCase();
        String[] words = ans.split("\\s+");

        Set<String> uniqueWords = new HashSet<>();

        for (String word : words) {
            uniqueWords.add(word);
        }

        StringBuilder result = new StringBuilder();

        for (String word : uniqueWords) {
            result.append(word).append(" ");
        }

        return result.toString().trim();
    }
}
