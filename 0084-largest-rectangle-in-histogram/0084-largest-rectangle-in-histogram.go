

func largestRectangleArea(heights []int) int {
    idx_stck := []int{}
    mx := 0
    l := len(heights)
    for i := 0; i <= l; i++ {
        curr := 0
        if i < l { curr = heights[i] }
        for len(idx_stck) > 0 && heights[idx_stck[len(idx_stck)-1]] > curr {
            h := heights[idx_stck[len(idx_stck)-1]]
            idx_stck = idx_stck[:len(idx_stck)-1]
            w := i
            if len(idx_stck) > 0 { w = i - idx_stck[len(idx_stck)-1] - 1 }
            mx = max(mx, h*w)
        }
        idx_stck = append(idx_stck, i)
    }
    return mx
}