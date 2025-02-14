package sprint;

public class BasicCalc {
    public int doOperation(int a, char op, int b) {
       if (op != '+' && op != '-' && op != '%' && op != '*' && op != '/' ) {
        return 0;
       }

        if (op == '+') {
            return a + b;
        } else if (op == '-') {
            return a - b;
        } else if (op == '/') {
            if (b == 0) {
                return 0;
            }
            return a / b;
        } else if (op == '*') {
            return a * b;
        } else if (op == '%') {
            if (b == 0) {
                return 0;
            }
            return a % b;
        } else {
            return 0;
        }
    }
}
