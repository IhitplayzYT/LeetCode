impl Solution {
pub fn last_stone_weight(stones: Vec<i32>) -> i32 {
    use std::collections::BinaryHeap;
    let mut hp:BinaryHeap<i32> = stones.into_iter().collect();
    while hp.len() > 1{
        let (y,x) = (hp.pop().unwrap(),hp.pop().unwrap());
        if x != y{
            hp.push(y-x);
        }
    }
    hp.pop().unwrap_or(0)
}
}