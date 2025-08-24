package protein
import "errors"

var codon2pro = []struct{
    codon []string
    protein string
}{
    {[]string{"AUG"},"Methionine"},
    {[]string{"UUU","UUC"},"Phenylalanine"},
    {[]string{"UUA", "UUG"}, "Leucine"},
    {[]string{"UGG"},"Tryptophan"},
	{[]string{"UAU","UAC"},"Tyrosine"},
	{[]string{"UGU","UGC"},"Cysteine"},
    {[]string{"UCU","UCC","UCA","UCG"},"Serine"},
    {[]string{"UAA","UAG","UGA"},"STOP"},
}
var ErrStop = errors.New("stop codon")
var ErrInvalidBase = errors.New("invalid RNA codon")

func FromRNA(rna string) ([]string, error) {
    
    codonLen := 3
    var result []string

    for i := 0; i < len(rna); i += codonLen {
        codon := rna[i : i+codonLen]
        protein, err := FromCodon(codon)

        if err == ErrStop {
            return result, nil
        }
        if err == ErrInvalidBase {
            return nil, err
        }

        result = append(result, protein)
    }

    return result, nil
}

func FromCodon(codon string) (string, error) {
    for _, val := range codon2pro{
        for _, sym := range val.codon{
            if sym == codon{
                if val.protein == "STOP"{
                    return "", ErrStop
                }
                return val.protein, nil
            }
        }
    }
    return "", ErrInvalidBase
}
