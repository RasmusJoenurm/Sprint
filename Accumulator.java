package sprint;

public class Accumulator {

    public int accumulate(int n) {
        if (n <= 0) {
            return 0;
        }
        int result = 0;
        for (int i = 0; i <= n; i++) {
            result = result + i;
        }
        return result;
    }
    
}