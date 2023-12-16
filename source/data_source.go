package source

import "github.com/kirillbdev/metronig/model"

type DataSource interface {
	Store(metrics *model.Metrics) error
}
