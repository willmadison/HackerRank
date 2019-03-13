package main

func main() {
}

type Map struct {
	cities map[int]City
}

type City struct {
	id                           int
	HasLibrary, AccessibleByRoad bool
	AdjacentCities               map[int]City
}
