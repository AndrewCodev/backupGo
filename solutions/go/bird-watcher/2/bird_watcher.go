package birdwatcher

const daysPerWeek = 7 

func TotalBirdCount(birdsPerDay []int) int {
    sum := 0
    for _, birds := range birdsPerDay {
        sum += birds
    }
    return sum
}

func BirdsInWeek(birdsPerDay []int, week int) int {
    var sum int
    firstDay := (week - 1) * daysPerWeek
    lastDay := firstDay + daysPerWeek

    for i := firstDay; i < lastDay; i++{
        sum += birdsPerDay[i]
    }
    return sum
}

func FixBirdCountLog(birdsPerDay []int) []int {
    n := len(birdsPerDay)
    for i := 0; i < n; i += 2{
        birdsPerDay[i]++
    }
    return birdsPerDay
}