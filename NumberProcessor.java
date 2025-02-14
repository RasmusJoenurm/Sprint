package sprint;

import java.util.List;
import java.util.Optional;

public class NumberProcessor {
    public Optional <Integer> processNumbers(List<Integer> list) {
        return list.stream()
        .filter(number -> number >= 10)
        .reduce((a, b) -> a * b);
    }
}
