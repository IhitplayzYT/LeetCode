func checkValidString(s string) bool {
    l, h := 0, 0
    for _, c := range s {
        if c == '(' {
            l++;h++
        } else if c == ')' {
            l--;h--
        } else {
            l--;h++
        }
        if h < 0 { return false }
        if l < 0 { l = 0 }
    }
    return l == 0
}