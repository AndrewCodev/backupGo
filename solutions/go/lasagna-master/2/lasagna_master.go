package lasagnamaster

const (
	noodleLayer = "noodles"
	sauceLayer  = "sauce"

	noodleGramsPerLayer = 50
	sauceLitersPerLayer = 0.2

    basePortions = 2
)

func PreparationTime(layers []string, avgTime int) int {
    timePerLayer := avgTime
    if timePerLayer == 0 {
        timePerLayer = 2
    }
    return len(layers) * timePerLayer
}

func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64

	for _, layer := range layers {
		switch layer {
		case noodleLayer:
			noodles += noodleGramsPerLayer
		case sauceLayer:
			sauce += sauceLitersPerLayer
		default:
			// other layers are intentionally ignored
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendIngredients, myIngredients []string) {
	lastFriendIndex := len(friendIngredients) - 1
	lastMyIndex := len(myIngredients) - 1

	myIngredients[lastMyIndex] = friendIngredients[lastFriendIndex]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaled := make([]float64, len(quantities))
	factor := float64(portions) / basePortions

	for i, q := range quantities {
		scaled[i] = q * factor
	}

	return scaled
}