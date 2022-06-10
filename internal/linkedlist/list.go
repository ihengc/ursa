package linkedlist

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/7 20:05
 * @description: double linked list

可以使用泛型的双向链表
 ***************************************************************/

// Element 表示链表中的节点
type Element[T any] struct {
	next, prev *Element[T] // 当前节点的前后节点指针
	list       *List[T]    // 双向链表指针，用于确定当前节点属于哪个双向链表
	Value      T           // 存放具体的数据
}

// Next 返回当前节点的下一个节点指针
// 双向链表的root节点为空节点(哨兵节点)
func (e *Element[T]) Next() *Element[T] {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// Prev 返回当前阶段的前一个节点指针
func (e *Element[T]) Prev() *Element[T] {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// List 双向链表
type List[T any] struct {
	root Element[T] // 根节点(哨兵节点)
	len  int        // 双向链表长度
}

// Init 初始化链表
func (l *List[T]) Init() *List[T] {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// New 创建链表
func New[T any]() *List[T] { return new(List[T]).Init() }

// Len 返回链表的长度
func (l *List[T]) Len() int { return l.len }

// Front 返回链表的头节点
func (l *List[T]) Front() *Element[T] {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Back 返回链表的尾节点
func (l *List[T]) Back() *Element[T] {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// lazyInit 延迟初始链表;在使用时初始化
func (l *List[T]) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// insert 在at节点后面插入e节点,并返回插入的节点e指针
func (l *List[T]) insert(e, at *Element[T]) *Element[T] {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

func (l *List[T]) insertValue(v T, at *Element[T]) *Element[T] {
	return l.insert(&Element[T]{Value: v}, at)
}

func (l *List[T]) remove(e *Element[T]) {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	e.list = nil
	l.len--
}

func (l *List[T]) move(e, at *Element[T]) {
	if e == at {
		return
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
}

func (l *List[T]) Remove(e *Element[T]) T {
	if e.list == l {
		l.remove(e)
	}
	return e.Value
}

func (l *List[T]) PutFront(v T) *Element[T] {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

func (l *List[T]) PushBack(v T) *Element[T] {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

func (l *List[T]) InsertBefore(v T, mark *Element[T]) *Element[T] {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.prev)
}

func (l *List[T]) InsertAfter(v T, mark *Element[T]) *Element[T] {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark)
}

func (l *List[T]) MoveToFront(e *Element[T]) {
	if e.list != l || l.root.next == e {
		return
	}
	l.move(e, &l.root)
}

func (l *List[T]) MoveToBack(e *Element[T]) {
	if e.list != l || l.root.prev == e {
		return
	}
	l.move(e, l.root.prev)
}

func (l *List[T]) MoveBefore(e, mark *Element[T]) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark.prev)
}

func (l *List[T]) MoveAfter(e, mark *Element[T]) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark)
}

func (l *List[T]) PushBackList(other *List[T]) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

func (l *List[T]) PushFrontList(other *List[T]) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}

func (l *List[T]) PopFront() *Element[T] {
	if l.len == 0 {
		return nil
	}
	e := l.root.next
	l.remove(e)
	return e
}
