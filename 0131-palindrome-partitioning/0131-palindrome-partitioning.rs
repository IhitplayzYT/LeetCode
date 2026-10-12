impl Solution {
pub fn partition(s: String) -> Vec<Vec<String>> {
    let mut ret = vec![];
    _recurse(0,vec![],s.clone(),&mut ret);
    ret
}
}


pub fn b_is_palindrome(s1: &str) -> bool{
    s1.as_bytes().iter().zip(s1.as_bytes().iter().rev()).all(|(a,b)| a == b)
}

pub fn _recurse(st:usize,mut pth:Vec<String>,s: String,ret: &mut Vec<Vec<String>>){
    let l = s.len();
    if st == l{ 
        ret.push(pth.clone());
        return;
    }
    for ed in st+1..l+1{
        if b_is_palindrome(&s[st..ed]){
            pth.push(s[st..ed].to_owned());
            _recurse(ed,pth.clone(),s.clone(),ret);
            pth.pop();
        }
    }
}