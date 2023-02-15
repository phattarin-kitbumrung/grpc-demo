package services

import (
	"io"
	"image/png"
	"os"
	"path/filepath"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type qrServer struct {
}

func NewQrServer() QrServer {
	return qrServer{}
}

func (qrServer) mustEmbedUnimplementedQrServer() {}

func (qrServer) Qr(stream Qr_QrServer) error {
	os.RemoveAll("images")
	os.Mkdir("images", os.ModePerm)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		qrCode, _ := qr.Encode(req.Message, qr.M, qr.Auto)
		qrCode, _ = barcode.Scale(qrCode, 200, 200)
		path := filepath.Join("images", req.Filename)
		file, _ := os.Create(path)
		defer file.Close()
		png.Encode(file, qrCode)

		res := QrResponse{
			Message: req.Message,
			Filename: req.Filename,
		}
		err = stream.Send(&res)
		if err != nil {
			return err
		}
	}

	return nil
}
