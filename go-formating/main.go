package main

import (
	"fmt"
	"time"
)

type ConfigItem struct {
	Key           string
	Value         any
	PreviousValue any
	Version       int
	IsSet         bool
	UpdatedAt     time.Time
}

func (c ConfigItem) String() string {
	return fmt.Sprintf(
		"Key: %s, Value: %v, PreviousValue: %v, Version: %d, IsSet: %t, UpdatedAt: %s",
		c.Key,
		c.Value,
		c.PreviousValue,
		c.Version,
		c.IsSet,
		c.UpdatedAt.Format(time.RFC3339),
	)
}

func (c *ConfigItem) UpdateValue(newValue any) {
	c.PreviousValue = c.Value
	c.Value = newValue
	c.Version++
	c.UpdatedAt = time.Now()
}

func main() {
	item := ConfigItem{
		Key:       "app_name",
		Value:     "MyApp",
		Version:   1,
		IsSet:     true,
		UpdatedAt: time.Now(),
	}

	fmt.Println("Initial Config:")
	fmt.Println(item)

	item.UpdateValue("MyAwesomeApp")

	fmt.Println("\nAfter Update:")
	fmt.Println(item)

	item.UpdateValue("SuperApp")

	fmt.Println("\nAfter Second Update:")
	fmt.Println(item)
}