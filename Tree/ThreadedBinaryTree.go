package Tree

import "fmt"

// 含有n个结点的二叉树的二叉链表中，其空的指针域的个数为n+1
// 利用空链域(变成线索)指向结点前驱与后继(具体根据遍历算法)

type ElemType4 int

// 线索二叉树结点
type ThreadNode struct {
	data           ElemType4
	lChild, rChild *ThreadNode
	ltag, rtag     int // 0: 指向孩子, 1: 指向前驱或后继
}

var pre *ThreadNode = nil // 全局变量,记录前驱

// 访问结点函数
func visit(node *ThreadNode) {
	if node.lChild == nil { // 左子树为空
		node.lChild = pre // 建立前驱线索
		node.ltag = 1
	}
	if pre != nil && pre.rChild == nil {
		pre.rChild = node // 建立后继线索
		pre.rtag = 1
	}
	pre = node
}

// 中序线索化(先序、后序方法类似)
func inThread(root *ThreadNode) {
	if root != nil {
		inThread(root.lChild) // 中序遍历左子树
		visit(root)           // 访问根结点
		inThread(root.rChild) // 中序遍历右子树
	} else {
		return
	}
}

// 中序遍历线索二叉树
func inOrderTraverse(root *ThreadNode) {
	node := root
	for node != nil && node.ltag == 0 {
		node = node.lChild
	}
	for node != nil {
		fmt.Printf("%d ", node.data)
		if node.rtag == 1 {
			node = node.rChild
		} else {
			node = node.rChild
			for node != nil && node.ltag == 0 {
				node = node.lChild
			}
		}
	}
}

// 中序线索二叉树寻找中序后继(寻找前驱思路也类似)
// 找到以node为根的子树中第一个被中序遍历的结点
func firstNode(node *ThreadNode) *ThreadNode {
	for node.ltag == 0 { // 循环找到最左下结点(不一定是叶子结点)
		node = node.lChild
	}
	return node
}

// 在中序线索二叉树中找到结点node的后继结点
func nextNode(node *ThreadNode) *ThreadNode {
	if node.rtag == 0 {
		return firstNode(node.rChild) // 右子树最左下结点
	} else {
		return node.rChild // rtag == 1直接返回后继线索
	}
}

// 对中序线索二叉树进行中序遍历(利用线索实现的非递归算法)
func newInOrder(tree *ThreadNode) {
	for node := firstNode(tree); node != nil; node = nextNode(node) {
		fmt.Println(node.data)
	}
}

/*
func main() {

	root := &ThreadNode{data: 1, ltag: 0, rtag: 0}
	root.lChild = &ThreadNode{data: 2, ltag: 0, rtag: 0}
	root.rChild = &ThreadNode{data: 3, ltag: 0, rtag: 0}

	inThread(root)
	// 处理最后一个结点的右孩子为空的问题
	if pre != nil && pre.rChild == nil {
		pre.rtag = 1
		pre.rChild = nil // 或者指向一个特殊的头结点
	}

	inOrderTraverse(root)
	newInOrder(root)
}
*/
