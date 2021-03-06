package metrics

import (
	"errors"
	"fmt"

	"github.com/oslokommune/okctl-metrics-service/pkg/endpoints/metrics/types"

	"github.com/prometheus/client_golang/prometheus"
)

const userAgentKey = "useragent"

type MetricRegistry struct {
	counters map[types.Category]map[types.Action]*prometheus.CounterVec
}

func (m *MetricRegistry) Increment(userAgent string, event Event) error {
	counter, ok := m.counters[event.Category][event.Action]
	if !ok {
		return errors.New("metric not found")
	}

	if event.Labels == nil {
		event.Labels = map[string]string{}
	}

	event.Labels[userAgentKey] = userAgent

	metric, err := counter.GetMetricWith(event.Labels)
	if err != nil {
		return fmt.Errorf("invalid labels: %w", err)
	}

	metric.Inc()

	return nil
}

func (m *MetricRegistry) Add(d types.Definition) {
	for _, action := range d.Actions {
		m.addCounter(d.Category, d.Labels, action)
	}
}

func NewMetricRegistry() *MetricRegistry {
	m := &MetricRegistry{
		counters: make(map[types.Category]map[types.Action]*prometheus.CounterVec),
	}

	return m
}

func (m *MetricRegistry) addCounter(category types.Category, labels []string, action types.Action) {
	if _, ok := m.counters[category]; !ok {
		m.counters[category] = make(map[types.Action]*prometheus.CounterVec)
	}

	m.counters[category][action] = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: defaultMetricNamespace,
		Subsystem: category.String(),
		Name:      action.String(),
	}, append(labels, userAgentKey))

	prometheus.MustRegister(m.counters[category][action])
}

func (m *MetricRegistry) Reset() {
	for category := range m.counters {
		for action := range m.counters[category] {
			prometheus.Unregister(m.counters[category][action])

			delete(m.counters[category], action)
		}

		delete(m.counters, category)
	}
}
