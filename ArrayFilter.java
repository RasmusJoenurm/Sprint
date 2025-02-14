package sprint;

import java.util.ArrayList;
import java.util.List;

public class ArrayFilter {
    public int[][] filterBySum(int[][] array, int n) {
        List<int[]> filteredList = new ArrayList<>();
        
        for (int[] subArray : array) {
            int sum = 0;

            for (int num : subArray)
            sum += num;
         if (sum >= n) {
            filteredList.add(subArray);
         }
         }
         return filteredList.toArray(new int[filteredList.size()][]);
    }
}
