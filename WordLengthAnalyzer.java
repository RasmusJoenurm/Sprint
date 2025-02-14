package sprint;

import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

public class WordLengthAnalyzer {

    public Map<Integer, Integer> analyzeWordLengths(List<String> list) {
        return list.stream()
                   .collect(Collectors.groupingBy(
                       word -> word.length(),  
                       Collectors.collectingAndThen(
                           Collectors.counting(),  
                           Long::intValue        
                       )
                   ));
    }
}
