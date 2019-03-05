package strand

var dnaToRnaMap = map[byte]byte{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

//ToRNA returns the RNA conversion from a DNA string
func ToRNA(dna string) string {
	rna := ""
	for i := 0; i < len(dna); i++ {
		rna += string(dnaToRnaMap[dna[i]])
	}
	return rna
}
