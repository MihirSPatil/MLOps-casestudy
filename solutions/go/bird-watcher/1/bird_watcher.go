package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var sum int
    for idx :=0; idx < len(birdsPerDay); idx++{
        sum+= birdsPerDay[idx]
    }
    return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    var startIdx int
    var endIdx int
    startIdx = 7*(week-1)
    endIdx = startIdx+7
    
    birdCount := TotalBirdCount(birdsPerDay[startIdx:endIdx])
    return birdCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for idx:=0; idx < len(birdsPerDay); idx++{
        if idx %2 ==0{
            birdsPerDay[idx] = birdsPerDay[idx] + 1
        }
    }
    return birdsPerDay
}
