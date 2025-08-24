package strain

func Keep [T any](collection []T, predicate func(T) bool) []T{
    var kvals []T
    for _, val := range collection{
        if predicate(val) == true{
            kvals = append(kvals, val)
        }
    }
    return kvals
}

func Discard [T any](collection []T, predicate func(T) bool) []T{
    var dvals []T
    for _, val := range collection{
        if predicate(val) == false{
            dvals = append(dvals, val)
        }
    }
    return dvals
}
