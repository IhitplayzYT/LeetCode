use std::collections::BinaryHeap;


struct KthLargest {
    hp: BinaryHeap<i32>,
    n: i32,
}

impl KthLargest {
    fn new(n: i32, nums: Vec<i32>) -> Self {
        KthLargest {n,hp: BinaryHeap::from_iter(nums.iter().map(|&n| -n))}
    }

    fn add(&mut self, val: i32) -> i32 {
        self.hp.push(-val);
        while self.hp.len() > self.n as usize {
            self.hp.pop();
        }
        -self.hp.peek().unwrap()
    }
}