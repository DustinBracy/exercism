package lasagna

func PreparationTime(layers []string, layerTime int) int {
	if layerTime == 0 {
		layerTime = 2
	}
	return layerTime * len(layers)
}

func Quantities(layers[]string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _,v := range layers {
		if v == "sauce" {
			sauce += .2
		}
		if v == "noodles" {
			noodles += 50
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsList []string, myList []string) {
	var secretIngredient string
	common := make(map[string]bool)
	for _, v := range myList {
		common[v] = true
	}
	for _, v := range friendsList {
		if !common[v] {
			secretIngredient = v
		}
	}
	myList[len(myList) - 1] = secretIngredient
}

func ScaleRecipe(ingredients []float64, portions int) []float64 {
	var scaledRecipe []float64
	for _, v := range ingredients {
		scaledRecipe = append(scaledRecipe, v * float64(portions)/2)
	}
	return scaledRecipe
}
