use std::cmp::min;
use std::collections::BTreeMap;
struct DetectSquares {
    pub storage: BTreeMap<(i32,i32),i32>
}

impl DetectSquares {

    fn new() -> Self {
        Self { storage: BTreeMap::new() }        
    }
    
    fn add(&mut self, point: Vec<i32>) {
        *self.storage.entry((point[0],point[1])).or_insert(0) += 1;
    }
    
    fn count(&self, point: Vec<i32>) -> i32 {
        let (x,y) = (point[0],point[1]);
        let mut ret = 0_i32;
        for (&(x2, y2), &cnt) in &self.storage {
            let dx = x2 - x;
            let dy = y2 - y;

            if dx == 0 || dx.abs() != dy.abs() { continue; }

            let a = self.storage.get(&(x2, y)).copied().unwrap_or(0);
            let b = self.storage.get(&(x, y2)).copied().unwrap_or(0);
            ret += cnt * a * b;
        }
        ret
    }
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * let obj = DetectSquares::new();
 * obj.add(point);
 * let ret_2: i32 = obj.count(point);
 */