package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTime int) int{
    if avgTime == 0 {
        avgTime = 2
    } else {
        avgTime = avgTime
    }
	totalPrepTime := len(layers) * avgTime
    return totalPrepTime
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64){
    noodlesQuantity := 0
    sauceQuantity := 0
    for _, item := range layers {
        if item == "noodles" {
            noodlesQuantity++
        }
        if item == "sauce" {
            sauceQuantity++
        }
    }
    return int(noodlesQuantity * 50), float64(sauceQuantity) * 0.2
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
    secretIngredient := friendsList[len(friendsList) - 1]
    for i := 0; i < len(myList); i++ {
        if myList[i] == "?" {
            myList[i] = secretIngredient
        }
    }
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, numberOfPortions int) []float64{
    var scaledQuantities []float64
    for q := 0; q < len(quantities); q++ {
        scaledQuantities = append(scaledQuantities, (quantities[q] * float64(numberOfPortions)) / 2.0)
    }
    return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
