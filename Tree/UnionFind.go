package Tree

// 并查集思想类似于树的双亲表示法

// UnionFind 结构体
type UnionFind struct {
	parent []int // 切片存储每个结点的父结点
	size   []int // 存储每个集合的大小
}

// 初始化并查集
func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i // 初始化为每个节点的父节点是它自己
		uf.size[i] = 1   // 初始化为每个集合的大小为1
	}
	return uf
}

// 查找根节点并进行路径压缩
func (uf *UnionFind) Find(p int) int { // 查找节点p的根结点
	if uf.parent[p] != p { // 如果p不是根结点
		uf.parent[p] = uf.Find(uf.parent[p]) // 则递归查找其父结点，并在查找过程中进行路径压缩，将p直接连接到根结点
	}
	return uf.parent[p]
}

// 合并两个集合
func (uf *UnionFind) Union(p, q int) { // 合并两个结点p和q所在的集合
	// 找到它们的根结点rootP和rootQ
	rootP := uf.Find(p)
	rootQ := uf.Find(q)
	if rootP != rootQ { // 如果它们不在同一个集合中,根据集合的大小进行合并
		if uf.size[rootP] < uf.size[rootQ] { // 较小的集合连接到较大的集合上
			uf.parent[rootP] = rootQ
			uf.size[rootQ] += uf.size[rootP]
		} else {
			uf.parent[rootQ] = rootP
			uf.size[rootP] += uf.size[rootQ]
		}
	}
}

// 检查两个节点是否连通
func (uf *UnionFind) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

/*
func main() {
	uf := NewUnionFind(10)
	uf.Union(1, 2)
	uf.Union(3, 4)
	uf.Union(2, 4)
	fmt.Println(uf.Connected(1, 3)) // 输出: true
	fmt.Println(uf.Connected(1, 5)) // 输出: false
}
*/
