package art

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
)

var aa = "" +
	"H4sIAGdplloAA3u/Z/97IhBXzeOGBgVcsvHxINkasKIaKwwAFiZWUkH//f7FEJTwqGfS+z170MxW" +
	"gDhlXycQPd0+93Hz3vgiWxu4KRCtYDVAWQgCCe+bBVISHw9UBZIHYoUnO/ZDlb3YNA9IgsTe75sB" +
	"RBDTgD4DmaMPxOjGgRSApLmgcvrggOiHK+ECC2LTzAVUBhTwAglC5CDiQFu5oD6CyCBkQVJcAE4i" +
	"5FepAQAA"

// DrawAA 貞子のAAを描画する
func DrawAA() {

	b, err := base64.StdEncoding.DecodeString(aa)
	if err != nil {
		panic(err)
	}
	r, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	if _, err = io.Copy(&buf, r); err != nil {
		panic(err)
	}
	fmt.Println(buf.String())

}
