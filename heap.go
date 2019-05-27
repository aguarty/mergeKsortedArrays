package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Heap struct {
	size int
	arr  []*Pointer
}

type Pointer struct {
	LineIndex int
	ValIndex  int
}

func NewHeap2(size int) *Heap {
	arr := make([]*Pointer, 1, size)
	arr[0] = &Pointer{}
	h := &Heap{size: 0, arr: arr}
	return h
}

func (h *Heap) proclateDown(parent int) {
	lChild := 2 * parent
	rChild := lChild + 1
	small := -1
	if lChild <= h.size {
		small = lChild
	}
	if rChild <= h.size && h.comp(lChild, rChild) {
		small = rChild
	}
	if small != -1 && h.comp(parent, small) {
		h.swap(parent, small)
		h.proclateDown(small)
	}
}

func (h *Heap) proclateUp(child int) {
	parent := child / 2
	if parent == 0 {
		return
	}
	if h.comp(parent, child) {
		h.swap(child, parent)
		h.proclateUp(parent)
	}
}
func (h *Heap) comp(i, j int) bool {
	return input[h.arr[i].LineIndex][h.arr[i].ValIndex] > input[h.arr[j].LineIndex][h.arr[j].ValIndex]
}

func (h *Heap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) Empty() bool {
	return (h.size == 0)
}

func (h *Heap) Add(value *Pointer) {
	h.size++
	h.arr = append(h.arr, value)
	h.proclateUp(h.size)
}

func (h *Heap) Remove() (*Pointer, bool) {
	if h.Empty() {
		fmt.Println("HeapEmptyError.")
		return &Pointer{}, false
	}
	value := h.arr[1]
	h.arr[1] = h.arr[h.size]
	h.size--
	h.proclateDown(1)
	h.arr = h.arr[0 : h.size+1]
	return value, true
}

var (
	countlines int
	countempty int
	input      [][]int
)

func main() {

	inputfile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inputfile.Close()

	outputfile, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer outputfile.Close()
	w := bufio.NewWriter(outputfile)

	scanner := bufio.NewScanner(inputfile)
	_ = scanner.Scan()
	countlines, _ = strconv.Atoi(scanner.Text())

	if countlines == 0 {
		return
	}

	input = make([][]int, countlines, countlines)
	for i := 0; scanner.Scan(); i++ {
		tmpbuf := scanner.Text()
		//fmt.Println(tmpbuf)
		tmparr := strings.Split(tmpbuf, " ")
		input[i] = make([]int, len(tmparr), len(tmparr))
		for idx, item := range tmparr {
			input[i][idx], _ = strconv.Atoi(item)
		}
	}

	hp := NewHeap2(countlines - countempty)

	for i := 0; i < countlines; i++ {
		if len(input[i]) > 1 {
			hp.Add(&Pointer{
				LineIndex: i,
				ValIndex:  1,
			})

		} else {
			countempty++
		}
	}

	for count := 0; count < countlines-countempty; {
		var (
			item *Pointer
			ok   bool
		)
		if item, ok = hp.Remove(); ok {
			//fmt.Print(*item.Value, " ")
			_, err := w.WriteString(strconv.Itoa(input[item.LineIndex][item.ValIndex]))
			if err != nil {
				panic(err)
			}
			_, err = w.WriteString(" ")
			if err != nil {
				panic(err)
			}
		}
		if item.ValIndex == len(input[item.LineIndex])-1 {
			count++
		} else {
			item.ValIndex++
			hp.Add(item)
		}
	}
	w.Flush()
}
