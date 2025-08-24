package lasagna
import "fmt"
// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int{
    timeVal :=2
    if timePerLayer > 0{
        timeVal = timePerLayer
    }
    return len(layers) * timeVal
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64){
    noodleCount :=0
    sauceCount :=0.0
    sauceVal :=0.2
    noodleVal:=50
    for idx:=0; idx < len(layers); idx++{
        if layers[idx] == "sauce"{
            sauceCount +=1
        } else if layers[idx] == "noodles"{
            noodleCount +=1
        }
    }
    return (noodleCount * noodleVal), (sauceCount * sauceVal)
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendRecipe []string, ownRecipe []string) {
	ownRecipe[len(ownRecipe)-1] = friendRecipe[len(friendRecipe)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(twoPortionQuantity []float64, portionsToCook int) []float64 {
    var quantityScaler float64 = float64(portionsToCook)/2.0
    var scaledQuantity []float64
    fmt.Println(quantityScaler)
    fmt.Println(twoPortionQuantity)
    for idx:=0; idx < len(twoPortionQuantity); idx++{
        scaledQuantity = append(scaledQuantity, twoPortionQuantity[idx] * quantityScaler)
    }
    return scaledQuantity
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
