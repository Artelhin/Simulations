package locs3

type FoodPool struct {
	Amount float64
}

func NewFoodPool() *FoodPool {
	return new(FoodPool)
}
