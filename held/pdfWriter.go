package held

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"strconv"
)

func strDelimit(str string, sepstr string, sepcount int) string {
	pos := len(str) - sepcount
	for pos > 0 {
		str = str[:pos] + sepstr + str[pos:]
		pos = pos - sepcount
	}
	return str
}

func createTable(pdf *gofpdf.Fpdf, header []string, data [][]string, columnSize []float64) {
	if pdf == nil {
		return
	}
	fmt.Println(header)
	fmt.Println("---")
	fmt.Println(data)
	fmt.Println("---")
	fmt.Println(columnSize)
	fmt.Println("---")

	cLen := len(columnSize)
	if cLen == 0 {
		return
	}
	fmt.Println("cLen > 0")
	if len(header) != 0 && cLen != len(header) {
		return
	}
	fmt.Println("len header 0 or eq to cLen")
	for _, v := range data {
		if len(v) != cLen {
			return
		}
	}
	fmt.Println("Check done")
	pdf.SetFillColor(128, 128, 128)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetDrawColor(128, 0, 0)
	pdf.SetLineWidth(.3)
	pdf.SetFont("Arial", "", 14)
	// display header fields
	for i, str := range header {
		pdf.CellFormat(columnSize[i], 7, str, "1", 0, "C", true, 0, "")
	}
	// only advance a line if we wrote a header
	if len(header) > 0 {
		pdf.Ln(-1)
	}
	// display data /w alternating background
	pdf.SetFillColor(224, 235, 255)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("", "", 0)
	//  Data
	fill := false
	for _, c := range data {
		for i, str := range c {
			pdf.CellFormat(columnSize[i], 6, str, "LR", 0, "", fill, 0, "")
		}
		pdf.Ln(-1)
		fill = !fill
	}
}

func (h *Held) allgemeineInformationenPDFData() (header []string, data [][]string, cols []float64) {
	header = make([]string, 0)
	data = make([][]string, 3)
	cols = []float64{25, 60, 45, 80}
	data[0] = []string{"Name", h.Name, "Spezies", h.Spezies.Name}
	data[1] = []string{"Kultur", h.Kultur.Name, "Profession", h.Profession.Name}
	data[2] = []string{"AP (frei)", strconv.Itoa(h.AP), "AP (ausgegeben)", strconv.Itoa(h.AP_spent)}
	return
}

func (h *Held) eigenschaftenPDFData() (header []string, data [][]string, cols []float64) {
	header = []string{"Eigenschaft", "Wert"}
	cols = []float64{70, 35}
	data = make([][]string, len(h.Eigenschaften.Eigenschaften))
	idx := 0
	for _, v := range h.Eigenschaften.Eigenschaften {
		data[idx] = make([]string, 2)
		data[idx][0] = v.Name
		data[idx][1] = strconv.Itoa(v.Wert)
		idx++
	}
	return
}

func (h *Held) fkPDFData() (header []string, data [][]string, cols []float64) {
	header = []string{"Kampftechnik", "SK", "Eigenschaften", "Wert", "FK"}
	cols = []float64{50, 10, 40, 20, 20}
	data = make([][]string, 3)
	idx := 0
	for _, v := range h.Kampftechniken.Kampftechniken {
		if v.IsFernkampf {
			data[idx] = make([]string, 5)
			data[idx][0] = v.Name
			data[idx][1] = v.SK
			data[idx][2] = ""
			for i, w := range v.Leiteigenschaften {
				data[idx][2] += w.Name
				if i < len(v.Leiteigenschaften) {
					data[idx][2] += " "
				}
			}
			data[idx][3] = strconv.Itoa(v.Wert)
			data[idx][4] = strconv.Itoa(v.FK())
			idx++
		}
	}
	return
}

func (h *Held) nkPDFData() (header []string, data [][]string, cols []float64) {
	header = []string{"Kampftechnik", "SK", "Eigenschaften", "Wert", "AT", "PA"}
	cols = []float64{50, 10, 40, 20, 10, 10}
	data = make([][]string, len(h.Kampftechniken.Kampftechniken)-3)
	idx := 0
	for _, v := range h.Kampftechniken.Kampftechniken {
		if !v.IsFernkampf {
			data[idx] = make([]string, 6)
			data[idx][0] = v.Name
			data[idx][1] = v.SK
			data[idx][2] = ""
			for i, w := range v.Leiteigenschaften {
				data[idx][2] += w.Name
				if i < len(v.Leiteigenschaften) {
					data[idx][2] += " "
				}
			}
			data[idx][3] = strconv.Itoa(v.Wert)
			data[idx][4] = strconv.Itoa(v.AT())
			if !v.NurAttacke {
				data[idx][5] = strconv.Itoa(v.PA())
			} else {
				data[idx][5] = "-"
			}
			idx++
		}
	}
	return
}

func (h *Held) ToFile(fname string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	// first page
	pdf.AddPage()
	// allgemeine helden-informationen:
	header, data, cols := h.allgemeineInformationenPDFData()
	createTable(pdf, header, data, cols)
	header, data, cols = h.eigenschaftenPDFData()
	createTable(pdf, header, data, cols)

	//pdf.AddPage()
	header, data, cols = h.fkPDFData()
	createTable(pdf, header, data, cols)
	header, data, cols = h.nkPDFData()
	createTable(pdf, header, data, cols)

	fmt.Println("/tmp/" + fname + ".x.pdf")
	err := pdf.OutputFileAndClose("/tmp/" + fname + "x.pdf")
	fmt.Println(err)
	return err
}
