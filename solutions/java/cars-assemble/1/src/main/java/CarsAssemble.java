public class CarsAssemble {
    
    private final int CARS_PER_HOUR = 221;

    public double productionRatePerHour(int speed) {
        double successRate =
        speed == 10 ? 0.77 :
        speed == 9  ? 0.8  :
        speed >= 5  ? 0.9  :
        speed >= 1  ? 1.0  : 0.0;

    return speed * CARS_PER_HOUR * successRate;
    }

    public int workingItemsPerMinute(int speed) {
        return (int)(productionRatePerHour(speed) / 60);
    }
}
