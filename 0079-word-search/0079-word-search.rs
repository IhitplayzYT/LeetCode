use std::collections::HashSet;
impl Solution {
pub fn exist(board: Vec<Vec<char>>, word: String) -> bool {
    let word = word.chars().collect::<Vec<char>>();
    for (i,row) in board.iter().enumerate(){
        for (j,ch) in row.iter().enumerate(){
            if ch == &word[0]{
                let mut seen = HashSet::new();
                seen.insert((i,j));
                if Self::is_present(&board,&mut seen,&word,i,j,1){ return true}
            }
        }
    }
    false
}

pub fn is_present(board: &Vec<Vec<char>>,seen: &mut HashSet<(usize,usize)>,word: &Vec<char>,i:usize,j:usize,k:usize) -> bool{
if k == word.len(){return true}
let dirs = [(-1,0),(1,0),(0,1),(0,-1)];
let (m,n) = (board.len() as i32,board[0].len() as i32);
let (i,j) = (i as i32,j as i32);
let possible = dirs.into_iter().map(|(dx,dy)| { 
    ((i + dx).max(0).min(m-1) as usize, (j + dy).max(0).min(n-1) as usize)
}).filter(|x| !seen.contains(x)).collect::<Vec<(usize,usize)>>();
for (x,y) in possible{
    if board[x][y] == word[k]{
        seen.insert((x,y));
        if Self::is_present(board, seen, word,x,y,k+1){
            return true;
        }
        seen.remove(&(x,y));
    }
}
return false;
}
}