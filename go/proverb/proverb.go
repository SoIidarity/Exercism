package proverb

// Proverb iterates over the rhymes array to generate the riming pattern
func Proverb(rhyme []string) []string {
	returnStrings := make([]string, len(rhyme))
	if len(rhyme) == 0 {
		return returnStrings
	}
	for i := 0; i < len(rhyme)-1; i++ {
		returnStrings[i] = "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."
	}
	returnStrings[len(returnStrings)-1] = "And all for the want of a " + rhyme[0] + "."
	return returnStrings
}
