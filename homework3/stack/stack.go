package stack

import "fmt"

type stack struct {
	top  *element
	size int64
}

type element struct {
	data interface{}
	next *element
}

func New() *stack {
	return &stack{}
}

func (st *stack) Push(data interface{}) {
	st.top = &element{data, st.top}
	st.size++
}

func (st *stack) Pop() (data interface{}) {
	if st.size > 0 {
		data, st.top = st.top.data, st.top.next
		st.size--
		return
	}
	return nil
}

func (st *stack) Top() interface{} {
	return st.top.data
}

func (st *stack) PrintAll() {
	var el *element
	el = st.top
	for el != nil {
		fmt.Println(el.data)
		el = el.next
	}
}

func (st *stack) GetSize() int64 {
	return st.size
}

func (st *stack) IsEmpty() bool {
	return st.size == 0
}
