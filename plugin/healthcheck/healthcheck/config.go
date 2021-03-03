package healthcheck

import "time"

// Config struct
type Config struct {
	URL                string `hcl:"url"`
	Timeout            int    `hcl:"timeout,optional"`
	Interval           int    `hcl:"interval,optional"`
	UnhealthyThreshold int    `hcl:"unhealthy_threshold,optional"`
	HealthyThreshold   int    `hcl:"healthy_threshold,optional"`
}

// TimeoutDuration return timeout in second
func (c Config) TimeoutDuration() time.Duration {
	return time.Duration(c.Timeout) * time.Second
}

// IntervalDuration return interval in second
func (c Config) IntervalDuration() time.Duration {
	return time.Duration(c.Interval) * time.Second
}
