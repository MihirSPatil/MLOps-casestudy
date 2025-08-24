package dndcharacter
import ("math/rand";
        "sort";
        "math")

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor((float64(score)-10)/2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
    var vals []int
    sum:= 0
	for i:=0; i< 4; i++{
        vals = append(vals, rand.Intn(6)+1)
    }
    sort.Ints(vals)
    for i:= len(vals)-1; i>0; i--{
        sum += vals[i]
    }
    return sum
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	char := Character{
    	Strength : Ability(),
    	Dexterity : Ability(),
		Constitution : Ability(),
		Intelligence : Ability(),
		Wisdom : Ability(),    
		Charisma : Ability(),    
        }
    	char.Hitpoints = 10 + Modifier(char.Constitution)
    return char
}
