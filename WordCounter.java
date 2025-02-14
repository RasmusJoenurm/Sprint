package sprint;


public class WordCounter {
    public int countWords(String sentence) {
        if (sentence.isEmpty()) {
            return 0;
        }
        int wordCount = 0;
        boolean inWord = false;
        for (int i = 0; i < sentence.length(); i++) {
            char currentCharacter = sentence.charAt(i);
            if (Character.isLetterOrDigit(currentCharacter)) {
                if (!inWord) {
                    wordCount++;
                    inWord = true;
                }
            } else {
                inWord = false;
            }
        }
        return wordCount;
    }
}