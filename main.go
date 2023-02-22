package main

import (
	"fmt"
)

func main() {
	ns := GetNutritionScore(NutritionalData{
		Energy:  EnergyFromKcal(10),
		Sugar:   SugarGram(10),
		SFA:     SFAGram(2),
		Protein: ProteinGram(500),
		Fiber:   FiberGram(60),
		Sodium:  SodiumMilligram(4),
		Fruits:  FruitsPercent(2),
	}, Food)

	fmt.Printf("The nutritional score is: %d\n", ns)
	fmt.Printf("The nutriScore: %s\n", ns.GetNutriScore())
}
