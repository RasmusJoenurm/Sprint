package sprint;

public class DigitSum {

    public static int sumOfDigits(int number) {
        if (number < 0) {
            number = (number * (-1));
        }
        var sum = 0;
        for (; number > 0; number /= 10) {
            sum += number % 10;
        }
        return sum;
    }

}