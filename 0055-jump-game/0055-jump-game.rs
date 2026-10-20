impl Solution {
    pub fn can_jump(nums: Vec<i32>) -> bool {
        let l = nums.len();
        let mut dp = vec![-1;l];
        dp[0] = 1;
        for (i,v) in nums.iter().enumerate(){
            if dp[l-1] == 1{ return true; }
            if dp[i] == -1 { continue; }
            for k in 1..=std::cmp::min(*v as usize,l-1){
                if i + k < l {dp[i + k] = 1;}
                else{break}
            }
        }
        dp[l-1] == 1 
    }
}