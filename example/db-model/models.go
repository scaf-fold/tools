package db_model

import (
	"github.com/scaf-fold/tools/pkg/xconverter"
)

func Gen(conf string) {
	// executor will generate db model and query and so on
	xconvert.NewModelConverter(conf).Gen()
}
