using System;
class Lasagna
{
    private int ExpectedMinutes = 40;
    public int ExpectedMinutesInOven()
    {
        return ExpectedMinutes;
    }
    public int RemainingMinutesInOven(int elapsedMinutes)
    {

        return ExpectedMinutes - elapsedMinutes;


    }
    public int PreparationTimeInMinutes(int layers)
    {
        return layers * 2;
    }
    public int ElapsedTimeInMinutes(int layers, int timePassed)
    {
        return PreparationTimeInMinutes(layers) + timePassed;
    }
}