type Pt struct {
    x, y int
}

func numIslands(grid [][]byte) int {
    if len(grid) == 0 { return 0 }
    seen := make(map[Pt]bool)
    m, n := len(grid), len(grid[0])
    dz := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
    var f func(int, int)
    f = func(x, y int) {
        if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] != '1' { return }
        pt := Pt{x: x, y: y}
        if seen[pt] { return }
        seen[pt] = true
        for _, v := range dz {
            f(x+v[0], y+v[1])
        }
    }

    cnt := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            pt := Pt{x: i, y: j}
            if grid[i][j] == '1' && !seen[pt] {
                cnt+=1
                f(i, j)
            }
        }
    }

    return cnt
}