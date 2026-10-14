impl Solution {
pub fn find_kth_largest(nums: Vec<i32>,mut k: i32) -> i32 {
    use std::collections::BinaryHeap;
    let mut hp: BinaryHeap<i32> = nums.into_iter().collect();
    while k > 1{
        if hp.pop().is_none(){
            return -1;
        }
        k -= 1;
    }
    hp.pop().unwrap_or(-1)
}
}