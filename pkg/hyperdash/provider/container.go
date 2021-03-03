package provider

import (
	"fmt"
	"os"
	"plugin"
	"strings"
)

// Factory type
type Factory func() Provider

// Container struct
type Container struct {
	providers map[string]Factory
}

// NewContainer constructor
func NewContainer() *Container {
	return &Container{
		providers: make(map[string]Factory),
	}
}

// LoadProviderDir to container
func (c *Container) LoadProviderDir(path string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("os read dir failed: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filename := strings.TrimRight(path, "/") + "/" + entry.Name()

		if err := c.LoadProvider(filename); err != nil {
			return err
		}
	}

	return nil
}

// LoadProvider to container
func (c *Container) LoadProvider(filename string) error {
	p, err := plugin.Open(filename)
	if err != nil {
		return fmt.Errorf("plugin open failed: %w", err)
	}

	v, err := p.Lookup("Provider")
	if err != nil {
		return fmt.Errorf("plugin lookup failed: %w", err)
	}

	provider := v.(func() Provider)

	return c.AddProvider(provider)
}

// AddProvider to container
func (c *Container) AddProvider(factory Factory) error {
	name := factory().Info().Name

	if _, ok := c.providers[name]; ok {
		return fmt.Errorf("provider %s already loaded", name)
	}

	c.providers[name] = factory

	return nil
}

// Get provider by name
func (c *Container) Get(name string) (Provider, error) {
	provider, ok := c.providers[name]
	if !ok {
		return nil, fmt.Errorf("provider %s is not exists", name)
	}

	return provider(), nil
}
