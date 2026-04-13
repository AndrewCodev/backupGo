public class Lasagna {

    private static final int EXPECTED_MINUTES_IN_OVEN = 40;
    private static final int MINUTES_PER_LAYER = 2;
    
    public int expectedMinutesInOven(){
        return EXPECTED_MINUTES_IN_OVEN;
    }
    
    public int remainingMinutesInOven(int minutesInOven){
        if (minutesInOven < 0) {
            throw new IllegalArgumentException("Minutes in oven cannot be negative.");
        }
        return EXPECTED_MINUTES_IN_OVEN - minutesInOven;
    }
    
    public int preparationTimeInMinutes(int layers){
        if (layers < 0) {
            throw new IllegalArgumentException("Number of layers cannot be negative.");
        }
        return layers * MINUTES_PER_LAYER;
    }
    
    public int totalTimeInMinutes(int layers, int minutesInOven){
         if (layers < 0) {
            throw new IllegalArgumentException("Number of layers cannot be negative.");
        }
        return preparationTimeInMinutes(layers) + minutesInOven;
    }
}
