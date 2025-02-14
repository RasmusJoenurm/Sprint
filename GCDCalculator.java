package sprint;

public class GCDCalculator {

    public int gcd(int a, int b) {
        a = Math.abs(a);
        b = Math.abs(b);
        
        if (a == 0) {
            return b;
        }
        if (b == 0) {
            return a;
        }
        while (a != b) {
            if (a > b) {
                a = a - b;
            } else {
                b = b - a;
            }
        }
        return a;
    }
}