package simple

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestWorld(t *testing.T) {
	rand.NewSource(time.Now().UnixNano())
	c := new(Creature)
	c.MutateChance = 0.9
	c.ReproduceChance = 0.1

	world := NewWorld()
	world.Populate(c)
	fmt.Println("World is ready")
	for {
		time.Sleep(2*time.Second)
		a, b, c := world.Tick(100)
		fmt.Printf("===Age: %d===\nBorn: %d, Died: %d, Population increase: %d\nPopulation: %d\n", world.Age, a, b, c, world.Population)
		p := "%"
		fmt.Printf("AvgMutationRate: %.2f%s\nAvgReproductionRate: %.2f%s\n\n", world.AvgMut*100, p, world.AvgRepr*100, p)
		continue
	}
}
