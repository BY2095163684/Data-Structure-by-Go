package Tree

import "fmt"

// 树采用孩子-兄弟表示法化为二叉树,即可使用二叉树的方法来操作树
// 左孩子右兄弟
// 森林也可用这种方法转化为二叉树

type CSNode struct {
	data                 int
	firstChild, rightSib *CSNode // 第一个孩子和右兄弟指针
}

// 创建新结点
func createNode(data int) *CSNode {
	return &CSNode{data: data}
}

// 先序遍历
func newpreOrder(node *CSNode) {
	if node != nil {
		fmt.Printf("%v ", node.data)
		newpreOrder(node.firstChild)
		newpreOrder(node.rightSib)
	}
}

/*
     1
   / | \
  2  3  4
 / \
5   6

			1
		   /
		  2
-->		 / \
		5   3
		\   \
		 6   4
*/

/*
func main() {
	// 创建树节点
	root := createNode(1)
	root.firstChild = createNode(2)
	root.firstChild.rightSib = createNode(3)
	root.firstChild.rightSib.rightSib = createNode(4)
	root.firstChild.firstChild = createNode(5)
	root.firstChild.firstChild.rightSib = createNode(6)

	// 先序遍历
	fmt.Print("Pre-order traversal: ")
	newpreOrder(root)
	fmt.Println()
}
*/
