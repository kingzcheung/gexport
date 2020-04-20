package cmd

import (
	"bytes"
	"fmt"
	"github.com/kingzcheung/gexport"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path"
)

var name string
var outfile string

var rootCmd = &cobra.Command{
	Use:   "gexport",
	Short: "gexport 是一个转换sql/json到struct的命令",
	Run:   rootRun,
}

func init() {
	rootCmd.Flags().StringVarP(&name, "name", "n", gexport.DefaultStructName, "struct name")
	rootCmd.Flags().StringVarP(&outfile, "outfile", "o", "", "outfile name")
}
func rootRun(cmd *cobra.Command, args []string) {

	var buf bytes.Buffer
	_, err := buf.ReadFrom(os.Stdin)
	if err != nil {
		log.Fatalln(err)
		return
	}

	gx := gexport.New(buf.String(), exportType(buf.Bytes()))
	//设置名称
	gx.StructName = name
	gx.Parse()
	if gx.Error() != nil {
		log.Fatalln("数据解析错误！")
	}

	output := gx.Output()

	for _, g := range output {
		if outfile != "" {
			//写入文件
			if err := writeFile(outfile, g); err != nil {
				log.Fatalln(err)
			}
			continue
		}
		fmt.Println(g)
	}
}

func writeFile(outfile string, g string) error {
	var fs = afero.NewOsFs()
	outDir := path.Dir(outfile)

	_ = fs.MkdirAll(outDir, 0777)

	if err := afero.WriteFile(fs, outfile, []byte(g), 0777); err != nil {
		return err
	}

	return nil
}

func exportType(str []byte) gexport.ExportType {
	if bytes.HasPrefix(str, []byte{123}) {
		return gexport.JSON
	}
	return gexport.SQL
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
