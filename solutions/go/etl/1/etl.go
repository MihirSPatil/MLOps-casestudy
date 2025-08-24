package etl
import ("strings")

func Transform(in map[int][]string) map[string]int {
	letterval := make(map[string]int)
    for val, str := range(in){
        for _, s := range str{
            lowStr := strings.ToLower(s)
        	letterval[lowStr] = val
        }
    }
    return letterval
}
