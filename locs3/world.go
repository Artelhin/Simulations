package locs3

type World struct {
	Age       int
	Locations []*Location
}

func NewWorld(l []*Location) *World {
	w := new(World)
	w.Locations = l
	w.Age = 0
	return w
}

type TickInfo struct {
	World           int
	Population      int
	Born            int
	Died            int
	LocalPopulation map[string]int
	LocalBorn       map[string]int
	LocalDied       map[string]int
	LocalLeft       map[string]int
	LocalArrived    map[string]int
}

func (w *World) Tick() *TickInfo {
	ti := new(TickInfo)
	ti.LocalLeft = map[string]int{}
	ti.LocalDied = map[string]int{}
	ti.LocalBorn = map[string]int{}
	ti.LocalArrived = map[string]int{}
	ti.LocalPopulation = map[string]int{}
	for _, l := range w.Locations {
		ti.LocalArrived[l.Name] = l.Arrive()
		l.EatPhase()
		ti.LocalDied[l.Name], ti.LocalLeft[l.Name] = l.DeathAndMigratePhase()
		ti.LocalPopulation[l.Name] = l.Population
		ti.LocalBorn[l.Name] = l.ReproducePhase()

		ti.Population += ti.LocalPopulation[l.Name]
		ti.Born += ti.LocalBorn[l.Name]
		ti.Died += ti.LocalDied[l.Name]
	}
	w.Age += 1
	return ti
}
