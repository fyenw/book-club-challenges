package algorithm

func RepresentGraph(airports []string, routes [][2]string) map[string][]string {
	airportPartners := map[string][]string{}
	for _, airport := range airports {
		airportPartners[airport] = []string{}
	}
	for _, route := range routes {
		airportPartners[route[0]] = append(airportPartners[route[0]], route[1])
		airportPartners[route[1]] = append(airportPartners[route[1]], route[0])
	}
	return airportPartners
}

func RouteExists(start, end string, airportPartners map[string][]string) bool {
	consideredAirports := map[string]bool{start: true}
	currentAirports := map[string]bool{start: true}
	for len(currentAirports) != 0 {
		newAirports := map[string]bool{}
		for airport := range currentAirports {
			for _, partner := range airportPartners[airport] {
				if partner == end {
					return true
				}
				if !consideredAirports[partner] {
					consideredAirports[partner] = true
					newAirports[partner] = true
				}
			}
		}
		currentAirports = newAirports
	}
	return false
}
