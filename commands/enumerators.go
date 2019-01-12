package commands

import (
	"sort"

	"github.com/last-ent/testy-go/classifier"
)

type OrderedMap struct {
	om    map[int]string
	order []int
}

type MapPair struct {
	key   int
	value string
}

func (o *OrderedMap) add(key int, value string) {
	o.om[key] = value
	o.order = append(o.order, key)
}

func (o *OrderedMap) Iterator() []MapPair {
	pairs := []MapPair{}
	for _, key := range o.order {
		pairs = append(pairs, MapPair{key, o.om[key]})
	}
	return pairs
}

func (o *OrderedMap) AddAll(values []string) {
	sort.Strings(values)
	for i, value := range values {
		o.add(i+1, value)
	}
}

func (o *OrderedMap) Get(key int) string {
	return o.om[key]
}

func NewOrderedMap() *OrderedMap {
	return &OrderedMap{
		om:    map[int]string{},
		order: []int{},
	}
}

// ListDirsAsMap converts TestableDirMeta.Dirs into a map
func ListDirsAsMap(meta *classifier.TestableDirMeta) *OrderedMap {
	dirMap := NewOrderedMap()
	dirs := []string{}
	for _, dir := range meta.Dirs {
		dirs = append(dirs, prefixDir(meta.ParentDir, dir))
	}

	dirMap.AddAll(dirs)

	return dirMap
}

// ListFilesAsMap converts TestableDirMeta.Files into a map
func ListFilesAsMap(meta *classifier.TestableDirMeta) *OrderedMap {
	fileMap := NewOrderedMap()
	files := []string{}
	for _, file := range meta.Files {
		files = append(files, prefixDir(meta.ParentDir, file))
	}
	fileMap.AddAll(files)
	return fileMap
}
