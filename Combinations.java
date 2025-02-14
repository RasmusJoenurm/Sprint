package sprint;

import java.util.ArrayList;
import java.util.List;

public class Combinations {
   
    public List<String> combN(int n) {
        List<String> validComb = new ArrayList<>();

        if (n <= 0) {
            return validComb;
        }

        for(int i = 0; i <= 9; i++) {
            generateCombinations(validComb, String.valueOf(i), i + 1, n - 1);
        }
        
        return validComb;
    }

    private void generateCombinations(List<String> validComb, String currentString, int start, int remainingDigits) {
        if (remainingDigits == 0) {
            validComb.add(currentString);
            return;
        }
        for (int i = start; i <= 9; i++) {
            generateCombinations(validComb, currentString + i, i + 1, remainingDigits - 1);
        }
    }
}
