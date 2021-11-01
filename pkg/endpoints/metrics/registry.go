package metrics

import (
	"errors"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

type MetricRegistry struct {
	counters map[string]map[Category]map[Action]*prometheus.CounterVec
}

func (m *MetricRegistry) Increment(userAgent string, event Event) error {
	counter, ok := m.counters[userAgent][event.Category][event.Action]
	if !ok {
		return errors.New("metric not found")
	}

	metric, err := counter.GetMetricWith(event.Labels)
	if err != nil {
		return fmt.Errorf("invalid labels: %w", err)
	}

	metric.Inc()

	return nil
}

func (m *MetricRegistry) Add(d Definition) {
	for agent := range m.counters {
		m.addCounters(agent, d.Category, d.Labels, d.Actions...)
	}
}

func NewMetricRegistry(legalUserAgents []string) *MetricRegistry {
	m := &MetricRegistry{
		counters: make(map[string]map[Category]map[Action]*prometheus.CounterVec),
	}

	for _, agent := range legalUserAgents {
		m.counters[agent] = make(map[Category]map[Action]*prometheus.CounterVec)
	}

	return m
}

func (m *MetricRegistry) addCounters(userAgent string, category Category, labels []string, actions ...Action) {
	for _, action := range actions {
		m.addCounter(userAgent, category, labels, action)
	}
}

func (m *MetricRegistry) addCounter(userAgent string, category Category, labels []string, action Action) {
	if _, ok := m.counters[userAgent][category]; !ok {
		m.counters[userAgent][category] = make(map[Action]*prometheus.CounterVec)
	}

	m.counters[userAgent][category][action] = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: userAgent,
		Subsystem: category.String(),
		Name:      action.String(),
	}, labels)

	prometheus.MustRegister(m.counters[userAgent][category][action])
}
