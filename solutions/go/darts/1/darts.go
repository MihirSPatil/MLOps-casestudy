package darts
import "math"

func Score(x, y float64) int {
	radius:= math.Sqrt(x*x +y*y)
    var score int
    switch {
        case radius <=1.0: score = 10
        case radius <= 5.0: score = 5
        case radius <= 10.0: score = 1
        default: score = 0
    }
    return score
}
