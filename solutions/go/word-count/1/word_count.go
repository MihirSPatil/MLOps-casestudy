package wordcount
import ("regexp"
       "strings")
type Frequency map[string]int

func WordCount(phrase string) Frequency {
	dict := make(Frequency)
    re := regexp.MustCompile(`[";:,.!?\s]+`)
    phrase_ := strings.ToLower(phrase)
    valid := regexp.MustCompile(`[a-z0-9]+(?:'[a-z0-9]+)*$`)
    splits := re.Split(phrase_, -1)
    for _, word := range splits{
        word = strings.TrimPrefix(word,"'")
        word = strings.TrimSuffix(word,"'")
        if valid.MatchString(word){
            dict[word] +=1
        }
    }
    return dict
}
