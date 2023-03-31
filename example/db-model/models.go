package db_model

import (
	"github.com/scaf-fold/tools/pkg/gormer"
)

func Gen(conf string) {
	// executor will generate db model and query and so on
	gormer.NewModelConverter(conf).Gen()
}
