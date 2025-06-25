package space

import "fmt"

type Planet string


func Age(seconds float64, planet Planet) float64 {
    earthyears := map[Planet]float64{
        "Mercury": 0.2408467, 
        "Venus": 0.61519726, 
        "Earth": 1.0,
        "Mars": 1.8808158, 
        "Jupiter": 11.862615, 
        "Saturn": 29.447498, 
        "Uranus": 84.016846, 
        "Neptune": 164.79132, 
    }
    
	fmt.Println("Input_Seconds:", seconds, "Input_Planet:", planet)
	_, exists := earthyears[planet]
    if exists {
        return seconds / (earthyears[planet] * 31557600)
    } else {
        return -1.0
    }
    
    
}