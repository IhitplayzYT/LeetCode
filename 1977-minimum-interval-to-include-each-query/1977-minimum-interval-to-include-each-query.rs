
use std::collections::BinaryHeap;
use std::cmp::Reverse;

impl Solution {
    pub fn min_interval(mut intervals: Vec<Vec<i32>>, queries: Vec<i32>) -> Vec<i32> {
        intervals.sort();
        let mut queries: Vec<(i32, usize)> = queries.into_iter().enumerate().map(|(i, q)| (q, i)).collect();
        queries.sort();
        let mut ret = vec![-1; queries.len()];
        let mut mn_hp = BinaryHeap::new();
        let mut i = 0;
        for (q, idx) in queries {
            while i < intervals.len() && intervals[i][0] <= q {
                let (st,ed) = (intervals[i][0],intervals[i][1]);
                mn_hp.push(Reverse((ed - st + 1, ed)));
                i += 1;
            }

            while let Some(Reverse((sz, ed))) = mn_hp.peek() {
                if *ed < q { mn_hp.pop();}
                else {
                    ret[idx] = *sz;
                    break;
                }
            }
        }
        ret
    }
}