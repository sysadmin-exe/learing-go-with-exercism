package meteorology

import "fmt" 

type Stringer interface {
    String() string
}

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

func (tu  TemperatureUnit) String() string {
    units := []string{"°C", "°F"}
	return units[tu]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t  Temperature) String() string {
    return fmt.Sprintf("%v %v", t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

func (su  SpeedUnit) String() string {
    units := []string{"km/h", "mph"}
	return units[su]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (s  Speed) String() string {
    return fmt.Sprintf("%v %v", s.magnitude, s.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (m  MeteorologyData) String() string {
    return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity", m.location, m.temperature, m.windDirection, m.windSpeed, m.humidity)
}