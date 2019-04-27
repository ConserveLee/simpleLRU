package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type LRU struct {
	Len  int
	Map  map[int]*list.Element //哈希表
	link *list.List            //链表
}

type Elem struct {
	//键值对
	key   int
	value interface{}
}

func NewLRU(Len int) *LRU {
	//示例化一个LRU缓存
	return &LRU{
		Len:  Len,
		Map:  map[int]*list.Element{},
		link: list.New(),
	}
}

func (l *LRU) Get(key int) (value interface{}, ok bool) {
	if e, ok := l.Map[key]; ok {
		//元素e插入到链表的头结点
		l.link.MoveToFront(e)
		value = e.Value.(*Elem).value
	}
	return
}

func (l *LRU) Set(key int, value interface{}) {
	if e, ok := l.Map[key]; ok {
		//移动到头结点并更新值
		l.link.MoveToFront(e)
		e.Value.(*Elem).value = value
		return
	}
	v := l.link.PushFront(&Elem{key, value})
	l.Map[key] = v
	if l.link.Len() > l.Len {
		//删除最后一个结点
		l.DeleteLast()
	}
}

func (l *LRU) DeleteLast() {
	e := l.link.Back()
	if e != nil {
		l.removeElem(e)
	}
}

func (l *LRU) delete(key int) {
	if e, ok := l.Map[key]; ok {
		l.removeElem(e)
	}
}

func (l *LRU) removeElem(e *list.Element) {
	l.link.Remove(e)
	v := e.Value.(*Elem)
	delete(l.Map, v.key)
}

func (l *LRU) print() {
	tmp := l.link.Front()
	for tmp != nil {
		fmt.Printf("%v -> ", tmp.Value.(*Elem).value)
		tmp = tmp.Next()
	}
	fmt.Println()
}

var index int

func enter(l *LRU, v interface{}) {
	index++
	l.Set(index, v)
	l.print()
}

func main() {
	l := NewLRU(5)
	fmt.Println("please enter a value")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		t := scanner.Text()
		enter(l, t)
	}

}
