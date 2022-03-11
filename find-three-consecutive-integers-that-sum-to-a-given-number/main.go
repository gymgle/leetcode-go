func sumOfThree(num int64) []int64 {
    x := num / 3
    if x * 3 == num {
        return []int64{x-1, x, x+1}
    }
    return []int64{}
}
