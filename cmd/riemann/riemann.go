package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

var (
	options = gofpdf.ImageOptions{ImageType: "png", ReadDpi: true, AllowNegativePosition: false}
)

func Names(resType string) (names []string) {
	for key := range res {
		if strings.HasPrefix(key, resType) {
			name := strings.TrimPrefix(key, resType)
			names = append(names, name)
		}
	}
	return
}

func FontBytes(name string) (b []byte) {
	if !strings.HasPrefix(name, "font.") {
		name = fmt.Sprintf("font.%s", name)
	}
	in := res[name]
	r, err := gzip.NewReader(bytes.NewBuffer(in))
	if err != nil {
		return
	}
	defer r.Close()
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return
}

func LoadFonts(pdf *gofpdf.Fpdf) {
	for _, fn := range Names("font.") {
		pdf.AddUTF8FontFromBytes(fn, "", FontBytes(fn))
	}
}

func RegisterPngImages(pdf *gofpdf.Fpdf) {
	for _, png := range Names("png.") {
		pdf.RegisterImageOptionsReader(png, options, bytes.NewBuffer(res[fmt.Sprintf("png.%s", png)]))
	}
}

func NewFpdf(header, footer, body string, fsize float64) (pdf *gofpdf.Fpdf) {
	pdf = gofpdf.New("P", "mm", "A4", "")
	LoadFonts(pdf)
	RegisterPngImages(pdf)
	pdf.SetLeftMargin(12.5)
	pdf.SetTopMargin(7.5)
	pdf.SetHeaderFuncMode(func() {
		pdf.SetFont(header, "", fsize)
		pdf.Cell(50, 10, "姓名：")
		pdf.Cell(48, 10, "日期：")
		pdf.Cell(46, 10, "时间：")
		pdf.Cell(25, 10, "分数：")
		pdf.ImageOptions("riemann", -1, 6, 10, 10, false, options, 0, "")
		width, _ := pdf.GetPageSize()
		pdf.Line(10, 20, width-10, 20)
		pdf.SetFont(body, "", fsize)
	}, true)
	pdf.SetFooterFunc(func() {
		pdf.SetFont(footer, "", fsize)
		pdf.SetY(-15)
		width, _ := pdf.GetPageSize()
		pdf.SetX(width/2 - 5)
		pdf.Cell(10, 10, fmt.Sprintf("%-2d", pdf.PageNo()))
	})
	return
}

func carry(a, b int) bool {
	a = a % 10
	b = b % 10
	if a+b >= 10 {
		return true
	}
	return false
}

func borrow(a, b int) bool {
	a = a % 10
	b = b % 10
	if a-b < 0 {
		return true
	}
	return false
}

func WritePages100(pdf *gofpdf.Fpdf, pages, maximum, level int) {
	ops := map[int]string{0: "+", 1: "-"}
	for p := 0; p < pages; p++ {
		pdf.AddPage()
		pdf.Ln(3.5)
		easiers := 0
		for i := 0; i < maximum; i++ {
			var a, b, c, o int
			for {
				a = rand.Intn(maximum)
				b = rand.Intn(maximum)
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
					if c > maximum {
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
}
