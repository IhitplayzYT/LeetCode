use std::collections::{VecDeque, HashSet};

impl Solution {
    pub fn oranges_rotting(grid: Vec<Vec<i32>>) -> i32 {
        let m = grid.len();
        if m == 0 {return 0;}
        let n = grid[0].len();
        if n == 0 {return 0;}

        let dir = [(1, 0), (0, 1), (-1, 0), (0, -1)];
        let mut que = VecDeque::new();
        let mut time = 0;
        let mut seen = HashSet::new();

        for (i, p) in grid.iter().enumerate() {
            for (j, v) in p.iter().enumerate() {
                if *v == 2 {
                    que.push_back((i, j));
                    seen.insert((i, j));
                }
            }
        }
        let mut fresh = grid.iter().flatten().filter(|&&v| v == 1).count();
        while !que.is_empty() {
            let l = que.len();
            let mut rotted = false;
            for _ in 0..l {
                let (x, y) = que.pop_front().unwrap();
                for (dx, dy) in &dir {
                    let (nx, ny) = (x as i32 + dx, y as i32 + dy);
                    if nx < 0 || nx >= m as i32 || ny < 0 || ny >= n as i32{ continue; }
                    let (nx, ny) = (nx as usize, ny as usize);
                    if grid[nx][ny] != 1 || seen.contains(&(nx, ny)) { continue; }
                    fresh -= 1;
                    seen.insert((nx, ny));
                    que.push_back((nx, ny));
                    rotted = true;
                }
            }

            if rotted {
                time += 1;
            }
        }

        if fresh > 0 { -1 } else { time }
    }
}