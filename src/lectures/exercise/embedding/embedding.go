//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import "fmt"

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func (b *BandwidthUsage) averageBandwidth() int {
	sum := 0
	for i := 0; i < len(b.amount); i++ {
		sum += int(b.amount[i])
	}
	return sum / len(b.amount)
}

func (t *CpuTemp) averageTemp() int {
	sum := 0
	for i := 0; i < len(t.temp); i++ {
		sum += int(t.temp[i])
	}
	return sum / len(t.temp)
}

func (m *MemoryUsage) averageMemory() int {
	sum := 0
	for i := 0; i < len(m.amount); i++ {
		sum += int(m.amount[i])
	}
	return sum / len(m.amount)
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dash := Dashboard{bandwidth, temp, memory}
	fmt.Println("band with average :", dash.averageBandwidth())
	fmt.Println("temp with average :", dash.averageTemp())
	fmt.Println("memory with average :", dash.averageMemory())
}
