package lib

import (
	"fmt"
)

// Attribute is element of alert
type Attribute struct {
	Type    string   `json:"type"`
	Value   string   `json:"value"`
	Key     string   `json:"key"`
	Context []string `json:"context"`
}

// TimeRange has timestamps of alert begin and end
type TimeRange struct {
	Init float64 `json:"init"`
	Last float64 `json:"last"`
}

// Alert is extranted data from KinesisStream
type Alert struct {
	Name        string `json:"name"`
	Rule        string `json:"rule"`
	Key         string `json:"key"`
	Description string `json:"description"`

	Timestamp TimeRange   `json:"timestamp"`
	Attrs     []Attribute `json:"attrs"`
}

// Title returns string for Github issue title
func (x *Alert) Title() string {
	return fmt.Sprintf("%s: %s", x.Name, x.Description)
}

// AddAttribute just appends the attribute to the Alert
func (x *Alert) AddAttribute(attr Attribute) {
	x.Attrs = append(x.Attrs, attr)
}

// AddAttributes appends set of attribute to the Alert
func (x *Alert) AddAttributes(attrs []Attribute) {
	x.Attrs = append(x.Attrs, attrs...)
}

// FindAttributes searches and returns matched attributes
func (x *Alert) FindAttributes(key string) []Attribute {
	var attrs []Attribute
	for _, attr := range x.Attrs {
		if attr.Key == key {
			attrs = append(attrs, attr)
		}
	}

	return attrs
}

// Match checks attribute type and context.
func (x *Attribute) Match(context, attrType string) bool {
	if x.Type != attrType {
		return false
	}

	for _, ctx := range x.Context {
		if ctx == context {
			return true
		}
	}

	return false
}
