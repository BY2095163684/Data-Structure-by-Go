package LinearList

import "fmt"

const MaxSize = 10

//顺序栈
type SqStack struct {
	data [MaxSize]int //静态数组存放数据
	top  int          //记录栈顶
}

//链栈的实现与单链表无异,链栈的进栈出栈限制在链表的头结点

//初始化
func InitStack(stack *SqStack) {
	for i := 0; i < MaxSize; i++ {
		stack.data[i] = 0
	}
	stack.top = -1
}

//进栈
func Push(stack *SqStack, x int) bool {
	//栈满判断
	if stack.top == MaxSize-1 {
		fmt.Println("栈已满")
		return false
	}
	stack.top++               //栈顶变更
	stack.data[stack.top] = x //数据入栈
	return true
}

//出栈
func Pop(stack *SqStack, x *int) bool {
	//栈空判断
	if stack.top == -1 {
		fmt.Println("栈空")
		return false
	}
	*x = stack.data[stack.top] //数据出栈
	stack.data[stack.top] = 0
	stack.top-- //栈顶变更
	fmt.Printf("%v 出栈\n", *x)
	return true
}

/*
func main() {
	var stack SqStack
	InitStack(&stack)
	fmt.Println(stack)

	for i := 0; i < 20; i++ {
		full := Push(&stack, i)
		if !full {
			break
		}
	}
	fmt.Println(stack)

	x := 0
	for i := 0; i < 20; i++ {
		full := Pop(&stack, &x)
		if !full {
			break
		}
	}
	fmt.Println(stack)
}
*/
