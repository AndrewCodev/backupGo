public class Lasagna {

    private static final int EXPECTED_MINUTES_IN_OVEN = 40;
    private static final int MINUTES_PER_LAYER = 2;
    
    // TODO: define the 'expectedMinutesInOven()' method
    public int expectedMinutesInOven(){
        return EXPECTED_MINUTES_IN_OVEN;
    }
    
    // TODO: define the 'remainingMinutesInOven()' method
    public int remainingMinutesInOven(int minutesInOven){
        if (minutesInOven < 0) {
            throw new IllegalArgumentException("Minutes in oven cannot be negative.");
        }
        return EXPECTED_MINUTES_IN_OVEN - minutesInOven;
    }
    
    // TODO: define the 'preparationTimeInMinutes()' method
    public int preparationTimeInMinutes(int layers){
        if (layers < 0) {
            throw new IllegalArgumentException("Number of layers cannot be negative.");
        }
        return layers * MINUTES_PER_LAYER;
    }
    
    // TODO: define the 'totalTimeInMinutes()' method
    public int totalTimeInMinutes(int layers, int minutesInOven){
         if (layers < 0) {
            throw new IllegalArgumentException("Number of layers cannot be negative.");
        }
        return preparationTimeInMinutes(layers) + minutesInOven;
    }
}
