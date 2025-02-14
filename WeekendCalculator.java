package sprint;

import java.time.DayOfWeek;
import java.time.LocalDate;

public class WeekendCalculator {
    public long countWeekendDays(LocalDate start, LocalDate finish) {
        long weekendDayCount = 0;

        for (LocalDate date = start; !date.isAfter(finish); date = date.plusDays(1)) {
            DayOfWeek  dayOfWeek = date.getDayOfWeek();

            if (dayOfWeek == DayOfWeek.SATURDAY || dayOfWeek == DayOfWeek.SUNDAY) {
                weekendDayCount++;
            }
       }
       return weekendDayCount;
    }
}
