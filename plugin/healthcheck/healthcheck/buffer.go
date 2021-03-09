package healthcheck

import "sync"

// CircularStatus struct
type CircularStatus struct {
	values    []int
	size      int
	valuesMtx sync.RWMutex
}

// NewCircularStatus constructor
func NewCircularStatus(size int) *CircularStatus {
	return &CircularStatus{
		values: []int{},
		size:   size,
	}
}

// Cap return max cap of circular buffer
func (c *CircularStatus) Cap() int {
	return c.size
}

// Size of circular buffer
func (c *CircularStatus) Size() int {
	return len(c.values)
}

// Push status into circular buffer
func (c *CircularStatus) Push(status int) {
	c.valuesMtx.Lock()

	c.values = append(c.values, status)

	length := len(c.values)

	if length > c.size {
		c.values = c.values[length-c.size:]
	}

	c.valuesMtx.Unlock()
}

// Values returns all values of circular buffer
func (c *CircularStatus) Values() []int {
	values := []int{}

	c.valuesMtx.RLock()
	for _, v := range c.values {
		values = append(values, v)
	}
	c.valuesMtx.RUnlock()

	return values
}
