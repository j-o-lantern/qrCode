package qrCode
import (
	qr "github.com/skip2/go-qrcode"
	"fmt"
)

func getQr(str string) (string, error) {

	if qr, err := qr.New(str, qr.Highest); err != nil {
		return "", err
	} else {
		pixels := qr.Bitmap()
		return getQRString(pixels), nil
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