package Tree

import "fmt"

type ElemType3 int

type TreeNode3 struct {
	data           ElemType3
	lChild, rChild *TreeNode3
}

// 初始化根结点
func initRoot2(root *TreeNode3, x ElemType3) {
	root.data = x
	root.lChild = nil
	root.rChild = nil
}

// 创建二叉树结点
func createTreeNode2(data ElemType3) *TreeNode3 {
	newTreeNode3 := new(TreeNode3)
	newTreeNode3.data = data
	newTreeNode3.lChild = nil
	newTreeNode3.rChild = nil
	return newTreeNode3
}

// 队列结点
type linkNode struct {
	data *TreeNode3 // 存指针花销小一点
	next *linkNode
}

// 链式队列
type linkQueue struct {
	front *linkNode
	rear  *linkNode
}

// 队列初始化
func initQueue(Q *linkQueue) {
	Q.front = new(linkNode)
	Q.rear = Q.front
	Q.front.data = nil
	Q.front.next = nil
}

// rear入队
func enQueue(Q *linkQueue, t *TreeNode3) {
	newNode := new(linkNode)
	newNode.data = t
	newNode.next = nil
	Q.rear.next = newNode //新结点插入rear之后
	Q.rear = newNode      //表尾指针后移
}

// front出队
func deQueue(Q *linkQueue) (*TreeNode3, bool) {
	if Q.front == Q.rear { //判空
		return nil, false
	}
	p := Q.front.next     //要出队的是头结点的后一个结点
	x := p.data           //获取出队的数据
	Q.front.next = p.next //出队对应结点后连接剩余结点
	if Q.rear == p {
		Q.rear = Q.front //判断出队结点是否是最后一个结点
	}
	p.data = nil
	p.next = nil
	return x, true
}

/*
层次遍历:
1. 初始化一个辅助队列
2. 根结点入队
3. 若队列非空,则队头结点出队,访问该结点,并将其左、右孩子插入队尾(如果有的话)
4. 重复(3.)直至队列为空
*/

func levelOrder(root *TreeNode3, Q *linkQueue) {
	if root == nil {
		return
	}
	enQueue(Q, root)
	for Q.front != Q.rear {
		treeN, _ := deQueue(Q)
		if treeN != nil {
			fmt.Printf("deQueue: %v\n", treeN)
			if root.lChild != nil {
				enQueue(Q, treeN.lChild)
			}
			if root.rChild != nil {
				enQueue(Q, treeN.rChild)
			}
		}
	}
}

/*
func main() {
	var root TreeNode3
	initRoot2(&root, 1)
	var lQ linkQueue
	initQueue(&lQ)

	newTreeNode2 := createTreeNode2(2)
	newTreeNode3 := createTreeNode2(3)
	newTreeNode4 := createTreeNode2(4)
	newTreeNode5 := createTreeNode2(5)
	newTreeNode6 := createTreeNode2(6)

	root.lChild = newTreeNode2
	root.rChild = newTreeNode3
	newTreeNode2.lChild = newTreeNode4
	newTreeNode2.rChild = newTreeNode5
	newTreeNode3.lChild = newTreeNode6

	levelOrder(&root, &lQ)
}
*/
