use std::collections::*;
impl Solution {
    pub fn pacific_atlantic(heights: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let (mut s1,mut s2) = (HashSet::new(),HashSet::new());
        let (m,n) = (heights.len(),heights[0].len());
        // First set to process atlantic reachable nodes second for pacific reachable nodes
        let (mut q1,mut q2)= (VecDeque::new(),VecDeque::new());
        for i in 0..m{
            q1.push_back((i,0));
            q2.push_back((i,n-1));
        }
        for j in 0..n{
            q1.push_back((0,j));
            q2.push_back((m-1,j));
        }

        bfs(&mut q1,&mut s1,&heights);
        bfs(&mut q2,&mut s2,&heights);

        s1.intersection(&s2).into_iter().map(|(x,y)| vec![*x as i32,*y as i32]).collect()
    }
}

pub fn bfs(que: &mut VecDeque<(usize,usize)>,s: &mut HashSet<(usize,usize)>,heights:&Vec<Vec<i32>>){
    let dirs = [(0,1),(1,0),(-1,0),(0,-1)];
    while !que.is_empty(){
        let l = que.len();
        for _ in (0..l){
            let v = que.pop_front().unwrap();
            s.insert(v);
            for d in &dirs {
                let (x, y) = (v.0 as i32 + d.0, v.1 as i32 + d.1);
                if x < 0 || y < 0 || x >= heights.len() as i32 || y >= heights[0].len() as i32{ continue;}
                let u = (x as usize, y as usize);
                if s.contains(&u) || heights[u.0][u.1] < heights[v.0][v.1] { continue; }
                s.insert(u);
                que.push_back(u);
            }
        }
    }
}
    