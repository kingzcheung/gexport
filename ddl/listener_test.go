package ddl

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/kingzcheung/gexport/ddl/parser"
	"io/ioutil"
	"testing"
)

func TestParse(t *testing.T) {
	//is := antlr.NewInputStream(src)
	//lexer := parser.NewPhpArrayLexer(is)
	//stream := antlr.NewCommonTokenStream(lexer, 0)
	//parse := parser.NewPhpArrayParser(stream)
	//var listener = DefaultListener
	//var tree = parse.Php()
	//antlr.ParseTreeWalkerDefault.Walk(&listener, tree)
	b, err := ioutil.ReadFile("/Users/kingzcheung/GolandProjects/gexport/testdata/single.sql")
	if err != nil {
		return
	}
	is := antlr.NewInputStream(string(b))
	lexer := parser.NewDDLGrammarLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parse := parser.NewDDLGrammarParser(stream)
	var listener DdlListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, parse.Ddl())
}
