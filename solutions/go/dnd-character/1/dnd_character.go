package dndcharacter

import (
    "fmt";
    "math";
    "math/rand";
    "time"
)
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
    fmt.Print("Input Score:", score)
	return int(math.Floor(float64(score-10) / 2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	rand.Seed(time.Now().UnixNano()) // Seed the random generator

	rolls := make([]int, 4)
	for i := 0; i < 4; i++ {
		rolls[i] = rand.Intn(6) + 1 // Random int between 1 and 6
	}

	// Find the lowest roll
	lowest := rolls[0]
	for _, v := range rolls {
		if v < lowest {
			lowest = v
		}
	}

	// Sum all except the lowest
	sum := 0
	for _, v := range rolls {
		sum += v
	}
	sum -= lowest

	fmt.Println("Rolls:", rolls, "Lowest:", lowest, "Sum of highest 3:", sum)
	return sum
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	c := Character{
        Strength: Ability(),
    	Dexterity: Ability(),
    	Constitution: Ability(),
    	Intelligence: Ability(),
    	Wisdom: Ability(),
    	Charisma: Ability(),
    }
    c.Hitpoints = 10 + Modifier(c.Constitution)
    
    return c
}
