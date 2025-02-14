package sprint;


import java.util.List;
import java.util.stream.Collectors;

public class StreamPolice {
    public List<Integer> processNumbers(List <Integer> numbers) {
        
        return numbers.stream()
        .filter(num -> num >= 0)
        .filter(num -> !(num % 5 == 0 && num % 10 != 0))
        .collect(Collectors.toList());
    }
}
