package strand

func ToRNA(dna string) string {
	RNA_component :=map[rune]rune{'G':'C', 'C':'G', 'T':'A', 'A':'U'}
    rna := ""
    for _, val := range dna{
        rna += string(RNA_component[val])
    }
    return rna
}
