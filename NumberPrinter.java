package sprint;

public class NumberPrinter {
    public void printNums(int number) {
        if (number < 0) {
            System.out.println("Negative numbers are not allowed");
        } 
        for (int i = 0; i <= number; i++) {
            System.out.println(i);
        }
    }
}
