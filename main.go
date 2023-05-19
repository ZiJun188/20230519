package main

import (
	"fmt"
	"time"
)

type PriceData struct {
	Timestamp int64
	Price     float64
}

type MovingAverageTracker struct {
	TimeWindow     time.Duration
	PriceQueue     []PriceData
	CurrentSum     float64
	CurrentAverage float64
}

func NewMovingAverageTracker(timeWindow time.Duration) *MovingAverageTracker {
	return &MovingAverageTracker{
		TimeWindow: timeWindow,
		PriceQueue: []PriceData{},
	}
}

func (t *MovingAverageTracker) EventHandler(priceData PriceData) {
	t.PriceQueue = append(t.PriceQueue, priceData)
	t.CurrentSum += priceData.Price

	currentTime := time.Now().UnixNano() / int64(time.Millisecond)
	windowStart := currentTime - int64(t.TimeWindow)

	for len(t.PriceQueue) > 0 && t.PriceQueue[0].Timestamp < windowStart {
		t.CurrentSum -= t.PriceQueue[0].Price
		t.PriceQueue = t.PriceQueue[1:]
	}

	if len(t.PriceQueue) > 0 {
		newAverage := t.CurrentSum / float64(len(t.PriceQueue))
		if newAverage != t.CurrentAverage {
			t.CurrentAverage = newAverage
			t.PrintCurrentAverage()
		}
	}
}

func (t *MovingAverageTracker) PrintCurrentAverage() {
	fmt.Println("当前的移动平均值:", t.CurrentAverage)
}

func main() {
	tracker := NewMovingAverageTracker(time.Minute)

	priceDataStream := []PriceData{
		{Timestamp: 1682982215035, Price: 100.12},
		{Timestamp: 1682982216528, Price: 100.42},
		{Timestamp: 1682982221121, Price: 99.53},
	}

	for _, priceData := range priceDataStream {
		tracker.EventHandler(priceData)
	}
}
