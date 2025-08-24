package resistorcolor
// import "fmt"
import "strings"
// Colors returns the list of all colors.
func Colors() []string {
	colors:=[]string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
    return colors
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
    colors := Colors()
    var value int
    for i:=0; i<=len(colors)-1; i++{
        if strings.ToLower(color) == colors[i]{
            value = i
        } 
    }
    return value
}
