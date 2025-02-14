package sprint;

import java.util.ArrayList;
import java.util.List;

public class PrimeFinder {
    public static List<Integer> findPrimesUpTo(Integer limit) {
        List<Integer> prime = new ArrayList<>();

        for (int num = 2; num <= limit; num++) {
            if(isPrime(num)) {
                prime.add(num);
            }
       }
       return prime;
    }
    private static boolean isPrime(int num) {
        if (num <= 1) {
            return false;
        }
        for (int i = 2; i * i <= num; i++) {
            if (num % i == 0) {
                return false;
            }
        }
        return true;
    }
}
