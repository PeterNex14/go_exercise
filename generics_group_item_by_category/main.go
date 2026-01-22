package main

type Categorizable interface {
	Category() string
}

type Weapon struct {
	Name string
	Cat  string
}

type Potion struct {
	Name string
	Cat  string
}

func (w Weapon) Category() string { return w.Cat }
func (p Potion) Category() string  { return p.Cat }

func groupByCategory[T Categorizable](items []T) map[string][]T {
	maped := make(map[string][]T)
	
	if len(items) == 0 {
		return maped
	}

	for _, value := range items {
		maped[value.Category()] = append(maped[value.Category()], value)
	}
	
	return maped
}
