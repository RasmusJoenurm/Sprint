package sprint;

public class RecursivePalindrome {
    // Main method to initialize the recursion
    public boolean isPalindrome(String str) {
        if (str == null) return false; // Handle null input
        str = str.replaceAll("[^a-zA-Z0-9]", "").toLowerCase(); // Clean the string
        return isPalindromeHelper(str, 0, str.length() - 1); // Call helper method
    }

    // Helper method to perform the recursive palindrome check
    private boolean isPalindromeHelper(String str, int start, int end) {
        // Base case: when the start index is greater than or equal to the end index
        if (start >= end) return true;

        // Ensure we skip any non-alphanumeric characters
        if (!Character.isLetterOrDigit(str.charAt(start))) {
            return isPalindromeHelper(str, start + 1, end);
        }

        if (!Character.isLetterOrDigit(str.charAt(end))) {
            return isPalindromeHelper(str, start, end - 1);
        }

        // If characters at start and end are the same, move to the next character pair
        if (str.charAt(start) == str.charAt(end)) {
            return isPalindromeHelper(str, start + 1, end - 1);
        }

        // If characters don't match, it's not a palindrome
        return false;
    }
}

