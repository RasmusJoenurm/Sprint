package sprint;

import java.util.List;
import java.util.stream.Collectors;

public class StringToIntConverter {
    public List<Integer> convertStringListToIntList(List<String> list) {
        List<Integer> result = list.stream()
    .map(Integer::valueOf)
    .collect(Collectors.toList());

    return result;
    }
}
