package sprint;


public class SmallestDivisor {
    public int smallestDivisor(int number) {
        if (number <= 1) {
            return number;
        }
        int SquareRoot = (int) java.lang.Math.sqrt(number);
        for (int i = 2; i <= SquareRoot;i++) {
            if (number % i == 0) {
                return i;
            }
        }
        return number;
    }
}