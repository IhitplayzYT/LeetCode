impl Solution {
    pub fn multiply(num1: String, num2: String) -> String {
        if num1 == "0" || num2 == "0" { return "0".to_string();}

        let (mut n1, mut n2) = (num1.as_bytes().to_vec(),num2.as_bytes().to_vec());
        n1.reverse();
        n2.reverse();
        let mut ret = vec![0_u32; n1.len() + n2.len()];

        for (i, &x) in n2.iter().enumerate() {
            for (j, &y) in n1.iter().enumerate() {
                ret[i + j] += ((x - b'0') as u32) * ((y - b'0') as u32);
            }
        }

        for i in 0..ret.len() - 1 {
            ret[i + 1] += ret[i] / 10;
            ret[i] %= 10;
        }

        while ret.last() == Some(&0) {
            ret.pop();
        }

        ret.iter().rev().map(|&d| (b'0' + d as u8) as char).collect()
    }
}