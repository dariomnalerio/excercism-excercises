public class CarsAssemble {

    private static final int RATE = 221;

    private double getSuccessRate(int speed) {
        if (speed >= 1 && speed <= 4) {
            return 1.0;
        } else if (speed <= 8) {
            return 0.9;
        } else if (speed == 9) {
            return 0.8;
        } else {
            return 0.77;
        }
    }

    public double productionRatePerHour(int speed) {
        double successRate = getSuccessRate(speed);
        return RATE * speed * successRate;
    }

    public int workingItemsPerMinute(int speed) {
        return (int) (productionRatePerHour(speed) / 60);
    }
}
