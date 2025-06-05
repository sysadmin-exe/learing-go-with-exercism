package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return ("Welcome to my party, " + name + "!")
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	message := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
    return message
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	welcome := fmt.Sprintf("Welcome to my party, %s!\n", name)
    dir := fmt.Sprintf("You have been assigned to table %.3d. Your table is %s, exactly %.1f meters from here.\n", table, direction, distance)
    nei := fmt.Sprintf("You will be sitting next to %s.", neighbor)
    message := welcome + dir + nei
    return message
}
