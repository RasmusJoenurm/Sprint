package sprint;

public class PrimeChecker {
    public static boolean isPrime(int number) {
        if (number == 0 || number == 1 || number < 0) {
            return false;
        }
        for (int i = 2; i <= (number / 2);i++ ) {
            if (number % i == 0) {
                return false;
            }
        }
        return true;
    }
}