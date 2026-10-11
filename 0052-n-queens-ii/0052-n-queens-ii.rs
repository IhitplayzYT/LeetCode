use std::collections::HashSet;

impl Solution {

    pub fn total_n_queens(n: i32) -> i32{
        let n = n as usize;
        let mut ret = 0;
        let mut board = vec![vec!['.'; n]; n];
        Self::recurse_n_queen(&mut board, &mut ret, 0);
        ret
    }

    pub fn recurse_n_queen(board: &mut Vec<Vec<char>>,ret: &mut i32,row: usize,) {
        if row == board.len() {
            *ret += 1;
            return;
        }

        for idx in Self::valid_q_pos(board, row) {
            board[row][idx] = 'Q';
            Self::recurse_n_queen(board, ret, row + 1);
            board[row][idx] = '.';
        }
    }

    pub fn valid_q_pos(board: &Vec<Vec<char>>, row: usize) -> Vec<usize> {
        let n = board.len();
        let mut possib: HashSet<usize> = (0..n).collect();

        for i in 0..row {
            for j in 0..n {
                if board[i][j] == 'Q' {
                    possib.remove(&j);
                }
            }
        }

        possib.retain(|&col| {
            for i in 0..row {
                for j in 0..n {
                    if board[i][j] == 'Q'
                        && i.abs_diff(row) == j.abs_diff(col)
                    {
                        return false;
                    }
                }
            }
            true
        });

        let mut result: Vec<usize> = possib.into_iter().collect();
        result.sort_unstable();
        result
    }
}
