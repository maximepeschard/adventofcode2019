package rocket

// ModuleFuel returns the fuel requirement given a module's mass
func ModuleFuel(mass int) int {
	return mass/3 - 2
}

// TotalModuleFuel returns the fuel requirement given a module's mass
// when also taking into account the mass of the added fuel.
func TotalModuleFuel(mass int) int {
	total := 0
	fuel := ModuleFuel(mass)
	for fuel > 0 {
		total += fuel
		fuel = ModuleFuel(fuel)
	}

	return total
}
