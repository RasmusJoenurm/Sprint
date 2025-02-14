package sprint;

public class GCDRecursive {
    public Integer gcd(int a, int b) {
        if (a == 0 && b == 0) {
            return 0;
        } else if (a == 0) {
            return Math.abs(b);
        } else if (b == 0) {
            return Math.abs(a);
        } else {
            return gcd(b, a % b);
        }
    }
}
