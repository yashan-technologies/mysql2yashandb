package handler

import (
	"m2y/db"
	"m2y/defs/confdef"
	"m2y/internal/modules"
)

type CheckDataHandler struct {
	parallel   int
	sampleLine int
}

func NewCheckDataHandler(parallel, sampleLine int) *CheckDataHandler {
	return &CheckDataHandler{
		parallel:   parallel,
		sampleLine: sampleLine,
	}
}

func (c *CheckDataHandler) CheckData() error {
	conf := confdef.GetM2YConfig()
	var res [][]string
	var err error
	if len(conf.MySQL.Tables) != 0 {
		res, err = modules.CompareTables(db.MySQLDB, db.YashanDB, conf.MySQL.Database, conf.Yashan.RemapSchemas[0], conf.MySQL.Tables, c.parallel, c.sampleLine)
	} else {
		res, err = modules.CompareSchemas(db.MySQLDB, db.YashanDB, conf.MySQL.Schemas, conf.Yashan.RemapSchemas, conf.MySQL.ExcludeTables, c.parallel, c.sampleLine)
	}
	if err != nil {
		return err
	}
	modules.PrintCheckResults(res)
	return nil
}
