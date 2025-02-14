package sprint;

public class CharCounter {
    public int countOccurrences(String input, char target) {
       String LowerCaseString = input.toLowerCase();
       char LowerCaseCharacter = Character.toLowerCase(target);

       int count = 0;
       for (int i = 0; i < LowerCaseString.length(); i++) {
        if (LowerCaseString.charAt(i) == LowerCaseCharacter) {
            count = count + 1;
        }
       }
       return count;
    }

}