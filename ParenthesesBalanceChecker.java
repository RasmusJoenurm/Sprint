package sprint;


public class ParenthesesBalanceChecker {
    public boolean isBalanced(String str) {
        if (str == null) {
            return false;
        }
        return checkBalance(str, 0, 0);
    }
    private boolean checkBalance(String str, int index, int balance) {
        if (index == str.length()) {
            return balance == 0;
        }
        
        char currentChar = str.charAt(index);

        if (currentChar == '(') {
            balance++;
        } else if (currentChar == ')') {
            balance--;
        }
        if (balance < 0) {
            return false;
        }
        return checkBalance(str, index + 1, balance);
    }
}
