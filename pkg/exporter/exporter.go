package exporter

import (
	"github.com/Trois-Six/fauna-exporter/pkg/fauna"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog/log"
)

const (
	namespace                   = "fauna"
	transactionalReadOpsCost    = 0.5 / 1000000
	transactionalWriteOpsCost   = 2.5 / 1000000
	transactionalComputeOpsCost = 2.25 / 1000000
	dataStorageCost             = 0.25 / 1024 / 1024 / 1024
)

var (
	billingStartPeriod = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "billing", "start_period"),
		"Billing information - Start Period.",
		[]string{"date"}, nil,
	)
	billingEndPeriod = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "billing", "end_period"),
		"Billing information - End Period.",
		[]string{"date"}, nil,
	)
	billingTotalAmount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "billing", "total_amount"),
		"Billing information - Total amount.",
		[]string{"type"}, nil,
	)
	billingMetricAmountByteReadOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "billing", "metric_amount_byte_read_ops"),
		"Billing information - Metric amount byte read operations.",
		[]string{"type"}, nil,
	)
	billingMetricAmountByteWriteOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "billing", "metric_amount_byte_write_ops"),
		"Billing information - Metric amount byte write operations.",
		[]string{"type"}, nil,
	)
	billingMetricAmountComputeOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "billing", "metric_amount_compute_ops"),
		"Billing information - Metric amount compute operations.",
		[]string{"type"}, nil,
	)
	billingMetricAmountStorage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "billing", "metric_amount_storage"),
		"Billing information - Metric amount storage.",
		[]string{"type"}, nil,
	)
	billingMetricUsageByteReadOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "billing", "metric_usage_byte_read_ops"),
		"Billing information - Metric usage byte read operations.",
		[]string{"type"}, nil,
	)
	billingMetricUsageByteWriteOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "billing", "metric_usage_byte_write_ops"),
		"Billing information - Metric usage byte write operations.",
		[]string{"type"}, nil,
	)
	billingMetricUsageComputeOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "billing", "metric_usage_compute_ops"),
		"Billing information - Metric usage compute operations.",
		[]string{"type"}, nil,
	)
	billingMetricUsageStorage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "billing", "metric_usage_storage"),
		"Billing information - Metric usage storage.",
		[]string{"type"}, nil,
	)
	billingTROCost = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "billing", "transactional_read_ops_cost"),
		"Billing information - Transactional Read Ops cost.",
		[]string{"type"}, nil,
	)
	billingTWROCost = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "billing", "transactional_write_ops_cost"),
		"Billing information - Transactional Write Ops cost.",
		[]string{"type"}, nil,
	)
	billingTCOCost = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "billing", "transactional_compute_ops_cost"),
		"Billing information - Transactional Compute Ops cost.",
		[]string{"type"}, nil,
	)
	billingDataStorageCost = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "billing", "data_storage_cost"),
		"Billing information - Data Storage cost.",
		[]string{"type"}, nil,
	)
	usageByteReadOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "usage", "byte_read_ops"),
		"Usage information - byte read operations.",
		[]string{"collection"}, nil,
	)
	usageByteWriteOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "usage", "byte_write_ops"),
		"Usage information - byte write operations.",
		[]string{"collection"}, nil,
	)
	usageComputeOps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "usage", "compute_ops"),
		"Usage information - compute operations.",
		[]string{"collection"}, nil,
	)
	usageStorage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "usage", "storage"),
		"Usage information - storage.",
		[]string{"collection"}, nil,
	)
	usageVersions = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "usage", "versions"),
		"Usage information - versions.",
		[]string{"collection"}, nil,
	)
	usageIndexes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "usage", "indexes"),
		"Usage information - indexes.",
		[]string{"collection"}, nil,
	)
)

// Exporter collects Fauna metrics and exports them using the prometheus metrics package.
type Exporter struct {
	days   int
	client *fauna.Client
}

// NewExporter returns an initialized Exporter.
func NewExporter(days int, email, password string) *Exporter {
	return &Exporter{
		days:   days,
		client: fauna.NewFaunaClient(email, password),
	}
}

// Describe describes all the metrics ever exported by the Fauna exporter. It
// implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- billingStartPeriod
	ch <- billingEndPeriod
	ch <- billingTotalAmount
	ch <- billingMetricAmountByteReadOps
	ch <- billingMetricAmountByteWriteOps
	ch <- billingMetricAmountComputeOps
	ch <- billingMetricAmountStorage
	ch <- billingMetricUsageByteReadOps
	ch <- billingMetricUsageByteWriteOps
	ch <- billingMetricUsageComputeOps
	ch <- billingMetricUsageStorage
	ch <- billingTROCost
	ch <- billingTWROCost
	ch <- billingTCOCost
	ch <- billingDataStorageCost
	ch <- usageByteReadOps
	ch <- usageByteWriteOps
	ch <- usageComputeOps
	ch <- usageStorage
	ch <- usageVersions
	ch <- usageIndexes
}

