//checks given three values whether they form a triangle or not, and what kind of triangle
// https://golang.org/doc/effective_go.html#commentary
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
    // Pick values for the following identifiers used by the test program.
    NaT Kind = "Not a Triangle"
    Equ Kind = "Equilateral"
    Iso Kind = "Isoceles"
    Sca Kind = "Scalene"
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
    
   if a+b < c || b+c < a || a+c < b || a <=0 || b<=0 ||c<=0{
       return NaT
   }
    if a == b && b == c && c==a{
        return Equ
    } else if a!=b && b!=c && c!=a{
        return Sca
    } else {return Iso}
}
