use std::collections::BinaryHeap;
use std::cmp::Reverse;

struct MedianFinder {
    l: BinaryHeap<i32>,
    r: BinaryHeap<Reverse<i32>>,
}

impl MedianFinder {
    fn new() -> Self { MedianFinder {l: BinaryHeap::new(),r: BinaryHeap::new()} }
    
    fn add_num(&mut self, num: i32) {
        self.l.push(num);
        // Add the biggest elem in l to r
        if let Some(val) = self.l.pop() {
            self.r.push(Reverse(val));
        }
        // if l_r > l_l then take smallest in r and put in l
        if self.r.len() > self.l.len() {
            if let Some(Reverse(val)) = self.r.pop() {
                self.l.push(val);
            }
        }

    }
    
    fn find_median(&self) -> f64 {
        if self.l.len() > self.r.len() {
            *self.l.peek().unwrap() as f64
        } else {
            let a = *self.l.peek().unwrap();
            let Reverse(b) = *self.r.peek().unwrap();
            (a as f64 + b as f64) / 2.0
        }
    }
}