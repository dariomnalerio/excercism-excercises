package lasagna

func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		return 2 * len(layers)
	}
	return timePerLayer * len(layers)
}

func Quantities(layers []string) (int, float64) {
	sauceLayers := float64(0)
	noodlesLayers := 0

	for _, v := range layers {
		if v == "sauce" {
			sauceLayers++
		} else if v == "noodles" {
			noodlesLayers++
		}
	}

	noodles := noodlesLayers * 50
	sauce := sauceLayers * 0.2

	return noodles, sauce
}

func AddSecretIngredient(ing1, ing2 []string) {
	ing2[len(ing2)-1] = ing1[len(ing1)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	quantitiesCopy := make([]float64, len(quantities))
	copy(quantitiesCopy, quantities)

	for i := range quantitiesCopy {
		quantitiesCopy[i] *= float64(portions) / 2
	}

	return quantitiesCopy
}
