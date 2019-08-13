package simple

import (
	"math/rand"
)

type Creature struct {
	Age int
	Ate int
	MutateChance float64
	ReproduceChance float64
}

func (c *Creature) Eat(from *FoodPool) {
	/*if c.Ate > 2 {
		return
	}*/
	c.Age++
	c.Ate++
	from.Amount--
	//log.Printf("creature ate %d, food left: %d", c.Ate, from.Amount)
}

func (c *Creature) Survives() bool {
	res := c.Ate >= 1
	return res
}

func (c *Creature) Die(f *FoodPool) {
	//no actions
}

func (c *Creature) Reproduces() bool {
	a := rand.Float64()
	//log.Printf("cmp: %v, %v, %v, %v, %v", a, c.ReproduceChance*float64(c.Ate), c.ReproduceChance, c.Ate, float64(c.Ate))
	if a < c.ReproduceChance*float64(c.Ate - 1) {
		return true
	}
	c.Ate = 0
	return false
}

func (c *Creature) Reproduce() *Creature {
	nc := new(Creature)
	nc.Age = 0
	nc.Ate = 0
	if rand.Float64() < c.MutateChance {
		nc.MutateChance = (3.0 * c.MutateChance + rand.Float64())/4.0
		nc.ReproduceChance = (3.0 * c.ReproduceChance + rand.Float64())/4.0
	} else {
		nc.MutateChance = c.MutateChance
		nc.ReproduceChance = c.ReproduceChance
	}
	c.Ate = 0
	return nc
}