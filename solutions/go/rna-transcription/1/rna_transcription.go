package strand

func ToRNA(dna string) string {
    out := ""
    for _, d := range dna {
        switch d {
            case 'C':
            	out += "G"
            case 'T':
            	out += "A"
            case 'A':
            	out += "U"
            case 'G':
            	out += "C"
            default:
            	out += ""
        }
    }
    return out
}
