func sumOfSquares(nums []int) int {
    ret := 0
    l := len(nums)
    for i,v := range nums{
        if l % (i+1) == 0{
            ret += v * v
        }
    }
    return ret
}