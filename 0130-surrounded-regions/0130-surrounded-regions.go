func solve(board [][]byte)  {
    n := len(board)
    m := len(board[0])
    if n < 3 || m < 3 { return }

    var dfs func(r,c int)
    dfs = func(r int, c int) {
        board[r][c] = '@'
        if r > 0 && board[r - 1][c] == 'O' { dfs(r - 1, c) }
        if r < n - 1 && board[r + 1][c] == 'O' { dfs(r + 1, c) }
        if c > 0 && board[r][c - 1] == 'O' { dfs(r, c - 1) }
        if c < m - 1 && board[r][c + 1] == 'O' { dfs(r, c + 1)}
    }

    for r := 0; r < n; r++ {
        if board[r][0] == 'O' { dfs(r, 0) }
        if board[r][m - 1] == 'O' { dfs(r, m - 1) }
    }
    for c := 1; c < m - 1; c++ {
        if board[0][c] == 'O' { dfs(0, c)}
        if board[n - 1][c] == 'O' { dfs(n - 1, c) }
    }

    for r := 0; r < n; r++ {
        for c := 0; c < m; c++ {
            if board[r][c] == 'O' { 
                board[r][c] = 'X'
            } else if board[r][c] == '@' {
                board[r][c] = 'O'
            }
        }
    }    
}

