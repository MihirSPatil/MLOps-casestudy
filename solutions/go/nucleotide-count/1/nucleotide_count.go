package dna
import "fmt"
import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA []rune

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
    fmt.Println("string", string(d))
	var h Histogram = Histogram{'A':0, 'G':0, 'T':0, 'C':0}
    for _, rna := range d{
        if _, exists := h[rna]; exists{
            h[rna] +=1
        }else {return nil, errors.New("Wrong input")}
            }
	return h, nil
}
