package simple

import (
	"math/rand"
)

type World struct {
	Age int
	AvgMut float64
	AvgRepr float64
	FoodPool *FoodPool
	Creatures []*Creature
	Population int
}

func NewWorld() *World {
	return &World{
		Age: 0,
		FoodPool: NewFoodPool(),
		Creatures: make([]*Creature, 0),
	}
}

func (w *World) Tick(foodAmount int) (int, int, int) {
	w.FoodPool.Amount += foodAmount
	w.Age++

	for w.FoodPool.Amount > 0 {
		w.Creatures[rand.Intn(len(w.Creatures))].Eat(w.FoodPool)
	}

	born := 0
	died := 0
	for i, c := range w.Creatures {
		if !c.Survives() {
			c.Die(w.FoodPool)
			died++
			w.Creatures[i] = nil
		} else {
			if c.Reproduces() {
				w.Creatures = append(w.Creatures, c.Reproduce())
				born++
			}
		}
	}

	newPopulation := make([]*Creature, 0)
	for _, c := range w.Creatures {
		if c != nil {
			newPopulation = append(newPopulation, c)
		}
	}
	w.Creatures = newPopulation
	w.Population = len(w.Creatures)

	r := 0.0
	m := 0.0
	for _, c := range w.Creatures {
		m += c.MutateChance
		r += c.ReproduceChance
	}
	w.AvgRepr = r / float64(w.Population)
	w.AvgMut = m / float64(w.Population)

	return born, died, born-died
}

func (w *World) Populate(c *Creature) {
	w.Creatures = append(w.Creatures, c)
}