package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	fav := []int{2,6,9}
    return fav
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
    // Check if index is out of range
    if index < 0 || index >= len(slice) {
        return -1
    }
    
    // Return the item at the specified index
    return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
    if index < 0 || index >= len(slice) {
        var newSlice []int
        newSlice = append(slice, value)
    	return newSlice
    }
	slice[index] = value
    return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	newslice := append(values, slice...)
    return newslice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
    // Check if index is out of range
    if index < 0 || index >= len(slice) {
        return slice // Return original slice if index is invalid
    }
    
    // Remove the item by appending elements before and after the index
    return append(slice[:index], slice[index+1:]...)
}