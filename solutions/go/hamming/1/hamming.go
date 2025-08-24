package hamming
import "errors"
func Distance(a, b string) (int, error) {
	if len(a) != len(b){
        return 0, errors.New("The length of both the DNA sequences needs to be the same ")
    }
    ham_dist := 0
    
    for i := range a{
        if a[i] != b[i]{
            ham_dist+=1
        }
    }
    return ham_dist, nil
}
