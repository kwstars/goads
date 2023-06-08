package hashmap

import (
	"reflect"
	"testing"
)

func TestNewWithInitialCapacity(t *testing.T) {
	m := New[int, string](WithInitialCapacity[int, string](10))

	if m == nil || m.m == nil || len(m.m) != 0 {
		t.Errorf("Map not created with specified capacity")
	}
}

func TestPutAndGet(t *testing.T) {
	tests := []struct {
		name  string
		key   int
		value string
	}{
		{name: "Insert and retrieve normal value", key: 1, value: "one"},
		{name: "Insert and retrieve empty string", key: 2, value: ""},
		{name: "Insert and retrieve very large string", key: 3, value: string(make([]byte, 1<<20))},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := New[int, string]()
			m.Put(test.key, test.value)
			v, found := m.Get(test.key)

			if !found || v != test.value {
				t.Errorf("Got %v, expected %v", v, test.value)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name        string
		key         int
		initialMap  *Map[int, string]
		expectedMap *Map[int, string]
	}{
		{
			name: "Remove existing key",
			key:  1,
			initialMap: &Map[int, string]{
				m: map[int]string{
					1: "one",
					2: "two",
				},
			},
			expectedMap: &Map[int, string]{
				m: map[int]string{
					2: "two",
				},
			},
		},
		{
			name: "Remove non-existing key",
			key:  3,
			initialMap: &Map[int, string]{
				m: map[int]string{
					1: "one",
					2: "two",
				},
			},
			expectedMap: &Map[int, string]{
				m: map[int]string{
					1: "one",
					2: "two",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.initialMap.Remove(test.key)
			if !reflect.DeepEqual(test.initialMap, test.expectedMap) {
				t.Errorf("Got %v, expected %v", test.initialMap, test.expectedMap)
			}
		})
	}
}

func TestEmptyAndSize(t *testing.T) {
	m := New[int, string]()
	if !m.Empty() || m.Size() != 0 {
		t.Errorf("New map should be empty and have size 0")
	}

	m.Put(1, "one")
	if m.Empty() || m.Size() != 1 {
		t.Errorf("Map should have one element")
	}

	m.Remove(1)
	if !m.Empty() || m.Size() != 0 {
		t.Errorf("Map should be empty and have size 0 after removing the only element")
	}
}

func TestClear(t *testing.T) {
	m := New[int, string]()
	m.Put(1, "one")
	m.Put(2, "two")

	m.Clear()
	if !m.Empty() || m.Size() != 0 {
		t.Errorf("Map should be empty and have size 0 after Clear()")
	}
}
