package locs3

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestWorld(t *testing.T) {

	rand.NewSource(time.Now().UnixNano())

	locs := []string{"beach", "forest", "sea"}
	locations := make([]*Location, 0)
	for _, s := range locs {
		locations = append(locations, NewLocation(s))
	}
	connectLocations(locations[0], locations[1])
	connectLocations(locations[1], locations[2])
	world := NewWorld(locations)

	c := new(Creature)
	c.Speed = Speed{In: make(map[string]float64, 0)}
	c.Speed.In["sea"] = 1
	c.Speed.In["beach"] = 0
	c.Speed.In["forest"] = 0
	c.Weight = 0.5
	c.ReproductionRate = 0.5
	c.MutationRate = 0.2

	world.Locations[2].Populate(c)

	for {
		time.Sleep(3 * time.Second)
		for _, l := range world.Locations {
			l.Food.Amount = 10*rand.Float64() + 10*rand.Float64() + 10*rand.Float64()
		}
		ti := world.Tick()
		fmt.Printf("====Age: %d====\n", world.Age)
		for i, s := range locs {
			fmt.Printf("***%s***\nPopulation %d\nArrived: %d\nBorn: %d\nDied: %d\nLeft: %d\nFood left: %.2f\n\n",
				s, ti.LocalPopulation[s], ti.LocalArrived[s], ti.LocalBorn[s], ti.LocalDied[s], ti.LocalLeft[s], locations[i].Food.Amount)
		}
		fmt.Printf("Total population: %d\nTotal born: %d\nTotal died: %d\n\n", ti.Population, ti.Born, ti.Died)
	}

}
