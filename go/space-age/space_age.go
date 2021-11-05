package space

type Planet string

const multiplier = 365.25 * 24 * 60 * 60

var OrbitalPeriods = map[Planet]float64{
	"Mercury": 0.2408467 * multiplier,
	"Venus":   0.61519726 * multiplier,
	"Earth":   1 * multiplier,
	"Mars":    1.8808158 * multiplier,
	"Jupiter": 11.862615 * multiplier,
	"Saturn":  29.447498 * multiplier,
	"Uranus":  84.016846 * multiplier,
	"Neptune": 164.79132 * multiplier,
}

func Age(seconds float64, planet Planet) float64 {
	return seconds / OrbitalPeriods[planet]
}
