package listmanager

import (
	"math"
	"sync"
)

type Manager struct {
	list []float64
	mu   sync.RWMutex
}

func New() *Manager {
	return &Manager{list: make([]float64, 0)}
}

func (m *Manager) Add(number float64) []float64 {
	m.mu.Lock()
	defer m.mu.Unlock()
	if len(m.list) == 0 {
		m.list = append(m.list, number)
		return m.copy()
	}
	if m.SignsMatch(number, m.list[0]) {
		m.list = append(m.list, number)
	} else {
		reduce(m, math.Abs(number))
	}
	return m.copy()
}

func (m *Manager) List() []float64 {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.copy()
}

func (m *Manager) copy() []float64 {
	cp := make([]float64, len(m.list))
	copy(cp, m.list)
	return cp
}

func (m *Manager) SignsMatch(a, b float64) bool {
	if a == 0 || b == 0 {
		return true
	}
	return (a > 0 && b > 0) || (a < 0 && b < 0)
}

func reduce(m *Manager, amt float64) {
	newList := make([]float64, 0)
	reduction := amt
	for i, v := range m.list {
		av := math.Abs(v)
		if reduction >= av {
			reduction -= av
			continue
		}
		if reduction > 0 {
			if v > 0 {
				newList = append(newList, v-reduction)
			} else {
				newList = append(newList, v+reduction)
			}
			reduction = 0
			newList = append(newList, m.list[i+1:]...)
			break
		}
		newList = append(newList, v)
	}
	m.list = newList
}
