package space

import "math"

type Planet string

func calcEarthYearsWithSeconds(seconds float64) float64 {
	return seconds / (60 * 60 * 24 * 365.25)
}

func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case "Mercury":
		factor := 0.2408467
		earthYear := calcEarthYearsWithSeconds(seconds)
		return math.Round(earthYear/factor*100) / 100
	case "Venus":
		factor := 0.61519726
		earthYear := calcEarthYearsWithSeconds(seconds)
		return math.Round(earthYear/factor*100) / 100
	case "Earth":
		factor := 1.0
		earthYear := calcEarthYearsWithSeconds(seconds)
		return math.Round(earthYear/factor*100) / 100
	case "Mars":
		factor := 1.8808158
		earthYear := calcEarthYearsWithSeconds(seconds)
		return math.Round(earthYear/factor*100) / 100
	case "Jupiter":
		factor := 11.862615
		earthYear := calcEarthYearsWithSeconds(seconds)
		return math.Round(earthYear/factor*100) / 100
	case "Saturn":
		factor := 29.447498
		earthYear := calcEarthYearsWithSeconds(seconds)
		return math.Round(earthYear/factor*100) / 100
	case "Uranus":
		factor := 84.016846
		earthYear := calcEarthYearsWithSeconds(seconds)
		return math.Round(earthYear/factor*100) / 100
	case "Neptune":
		factor := 164.79132
		earthYear := calcEarthYearsWithSeconds(seconds)
		return math.Round(earthYear/factor*100) / 100
	default:
		return 0.0
	}
}
