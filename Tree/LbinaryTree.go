package Tree

import "fmt"

type ElemType2 int

type TreeNode2 struct {
	data           ElemType2
	lChild, rChild *TreeNode2
	// father *TreeNode2	// 三叉链表,方便找父结点
}

// 初始化根结点
func initRoot(root *TreeNode2, x ElemType2) {
	root.data = x
	root.lChild = nil
	root.rChild = nil
}

// 创建二叉树结点
func createTreeNode(data ElemType2) *TreeNode2 {
	newTreeNode2 := new(TreeNode2)
	newTreeNode2.data = data
	newTreeNode2.lChild = nil
	newTreeNode2.rChild = nil
	return newTreeNode2
}

/*
先序遍历: 根、左、右
中序遍历: 左、根、右
后序遍历: 左、右、根
*/

// 先序遍历
func preOrder(root *TreeNode2) {
	if root == nil {
		return
	}
	fmt.Printf("%v", root.data)
	preOrder(root.lChild)
	preOrder(root.rChild)
}

// 中序遍历
func inOrder(root *TreeNode2) {
	if root == nil {
		return
	}
	inOrder(root.lChild)
	fmt.Printf("%v", root.data)
	inOrder(root.rChild)
}

// 后序遍历
func postOrder(root *TreeNode2) {
	if root == nil {
		return
	}
	postOrder(root.lChild)
	postOrder(root.rChild)
	fmt.Printf("%v", root.data)
}

/*
func main() {
	var root TreeNode2
	initRoot(&root, 1)
	fmt.Println(root)

	newTreeNode2 := createTreeNode(2)
	newTreeNode3 := createTreeNode(3)
	newTreeNode4 := createTreeNode(4)
	newTreeNode5 := createTreeNode(5)
	newTreeNode6 := createTreeNode(6)

	root.lChild = newTreeNode2
	root.rChild = newTreeNode3
	newTreeNode2.lChild = newTreeNode4
	newTreeNode2.rChild = newTreeNode5
	newTreeNode3.lChild = newTreeNode6

	preOrder(&root)
	inOrder(&root)
	postOrder(&root)
}
*/
