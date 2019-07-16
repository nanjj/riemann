package main

import (
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	arithCmd := &cobra.Command{
		Use:     "arithmetic",
		Aliases: []string{"a", "arith", "am"},
		Args:    cobra.NoArgs,
		RunE:    ArithmeticE,
	}
	aFlags := arithCmd.Flags()
	aFlags.Int("degree", 50, "degree of difficulty, 0 easiest, 100 most difficult, default 50")
	aFlags.Int("pages", 1, "number of pages to print out, default 10")
	aFlags.Int("maximum", 100, "maximum of number, default 100")
	RootCmd.AddCommand(arithCmd)
}

func ArithmeticE(cmd *cobra.Command, args []string) (err error) {
	level, err := cmd.Flags().GetInt("degree")
	if err != nil {
		return
	}
	level = level % 101
	level = 100 - level
	pages, err := cmd.Flags().GetInt("pages")
	if err != nil {
		return
	}
	if pages <= 0 {
		pages = 10
	}
	maximum, err := cmd.Flags().GetInt("maximum")
	if err != nil {
		return
	}
	if maximum <= 0 {
		maximum = 100
	}
	pdf := NewFpdf("noto", "mono", "times", 16)
	if maximum == 100 {
		WritePages100(pdf, pages, maximum, level)
	}
	f, err := ioutil.TempFile(os.TempDir(), "e100*.pdf")
	if err != nil {
		panic(err)
	}
	fp := f.Name()
	defer os.Remove(fp)
	err = pdf.OutputFileAndClose(fp)
	if err != nil {
		panic(err)
	}
	err = Open(fp)
	return
}
