package sprint;

public class Factorial {
    public static Integer calculateFactorial(int n) {
        if (n == 0 || n == 1) {
            return 1;
        } else if (n > 0) {
            return n * calculateFactorial(n - 1);
        } else {
            return 0;
        }
    }
}
