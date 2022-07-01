package stackqueueandheap

//两个队列实现栈
import "container/list"

type CQueue struct {
	stack1, stack2 *list.List
}

func NewCQueue() CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (cqueue *CQueue) AppendTail(value int) {
	cqueue.stack1.PushBack(value)
}

func (cqueue *CQueue) DeleteHead() int {
	//如果第二个栈为空
	if cqueue.stack2.Len() == 0 {
		for cqueue.stack1.Len() > 0 {
			cqueue.stack2.PushBack(cqueue.stack1.Remove(cqueue.stack1.Back()))
		}
	}
	if cqueue.stack2.Len() != 0 {
		return cqueue.stack2.Remove(cqueue.stack2.Back()).(int)
	}
	return -1

}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
