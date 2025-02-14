package sprint;

public class AreaCalculator {

    public static double calculateArea(double side) {
        double area = side * side;
        return Math.round(area * 100.0) / 100.0;  
    }

    public static double calculateArea(double length, double width) {
        double area = length * width;
        return Math.round(area * 100.0) / 100.0;  
    }

    public static double calculateArea(double radius, boolean calculate) {
        if (!calculate) {
            return Double.NaN;  
        }
        double area = Math.PI * radius * radius;
        return Math.round(area * 100.0) / 100.0;  
    }
}
