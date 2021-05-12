package ddl

import (
	"fmt"
	"github.com/kingzcheung/gexport/ddl/parser"
)

type DdlListener struct {
	*parser.BaseDDLGrammarListener
}

func (d *DdlListener) ExitColumnsDefinition(ctx *parser.ColumnsDefinitionContext) {
	fmt.Println(ctx.GetText())
}

func (d *DdlListener) ExitExitCreateTable(ctx *parser.CreateTableContext) {
	//fmt.Println(ctx.K_CREATE().GetText())
}
