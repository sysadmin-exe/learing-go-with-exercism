package etl

import (
    "fmt"
    "strings"
)
func Transform(in map[int][]string) map[string]int {
	fmt.Println("Input Map:", in )

    out := make(map[string]int)

    for k, v := range in {
        switch k {
            case 1:
            	for _, a := range v {
                    newkey := strings.ToLower(a)
                    out[newkey] = 1
                }
            case 2:
            	for _, a := range v {
                    newkey := strings.ToLower(a)
                    out[newkey] = 2
                }
            case 3:
            	for _, a := range v {
                    newkey := strings.ToLower(a)
                    out[newkey] = 3
                }
            case 4:
            	for _, a := range v {
                    newkey := strings.ToLower(a)
                    out[newkey] = 4
                }
            case 5:
            	for _, a := range v {
                    newkey := strings.ToLower(a)
                    out[newkey] = 5
                }
            case 8:
            	for _, a := range v {
                    newkey := strings.ToLower(a)
                    out[newkey] = 8
                }
            case 10:
            	for _, a := range v {
                    newkey := strings.ToLower(a)
                    out[newkey] = 10
                }
        }
    }
	
    
	fmt.Println("Output Map:", out )
    return out
}
