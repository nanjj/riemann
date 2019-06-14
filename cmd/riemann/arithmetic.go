package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/jung-kurt/gofpdf"
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
	RootCmd.AddCommand(arithCmd)
}

func ArithmeticE(cmd *cobra.Command, args []string) (err error) {
	level, err := cmd.Flags().GetInt("degree")
	if err != nil {
		return
	}
	level = level % 100
	level = 100 - level
	pages, err := cmd.Flags().GetInt("pages")
	if err != nil {
		return
	}
	if pages <= 0 {
		pages = 10
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	fn := "mcust"
	pdf.AddUTF8FontFromBytes(fn, "", FontBytes())
	pdf.SetLeftMargin(12.5)
	pdf.SetTopMargin(7.5)
	carry := func(a, b int) bool {
		a = a % 10
		b = b % 10
		if a+b >= 10 {
			return true
		}
		return false
	}
	borrow := func(a, b int) bool {
		a = a % 10
		b = b % 10
		if a-b < 0 {
			return true
		}
		return false
	}
	options := gofpdf.ImageOptions{ImageType: "png", ReadDpi: true, AllowNegativePosition: false}
	pdf.RegisterImageOptionsReader("riemann", options, bytes.NewBuffer(res["riemann"]))
	pdf.SetHeaderFuncMode(func() {
		pdf.SetFont(fn, "", 16)
		pdf.Cell(44, 10, "姓名：")
		pdf.Cell(44, 10, "日期：")
		pdf.Cell(44, 10, "时间：")
		pdf.Cell(44, 10, "分数：")
		pdf.ImageOptions("riemann", -1, 6, 10, 10, false, options, 0, "")
		width, _ := pdf.GetPageSize()
		pdf.Line(10, 20, width-10, 20)
		pdf.SetFont("Courier", "", 16)
	}, true)
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("times", "I", 8)
		width, _ := pdf.GetPageSize()
		pdf.SetX(width/2 - 5)
		pdf.Cell(10, 10, fmt.Sprintf("%-2d", pdf.PageNo()))
	})
	ops := map[int]string{0: "+", 1: "-"}
	for p := 0; p < pages; p++ {
		pdf.AddPage()
		pdf.Ln(3.5)
		easiers := 0
		for i := 0; i < 100; i++ {
			var a, b, c, o int
			for {
				a = rand.Intn(100)
				b = rand.Intn(100)
				if a < 10 || b < 10 {
					easiers++
					if easiers >= level {
						continue
					}
				}

				c = 0
				o = rand.Intn(2)
				if o == 0 {
					c = a + b
					if c > 100 {
						continue
					}
					if !carry(a, b) {
						easiers++
						if easiers >= level {
							continue
						}
					}
				} else {
					c = a - b
					if c < 0 {
						continue
					}
					if !borrow(a, b) {
						easiers++
						if easiers >= level {
							continue
						}
					}
				}

				if i%4 == 0 {
					pdf.Ln(10.2)
				}
				s := fmt.Sprintf("%-2d %s %-2d = ", a, ops[o], b)
				pdf.CellFormat(49, 10, s, "", 0, "L", false, 0, "")
				break
			}
		}
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
