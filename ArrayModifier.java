package sprint;

import java.util.ArrayList;

public class ArrayModifier {
    public static ArrayList<Double> removeElementsBetween(ArrayList<Double> list, int index1, int index2) {
        
        index1 = Math.max(0, index1);
        index2 = Math.min(list.size(), index2);

        if (index2 < index1) {
            int temporary = index2;
            index2 = index1;
            index1 = temporary;
        }

        for (int i = index1; i < index2; i++ ) {
            list.remove(index1);
            }
            return list;
        }
    }