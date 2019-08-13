package locs3

import (
	"math/rand"
)

type Creature struct {
	Age              int
	Ate              float64
	Location         *Location
	Migrated         bool
	Speed            Speed
	Weight           float64
	MutationRate     float64
	ReproductionRate float64
}

func (c *Creature) SpeedIn(l *Location) float64 {
	return c.Speed.In[l.Name]
}

func (c *Creature) Eat() {
	needed := c.SpeedIn(c.Location) * c.Weight
	if c.Location.Food.Amount < needed {
		c.Ate += c.Location.Food.Amount
		c.Location.Food.Amount = 0
	} else {
		c.Ate += needed
		c.Location.Food.Amount -= needed
	}
}

func (c *Creature) Survives() bool {
	if c.Ate >= c.Weight*c.Speed.In[c.Location.Name] {
		c.Migrated = false
		return true
	}
	return false
}

func (c *Creature) Die() bool {
	if c.Ate > 0.5*c.SpeedIn(c.Location)*c.Weight && !c.Migrated {
		if c.Migrate() {
			return false
		}
	}
	return true
}

func (c *Creature) Reproduce() []*Creature {
	multiplier := int(c.Ate / c.Weight * c.SpeedIn(c.Location))
	if multiplier == 0 {
		return nil
	}
	offsprings := make([]*Creature, 0)
	for i := 0; i < multiplier; i++ {
		offsprings = append(offsprings, generateOffspring(c))
	}
	return offsprings
}

func (c *Creature) Migrate() bool {
	if len(c.Location.Near) == 0 {
		return false
	}
	to := rand.Intn(len(c.Location.Near))
	c.Location.Near[to].Incoming = append(c.Location.Near[to].Incoming, c)
	c.Location = c.Location.Near[to]
	c.Migrated = true
	return true
}

func generateOffspring(c *Creature) *Creature {
	nc := new(Creature)
	nc.Age = 0
	nc.Ate = 0
	nc.Location = c.Location
	nc.Speed = Speed{In: make(map[string]float64,0)}
	if rand.Float64() < c.MutationRate {
		nc.MutationRate = c.MutationRate * MutateMultiplier()
		nc.ReproductionRate = c.ReproductionRate * MutateMultiplier()
		nc.Weight = c.Weight * MutateMultiplier()
		for k, v := range c.Speed.In {
			nc.Speed.In[k] = v * MutateMultiplier() + rand.Float64() / 20
		}
	} else {
		nc.MutationRate = c.MutationRate
		nc.ReproductionRate = c.ReproductionRate
		nc.Weight = c.Weight
		for k, v := range c.Speed.In {
			nc.Speed.In[k] = v
		}
	}
	return nc
}

func MutateMultiplier() float64 {
	return rand.Float64()*0.4 + 0.8
}
