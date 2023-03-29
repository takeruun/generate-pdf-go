package generatepdfgo

import (
	"image"
	"os"

	"github.com/signintech/gopdf"
)

func generatepdfgo() {
	imagePath := "./images/base.jpg"
	imageFile, err := os.Open(imagePath)
	if err != nil {
		return
	}
	defer imageFile.Close()

	imagefile, _, err := image.Decode(imageFile)
	if err != nil {
		return
	}

	// 画像のサイズを取得
	bounds := imagefile.Bounds()
	width := float64(bounds.Dx())
	height := float64(bounds.Dy())

	// PDFファイルを作成
	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: width, H: height}})
	pdf.AddPage()

	// 画像をPDFに描画
	pdf.Image(imagePath, 0, 0, nil)

	// フォントを設定
	fontPath := "./ttf/ipaexm.ttf"
	pdf.AddTTFFont("IPAexM", fontPath)
	pdf.SetFont("IPAexM", "", 16)

	pdf.SetFontSize(30)
	pdf.SetTextColor(0, 0, 0)

	// 名前
	pdf.SetX(450)
	pdf.SetY(150)
	pdf.Cell(nil, "名前")

	// genrate pdf
	pdf.WritePdf("./pdfs/generatepdfgo.pdf")
}
