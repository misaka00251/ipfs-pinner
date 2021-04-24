package pinner // import "github.com/misaka00251/ipfs-pinner"

import (
	"fmt"
	"mime/multipart"

	"github.com/misaka00251/ipfs-pinner/pkg/infura"
	"github.com/misaka00251/ipfs-pinner/pkg/pinata"
)

type Config struct {
	Pinner string
	Apikey string
	Secret string
}

// Pin file to pinning network
func (cfg *Config) Pin(sourceFile multipart.File, sourceFileString string) (string, error) {

	var cid string
	var err error

	switch cfg.Pinner {
	default:
		err = fmt.Errorf("%s", "unknow pinner")
	case "infura":
		cid, err = infura.PinFile(sourceFile, sourceFileString)
	case "pinata":
		pnt := pinata.Pinata{Apikey: cfg.Apikey, Secret: cfg.Secret}
		cid, err = pnt.PinFile(sourceFile, sourceFileString)
	}

	return cid, err
}
