using System;

static class AssemblyLine
{
    public static double SuccessRate(int speed)
    {
        if (speed == 0)
        {
            return 0.0;
        }
        else if (speed < 5)
        {
            return 1.0;
        }
        else if (speed < 9)
        {
            return 0.9;
        }
        else if (speed == 9)
        {
            return 0.8;
        }
        else
        {
            return 0.77;
        }
    }

    public static double ProductionRatePerHour(int speed)
    {
        double carProduction = 221;
        double succesRate = SuccessRate(speed);
        double productionRatePerHour = speed * carProduction * succesRate;

        return productionRatePerHour;
    }

    public static int WorkingItemsPerMinute(int speed)
    {
        double productionRatePerHour = ProductionRatePerHour(speed);
        int productionRatePerMinute = (int)productionRatePerHour / 60;

        return productionRatePerMinute;
    }
}
