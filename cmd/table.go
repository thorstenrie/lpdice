package main

import (
	"errors"
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/thorstenrie/lpstats"
)

type Table struct {
	padding int
	header  []string
	rows    map[int]Row
	width   map[string]int
}

type Row map[string]string

func New() (*Table, error) {
	/*if (head == nil) || (len(head) == 0) {
		return nil, errors.New("empty header")
	}
	h := make(header)
	for i, c := range head {
		h[c] = i
	}*/
	t := &Table{
		padding: DefaultPadding,
		header:  make([]string, 0),
		rows:    make(map[int]Row),
	}
	return t, nil
}

func (t *Table) AddColumn(c string) error {
	if printable(c) != c {
		return errors.New("non-printable characters in column name")
	}
	if t.columnExists(c) {
		return errors.New("column already exists")
	}
	t.header = append(t.header, c)
	t.width[c] = utf8.RuneCountInString(c)
	return nil
}

func (t *Table) columnExists(c string) bool {
	for _, h := range t.header {
		if h == c {
			return true
		}
	}
	return false
}

func (t *Table) AddRow(r Row) error {
	if (r == nil) || (len(r) == 0) {
		return errors.New("empty row")
	}
	i := len(t.rows)
	t.rows[i] = make(Row)
	for column, cell := range r {
		if !t.columnExists(column) {
			t.AddColumn(column)
		}
		t.rows[i][column] = cell
		t.width[column] = lpstats.Max(t.width[column], utf8.RuneCountInString(column))
	}
	return nil
}

func (t *Table) Print() string {

	return ""
}

const (
	DefaultPadding = 2
)

type StrMap[T any] map[string]T

type PrintMap map[string]string

func (m *PrintMap) Print(prefix string) (string, error) {
	text := ""
	sm := (*StrMap[string])(m)
	max, _ := sm.MaxKeyLength()
	keys, _ := sm.Sorted()
	for _, k := range keys {
		text += prefix +
			k +
			strings.Repeat(" ", max+1-len(k)) +
			(*m)[k] +
			"\n"
	}
	return text, nil
}

func (m *StrMap[T]) MaxKeyLength() (int, error) {
	if (m == nil) || (len(*m) == 0) {
		return 0, errors.New("empty map")
	}
	max := 0
	for c := range *m {
		max = lpstats.Max(max, utf8.RuneCountInString(c))
	}
	return max, nil
}

func (m *StrMap[T]) Sorted() ([]string, error) {
	if (m == nil) || (len(*m) == 0) {
		return nil, errors.New("empty map")
	}
	keys := make([]string, 0, len(*m))
	for k := range *m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys, nil
}
