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
    if self.children == nil { self.children = make(map[byte] *NaryNode)}
    ch := word[0]
    if _, ok := self.children[ch]; !ok {
        self.children[ch] = &NaryNode{val: ch}
    }
    self.children[ch].add(word[1:])
}

func (self *NaryNode) search(word string) bool {
    if word == "" { return self.end}

    if word[0] == '.' {
        for _, nd := range self.children {
            if nd.search(word[1:]) { return true}
        }
        return false
    }

    nd, ok := self.children[word[0]]
    if !ok { return false}
    return nd.search(word[1:])
}

func (self *NaryNode) starts_with(word string) bool {
    if word == "" { return true}
    nd, ok := self.children[word[0]]
    if !ok { return false}
    return nd.starts_with(word[1:])
}

type Trie struct {
    root * NaryNode
}


func Constructor() Trie {
    return Trie{root:&NaryNode{}}
}


func (this *Trie) Insert(word string)  {
    this.root.add(word)
}


func (this *Trie) Search(word string) bool {
    return this.root.search(word)
}


func (this *Trie) StartsWith(prefix string) bool {
    return this.root.starts_with(prefix)
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */