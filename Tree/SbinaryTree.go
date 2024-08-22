package Tree

import (
	"math"
)

const MaxSize = 10

type ElemType int

// 定义树结点
type TreeNode struct {
	data    ElemType
	isEmpty bool
}

// 初始化树结点
func initTree(T []TreeNode) { // TreeNode切片,动态数组,存储相同类型的元素(TreeNode结构体),并且长度可以动态变化
	for i := 0; i < MaxSize; i++ {
		T[i].data = 0
		T[i].isEmpty = true
	}
}

// 切片（slice），引用类型，本质上是一个包含指向底层数组的指针、长度（len）和容量（cap）的结构体。
// 因此，当将切片作为参数传递给函数时，实际上是按值传递这个结构体的副本，但这个副本仍然指向相同的底层数组。

// 存入数据
func inputData(T []TreeNode, x ElemType) bool {
	for i := 0; i < len(T); i++ {
		if T[i].isEmpty {
			T[i].data = ElemType(x) // 自定义类型要进行转换
			T[i].isEmpty = false
			return true
		}
	}
	return false
}

// 左孩子结点
// 在完全二叉树中, i的左孩子是2i
func lChild(T []TreeNode, i int, e *ElemType) bool {
	if i >= MaxSize {
		return false
	}
	if 2*i < MaxSize && !T[2*i].isEmpty {
		*e = T[2*i].data
		return true
	}
	return false
}

// 右孩子结点
// 在完全二叉树中, i的右孩子是2i+1
func rChild(T []TreeNode, i int, e *ElemType) bool {
	if i >= MaxSize {
		return false
	}
	if 2*i+1 < MaxSize && !T[2*i+1].isEmpty {
		*e = T[2*i+1].data
		return true
	}
	return false
}

// 父结点
// 在完全二叉树中, i的父结点是[i/2]
func father(T []TreeNode, i int, e *ElemType) bool {
	if i >= MaxSize {
		return false
	}
	if i/2 < MaxSize && !T[i/2].isEmpty {
		*e = T[i/2].data
		return true
	}
	return false
}

// 判断i所在层次
func treeNodeFloor(i int) int {
	if i >= MaxSize {
		return -1
	}
	return int(math.Floor(math.Log2(float64(i))) + 1)
}

// 判断i是否为叶子结点
func isLeaf(i int) bool {
	return i > (MaxSize-1)/2
}

// 二叉树的顺序存储只适合完全二叉树

/*
// 除根节点外,任何一个结点有且仅有一个前驱
func main() {
	T := make([]TreeNode, MaxSize) // 分配内存,创造切片
	initTree(T)
	fmt.Println(len(T))
	fmt.Println(T)

	for i := ElemType(0); i < MaxSize; i++ {
		inputData(T, i)
	}
	fmt.Println(T)

	var e ElemType
	if lChild(T, 2, &e) {
		fmt.Println(e)
	}

	if rChild(T, 3, &e) {
		fmt.Println(e)
	}

	if father(T, 3, &e) {
		fmt.Println(e)
	}

	fmt.Println(treeNodeFloor(5))

	fmt.Println(isLeaf(9))
}
*/
