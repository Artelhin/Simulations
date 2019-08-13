package locs3

import "sort"

type Location struct {
	Name       string
	Food       *FoodPool
	Population int
	Near       []*Location
	Creatures  []*Creature
	Incoming   []*Creature
}

func NewLocation(name string) *Location {
	return &Location{
		Name: name,
		Food: NewFoodPool(),
		Near: make([]*Location, 0),
	}
}

func connectLocations(l1, l2 *Location) {
	l1.Near = append(l1.Near, l2)
	l2.Near = append(l2.Near, l1)
}

func (l *Location) Arrive() int {
	count := 0
	for _, c := range l.Incoming {
		l.Creatures = append(l.Creatures, c)
		count++
	}
	l.Incoming = make([]*Creature, 0)
	return count
}

func (l *Location) EatPhase() {
	sort.SliceStable(l.Creatures, func(i, j int) bool {
		return l.Creatures[i].Speed.In[l.Creatures[i].Location.Name] > l.Creatures[j].Speed.In[l.Creatures[j].Location.Name]
	})
	for l.Food.Amount > 0 && (l.Creatures != nil && len(l.Creatures) != 0){
		for _, c := range l.Creatures {
			c.Eat()
		}
	}
}

//возвращает количество умерших и покинувших локацию существ за фазу
func (l *Location) DeathAndMigratePhase() (int, int) {
	var (
		died int
		left int
	)
	for i, c := range l.Creatures {
		if !c.Survives() {
			if c.Die() {
				died++
			} else {
				left++
			}
			l.Creatures[i] = nil
		}
	}

	newPopulation := make([]*Creature, 0)
	for _, c := range l.Creatures {
		if c != nil {
			newPopulation = append(newPopulation, c)
		}
	}
	l.Creatures = newPopulation
	l.Population = len(newPopulation)

	return died, left
}

//возвращает количество рожденных за фазу
func (l *Location) ReproducePhase() int {
	born := 0
	newBorn := make([]*Creature, 0)
	for _, c := range l.Creatures {
		offsprings := c.Reproduce()
		if offsprings == nil {
			continue
		}
		for _, o := range offsprings {
			born++
			newBorn = append(newBorn, o)
		}
	}
	for _, c := range newBorn {
		l.Creatures = append(l.Creatures, c)
	}
	return born
}

func (l *Location) Populate(c *Creature) {
	l.Creatures = append(l.Creatures, c)
	c.Location = l
}
