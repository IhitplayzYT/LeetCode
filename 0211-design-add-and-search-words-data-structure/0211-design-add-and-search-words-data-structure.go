type NaryNode struct {
    val      byte
    children map[byte]*NaryNode
    end      bool
}

func (self *NaryNode) add(word string) {
    if word == "" {
        self.end = true
        return
    }
    if self.children == nil {
        self.children = make(map[byte]*NaryNode)
    }
    ch := word[0]
    if _, ok := self.children[ch]; !ok {
        self.children[ch] = &NaryNode{val: ch}
    }
    self.children[ch].add(word[1:])
}

func (self *NaryNode) search(word string) bool {
    if word == "" { return self.end}
    if word[0] == '.' {
        for _, nd := range self.children{
            if nd.search(word[1:]) { return true}
        }
        return false
    }
    nd, ok := self.children[word[0]]
    if !ok { return false }
    return nd.search(word[1:])
}

type WordDictionary struct {
    root *NaryNode
}

func Constructor() WordDictionary {
    return WordDictionary{root: &NaryNode{}}
}

func (self *WordDictionary) AddWord(word string) {
    self.root.add(word)
}

func (self *WordDictionary) Search(word string) bool {
    return self.root.search(word)
}