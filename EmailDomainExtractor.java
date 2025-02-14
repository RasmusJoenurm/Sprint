package sprint;

import java.util.List;
import java.util.stream.Collectors;

public class EmailDomainExtractor {
    public List<String> extractDomains(List<String> emails) {
        
        return emails.stream()

            .filter(email -> email.contains("@")
                    && email.indexOf('@') == email.lastIndexOf('@')  // Exactly one '@'
                    && email.indexOf('@') > 0                        // Something before '@'
                    && email.indexOf('@') < email.length() - 1)      // Something after '@'

            .map(email -> email.split("@")[1].toLowerCase())

            .distinct()

            .collect(Collectors.toList());
    }
}
