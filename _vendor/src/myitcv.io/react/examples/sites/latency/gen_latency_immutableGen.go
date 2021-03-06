// Code generated by immutableGen. DO NOT EDIT.

package main

//immutableVet:skipFile

import (
	"myitcv.io/immutable"
)

//
// latencies is an immutable type and has the following template:
//
// 	map[location]latency
//
type latencies struct {
	theMap  map[location]latency
	mutable bool
	__tmpl  _Imm_latencies
}

var _ immutable.Immutable = new(latencies)
var _ = new(latencies).__tmpl

func newLatencies(inits ...func(m *latencies)) *latencies {
	res := newLatenciesCap(0)
	if len(inits) == 0 {
		return res
	}

	return res.WithMutable(func(m *latencies) {
		for _, i := range inits {
			i(m)
		}
	})
}

func newLatenciesCap(l int) *latencies {
	return &latencies{
		theMap: make(map[location]latency, l),
	}
}

func (m *latencies) Mutable() bool {
	return m.mutable
}

func (m *latencies) Len() int {
	if m == nil {
		return 0
	}

	return len(m.theMap)
}

func (m *latencies) Get(k location) (latency, bool) {
	v, ok := m.theMap[k]
	return v, ok
}

func (m *latencies) AsMutable() *latencies {
	if m == nil {
		return nil
	}

	if m.Mutable() {
		return m
	}

	res := m.dup()
	res.mutable = true

	return res
}

func (m *latencies) dup() *latencies {
	resMap := make(map[location]latency, len(m.theMap))

	for k := range m.theMap {
		resMap[k] = m.theMap[k]
	}

	res := &latencies{
		theMap: resMap,
	}

	return res
}

func (m *latencies) AsImmutable(v *latencies) *latencies {
	if m == nil {
		return nil
	}

	if v == m {
		return m
	}

	m.mutable = false
	return m
}

func (m *latencies) Range() map[location]latency {
	if m == nil {
		return nil
	}

	return m.theMap
}

func (m *latencies) WithMutable(f func(mi *latencies)) *latencies {
	res := m.AsMutable()
	f(res)
	res = res.AsImmutable(m)

	return res
}

func (m *latencies) WithImmutable(f func(mi *latencies)) *latencies {
	prev := m.mutable
	m.mutable = false
	f(m)
	m.mutable = prev

	return m
}

func (m *latencies) Set(k location, v latency) *latencies {
	if m.mutable {
		m.theMap[k] = v
		return m
	}

	res := m.dup()
	res.theMap[k] = v

	return res
}

func (m *latencies) Del(k location) *latencies {
	if _, ok := m.theMap[k]; !ok {
		return m
	}

	if m.mutable {
		delete(m.theMap, k)
		return m
	}

	res := m.dup()
	delete(res.theMap, k)

	return res
}
func (s *latencies) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
	if s == nil {
		return true
	}

	if s.Mutable() {
		return false
	}
	return true
}