func pushBillingMetrics(ch chan<- prometheus.Metric, billing fauna.Billing) {
	ch <- prometheus.MustNewConstMetric(billingStartPeriod, prometheus.GaugeValue, 1, billing.StartPeriod)
	ch <- prometheus.MustNewConstMetric(billingEndPeriod, prometheus.GaugeValue, 1, billing.EndPeriod)
	ch <- prometheus.MustNewConstMetric(billingTotalAmount, prometheus.GaugeValue, float64(billing.TotalAmount)/100, "dollars")
	ch <- prometheus.MustNewConstMetric(billingMetricAmountByteReadOps, prometheus.GaugeValue, float64(billing.MetricAmount.ByteReadOps), "byte_read_ops")
	ch <- prometheus.MustNewConstMetric(billingMetricAmountByteWriteOps, prometheus.GaugeValue, float64(billing.MetricAmount.ByteWriteOps), "byte_write_ops")
	ch <- prometheus.MustNewConstMetric(billingMetricAmountComputeOps, prometheus.GaugeValue, float64(billing.MetricAmount.ComputeOps), "compute_ops")
	ch <- prometheus.MustNewConstMetric(billingMetricAmountStorage, prometheus.GaugeValue, float64(billing.MetricAmount.Storage), "storage")
	ch <- prometheus.MustNewConstMetric(billingMetricUsageByteReadOps, prometheus.GaugeValue, float64(billing.MetricUsage.ByteReadOps), "byte_read_ops")
	ch <- prometheus.MustNewConstMetric(billingMetricUsageByteWriteOps, prometheus.GaugeValue, float64(billing.MetricUsage.ByteWriteOps), "byte_write_ops")
	ch <- prometheus.MustNewConstMetric(billingMetricUsageComputeOps, prometheus.GaugeValue, float64(billing.MetricUsage.ComputeOps), "compute_ops")
	ch <- prometheus.MustNewConstMetric(billingMetricUsageStorage, prometheus.GaugeValue, float64(billing.MetricUsage.Storage), "storage")
	ch <- prometheus.MustNewConstMetric(billingTROCost, prometheus.GaugeValue, float64(billing.MetricUsage.ByteReadOps)*transactionalReadOpsCost, "dollars")
	ch <- prometheus.MustNewConstMetric(billingTWROCost, prometheus.GaugeValue, float64(billing.MetricUsage.ByteWriteOps)*transactionalWriteOpsCost, "dollars")
	ch <- prometheus.MustNewConstMetric(billingTCOCost, prometheus.GaugeValue, float64(billing.MetricUsage.ComputeOps)*transactionalComputeOpsCost, "dollars")
	ch <- prometheus.MustNewConstMetric(billingDataStorageCost, prometheus.GaugeValue, float64(billing.MetricUsage.Storage)*dataStorageCost, "dollars")
}

func pushUsageMetrics(ch chan<- prometheus.Metric, u fauna.UsageType, key string) {
	label := key
	if key == "" {
		label = "all"
	}
	ch <- prometheus.MustNewConstMetric(usageByteReadOps, prometheus.GaugeValue, float64(u.ByteReadOps), label)
	ch <- prometheus.MustNewConstMetric(usageByteWriteOps, prometheus.GaugeValue, float64(u.ByteWriteOps), label)
	ch <- prometheus.MustNewConstMetric(usageComputeOps, prometheus.GaugeValue, float64(u.ComputeOps), label)
	ch <- prometheus.MustNewConstMetric(usageStorage, prometheus.GaugeValue, float64(u.Storage), label)
	ch <- prometheus.MustNewConstMetric(usageVersions, prometheus.GaugeValue, float64(u.Versions), label)
	ch <- prometheus.MustNewConstMetric(usageIndexes, prometheus.GaugeValue, float64(u.Indexes), label)
}

// Collect fetches the metrics and delivers them as Prometheus metrics.
// It implements prometheus.Collector.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	if err := e.client.Login(); err != nil {
		log.Error().Str("exporter", "msg").Msgf("could not login: %s", err)
		return
	}

	billing, err := e.client.GetBillingUsage(e.days)
	if err != nil {
		log.Error().Str("exporter", "msg").Msgf("could not get billing usage: %s", err)
		return
	}

	pushBillingMetrics(ch, billing)

	usage, err := e.client.GetUsage(e.days)
	if err != nil {
		log.Error().Str("exporter", "msg").Msgf("could not get usage: %s", err)
		return
	}

	for k, v := range usage {
		pushUsageMetrics(ch, v, k)
	}
}
