package speed
import "fmt"
// TODO: define the 'Car' type struct
type Car struct {
    battery int
    batteryDrain int
    speed int
    distance int
}
// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
    return Car {
        battery: 100,
        batteryDrain: batteryDrain,
        speed: speed,
    }
}

// TODO: define the 'Track' type struct
type Track struct {
    distance int
}
// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track {
        distance: distance,
    }
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
    switch {
        case car.battery >= car.batteryDrain:
        	return Car {
                batteryDrain: car.batteryDrain,
                speed: car.speed,
                distance: car.distance + car.speed,
                battery: car.battery - car.batteryDrain,
            }
        default:
            return car 
    }
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
    maxDistance := (car.battery / car.batteryDrain) * car.speed
    fmt.Println(maxDistance)
    if maxDistance >= track.distance {
        return true
    } else {
        return false
    }
}
