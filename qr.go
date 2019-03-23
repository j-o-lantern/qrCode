package grCode
import (
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
	"os"
)

func getQr(str string) () {

	if qr, err := qrcode.New(str, qrcode.Highest); err != nil {
		fmt.Errorf("Error: %s", err)
	} else {
		pixels := qr.Bitmap()
		fmt.Fprint(os.Stderr, getQRString(pixels));
	}
}

func getQRString(pixels [][]bool) (string) {

	str := ""
	for ir, row := range pixels {
		lr := len(row)
		if ir == 0 || ir == 1 || ir == 2 ||
			ir == lr-1 || ir == lr-2 || ir == lr-3 {
			continue
		}
		for ic, col := range row {
			lc := len(pixels)
			if ic == 0 || ic == 1 || ic == 2 ||
				ic == lc-1 || ic == lc-2 || ic == lc-3 {
				continue
			}
			if col {
				str += fmt.Sprint("\033[48;5;0m  \033[0m")
			} else {
				str += fmt.Sprint("\033[48;5;7m  \033[0m")
			}
		}
		str += fmt.Sprintln()
	}

	return str
}