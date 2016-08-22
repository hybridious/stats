package apexstats

import (
	"github.com/apex/log"
	"github.com/segmentio/stats"
)

func NewBackend(logger log.Interface) stats.Backend {
	return stats.BackendFunc(func(e stats.Event) {
		logger.WithFields(fields(e)).Info(e.Name)
	})
}

func fields(e stats.Event) log.Fields {
	return log.Fields{
		"metric": log.Fields{
			"name":   e.Name,
			"type":   e.Type,
			"value":  e.Value,
			"sample": e.Sample,
			"tags":   tags(e.Tags),
			"time":   e.Time,
		},
	}
}

func tags(tags stats.Tags) log.Fields {
	fields := make(log.Fields, len(tags))

	for _, tag := range tags {
		fields[tag.Name] = tag.Value
	}

	return fields
}