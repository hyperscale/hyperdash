package apijson

import "time"

// Config struct
type Config struct {
	URL      string            `hcl:"url"`
	Timeout  int               `hcl:"timeout,optional"`
	Interval int               `hcl:"interval,optional"`
	Query    string            `hcl:"query"`
	Headers  map[string]string `hcl:"headers,optional"`
}

// TimeoutDuration return timeout in second
func (c Config) TimeoutDuration() time.Duration {
	return time.Duration(c.Timeout) * time.Second
}

// IntervalDuration return interval in second
func (c Config) IntervalDuration() time.Duration {
	return time.Duration(c.Interval) * time.Second
}
