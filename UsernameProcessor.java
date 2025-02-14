package sprint;

import java.util.List;

public class UsernameProcessor {
    public String findFirstUsername(List<String> list) {
        return list.stream()
        .findFirst()
        .orElse("Anonymous");
    }
}
