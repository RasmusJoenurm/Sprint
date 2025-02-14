package sprint;



import java.util.Arrays;

public class AnagramChecker {
    public static boolean areAnagrams(String input1, String input2) {
        String normalString1 = input1.toLowerCase();
        String normalString2 = input2.toLowerCase();

        if (normalString1.length() != normalString2.length()) {
            return false;
        }
        
        char[] charArray1 = normalString1.toCharArray();
        char[] charArray2 = normalString2.toCharArray();

        Arrays.sort(charArray1);
        Arrays.sort(charArray2);

        return Arrays.equals(charArray1, charArray2);
    }
}
