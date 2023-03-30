package db_model

import xconvert "github.com/scaf-fold/tools/xconverter"

func Gen(conf string) {
	// executor will generate db model and query and so on
	xconvert.NewModelConverter(conf).Gen()
}
