impl Solution {
pub fn k_closest(points: Vec<Vec<i32>>,mut k: i32) -> Vec<Vec<i32>> {
    use std::collections::BinaryHeap;
    use std::cmp::Reverse;
    let mut heap: BinaryHeap<Reverse<(i32,Vec<i32>)>> = points.into_iter().map(|pt| Reverse((pt[0].pow(2)+ pt[1].pow(2) ,pt))).collect();
    let mut ret = vec![];
    while k != 0{
        if let Some(Reverse(pt)) = heap.pop(){
            ret.push(pt.1);         
        }else{
            break;
        }
        k -= 1 
    }
    ret    
}
}