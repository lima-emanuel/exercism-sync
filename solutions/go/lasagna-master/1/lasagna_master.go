package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time <= 0 {
		return len(layers) * 2
	} else {
		return len(layers) * time
	}
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var (
		noodles int     = 0
		sauce   float64 = 0.0
	)

	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case "sauce":
			sauce += 0.2
		case "noodles":
			noodles += 50
		}
	}

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friend, mine []string) {
	mine[len(mine)-1] = friend[len(friend)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, n int) []float64 {
	var scaled []float64
	for i := 0; i < len(quantities); i++ {
		scaled = append(scaled, quantities[i]*float64(n)/2.0)
	}
	return scaled

}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
