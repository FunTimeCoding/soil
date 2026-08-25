package collector

import "time"

func (c *Collector) SetScrape(
	hypervisor string,
	success bool,
	duration time.Duration,
) {
	var value float64

	if success {
		value = 1
	}

	c.scrape.success.WithLabelValues(hypervisor).Set(value)
	c.scrape.duration.WithLabelValues(hypervisor).Set(duration.Seconds())
}
