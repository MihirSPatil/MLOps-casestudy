package space
import "strings"
type Planet string 
func Age(seconds float64, planet Planet) float64 {
	yearlyAge := seconds/31557600.0
    planetModifier := map[string]float64{
        "mercury": 0.2408467,
        "venus": 0.61519726,
        "earth": 1.0,
        "mars": 1.8808158,
        "jupiter": 11.862615,
        "saturn": 29.447498,
        "uranus": 84.016846,
        "neptune": 164.79132,
    }
    planetName := strings.ToLower(string(planet))
    name, ok := planetModifier[planetName]
    if !ok {
        return -1.0
        } else {
        return yearlyAge / name 
    }
}
