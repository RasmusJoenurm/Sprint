package sprint;


public class BetweenLimits {

  public String findRange(char from, char to) {
      if (from > to) {
        char temporary = from;
        from = to;
        to = temporary;
      }
      var result = new StringBuilder();
      for (char i = (char) (from + 1); i < to; i++) {
        result.append(i);
      }
      return result.toString();
    }
    
}