package uploader

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/toudi/jpk_vat_7/common"
)

var logger *log.Logger
var httpClient = http.DefaultClient

type Uploader struct {
	sourceFile      string
	workdir         string
	UseTestGateway  bool
	referenceNumber string
}

func UploaderInit(sourceFile string, verbose bool) *Uploader {
	uploader := &Uploader{sourceFile: sourceFile, workdir: path.Dir(sourceFile)}

	logger = log.New()
	if verbose {
		logger.SetLevel(log.DebugLevel)
	}

	return uploader
}

func (u *Uploader) UploadSAFTFile() error {
	logger.Infof("Rozpoczynam wysyłanie pliku; Adres bramki: %v", u.gatewayURL())
	var err error
	if err = u.uploadSAFTFile(); err != nil {
		return fmt.Errorf("Błąd rozpoczynania sesji: %v", err)
	}

	return nil
}

func (u *Uploader) saftRefNoFile() string {
	return path.Join(u.workdir, strings.TrimSuffix(path.Base(u.sourceFile), "-metadata.xml")+".ref")
}

func (u *Uploader) saftZipAesFile() string {
	return path.Join(u.workdir, strings.TrimSuffix(path.Base(u.sourceFile), "-metadata.xml")+".zip.aes")
}

func (u *Uploader) gatewayURL() string {
	if u.UseTestGateway {
		return common.TestGatewayURL
	}
	return common.ProductionGatewayURL
}
