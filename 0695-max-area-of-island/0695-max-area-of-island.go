type Pt struct {
    x, y int
}

func maxAreaOfIsland(grid [][]int) int {
    if len(grid) == 0 { return 0 }
    seen := make(map[Pt]bool)
    m, n := len(grid), len(grid[0])
    dz := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
    var f func(int, int) int
    f = func(x, y int) int {
        if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 { return 0 }
        pt := Pt{x: x, y: y}
        if seen[pt] { return 0 }
        seen[pt] = true
        area := 1
        for _, v := range dz {
            area += f(x+v[0], y+v[1])
        }
        return area
    }

    mx := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            pt := Pt{x: i, y: j}
            if grid[i][j] == 1 && !seen[pt] {
                mx = max(mx, f(i, j))
            }
        }
    }
    return mx
}