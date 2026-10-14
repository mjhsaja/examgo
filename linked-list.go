package main

import (
	"errors"
	"fmt"
	"strings"
)

// Node adalah satu elemen di dalam linked list.
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// LinkedList adalah singly linked list generik.
type LinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

var ErrIndexOutOfRange = errors.New("index di luar jangkauan")

// NewLinkedList membuat linked list kosong.
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Len mengembalikan jumlah elemen.
func (l *LinkedList[T]) Len() int { return l.size }

// IsEmpty mengecek apakah list kosong.
func (l *LinkedList[T]) IsEmpty() bool { return l.size == 0 }

// PushFront menambahkan elemen di awal list. O(1)
func (l *LinkedList[T]) PushFront(v T) {
	n := &Node[T]{Value: v, Next: l.head}
	l.head = n
	if l.tail == nil {
		l.tail = n
	}
	l.size++
}

// PushBack menambahkan elemen di akhir list. O(1)
func (l *LinkedList[T]) PushBack(v T) {
	n := &Node[T]{Value: v}
	if l.tail == nil {
		l.head = n
	} else {
		l.tail.Next = n
	}
	l.tail = n
	l.size++
}

// Insert menyisipkan elemen pada posisi index. O(n)
func (l *LinkedList[T]) Insert(index int, v T) error {
	if index < 0 || index > l.size {
		return ErrIndexOutOfRange
	}
	if index == 0 {
		l.PushFront(v)
		return nil
	}
	if index == l.size {
		l.PushBack(v)
		return nil
	}
	prev := l.nodeAt(index - 1)
	prev.Next = &Node[T]{Value: v, Next: prev.Next}
	l.size++
	return nil
}

// PopFront menghapus dan mengembalikan elemen pertama. O(1)
func (l *LinkedList[T]) PopFront() (T, bool) {
	var zero T
	if l.head == nil {
		return zero, false
	}
	v := l.head.Value
	l.head = l.head.Next
	if l.head == nil {
		l.tail = nil
	}
	l.size--
	return v, true
}

// RemoveAt menghapus elemen pada posisi index. O(n)
func (l *LinkedList[T]) RemoveAt(index int) (T, error) {
	var zero T
	if index < 0 || index >= l.size {
		return zero, ErrIndexOutOfRange
	}
	if index == 0 {
		v, _ := l.PopFront()
		return v, nil
	}
	prev := l.nodeAt(index - 1)
	target := prev.Next
	prev.Next = target.Next
	if target == l.tail {
		l.tail = prev
	}
	l.size--
	return target.Value, nil
}

// Remove menghapus kemunculan pertama dari v. O(n)
func (l *LinkedList[T]) Remove(v T) bool {
	var prev *Node[T]
	for cur := l.head; cur != nil; prev, cur = cur, cur.Next {
		if cur.Value != v {
			continue
		}
		if prev == nil {
			l.head = cur.Next
		} else {
			prev.Next = cur.Next
		}
		if cur == l.tail {
			l.tail = prev
		}
		l.size--
		return true
	}
	return false
}

// Get mengambil elemen pada posisi index. O(n)
func (l *LinkedList[T]) Get(index int) (T, error) {
	var zero T
	if index < 0 || index >= l.size {
		return zero, ErrIndexOutOfRange
	}
	return l.nodeAt(index).Value, nil
}

// IndexOf mencari posisi kemunculan pertama v, -1 jika tidak ada. O(n)
func (l *LinkedList[T]) IndexOf(v T) int {
	i := 0
	for cur := l.head; cur != nil; cur = cur.Next {
		if cur.Value == v {
			return i
		}
		i++
	}
	return -1
}

// Contains mengecek apakah v ada di dalam list.
func (l *LinkedList[T]) Contains(v T) bool { return l.IndexOf(v) != -1 }

// Reverse membalik urutan list secara in-place. O(n)
func (l *LinkedList[T]) Reverse() {
	var prev *Node[T]
	cur := l.head
	l.tail = l.head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	l.head = prev
}

// ToSlice mengubah list menjadi slice.
func (l *LinkedList[T]) ToSlice() []T {
	out := make([]T, 0, l.size)
	for cur := l.head; cur != nil; cur = cur.Next {
		out = append(out, cur.Value)
	}
	return out
}

// String menampilkan list dalam format: [1 -> 2 -> 3]
func (l *LinkedList[T]) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for cur := l.head; cur != nil; cur = cur.Next {
		fmt.Fprint(&sb, cur.Value)
		if cur.Next != nil {
			sb.WriteString(" -> ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

func (l *LinkedList[T]) nodeAt(index int) *Node[T] {
	cur := l.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur
}

func main() {
	list := NewLinkedList[int]()

	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)
	list.PushFront(5)
	fmt.Println("Setelah push     :", list, "| len =", list.Len())

	_ = list.Insert(2, 15)
	fmt.Println("Insert(2, 15)    :", list)

	v, _ := list.Get(3)
	fmt.Println("Get(3)           :", v)
	fmt.Println("IndexOf(30)      :", list.IndexOf(30))
	fmt.Println("Contains(99)     :", list.Contains(99))

	list.Remove(15)
	fmt.Println("Remove(15)       :", list)

	removed, _ := list.RemoveAt(0)
	fmt.Println("RemoveAt(0)      :", list, "| dihapus =", removed)

	list.Reverse()
	fmt.Println("Reverse()        :", list)

	front, _ := list.PopFront()
	fmt.Println("PopFront()       :", list, "| diambil =", front)

	if _, err := list.Get(10); err != nil {
		fmt.Println("Get(10)          : error ->", err)
	}

	fmt.Println("ToSlice()        :", list.ToSlice())
}
