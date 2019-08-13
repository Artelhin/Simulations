package simple

type FoodPool struct {
	Amount int
}

func NewFoodPool() *FoodPool {
	return new(FoodPool)
}