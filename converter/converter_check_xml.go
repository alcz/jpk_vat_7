package converter

import (
	"strings"
	"fmt"
	"io/ioutil"
	"encoding/xml"

//	"github.com/toudi/jpk_vat_7/common"
)

type saftFormCode struct {
        Value         string `xml:",chardata"`
        SystemCode    string `xml:"kodSystemowy,attr"`
        SchemaVersion string `xml:"wersjaSchemy,attr"`
}

type saftHeader struct {
        FormCode saftFormCode `xml:"KodFormularza"`
}

type saftXml struct {
        XmlName xml.Name
        Header  saftHeader `xml:"Naglowek"`
}

func (c *Converter) checkXml() bool {
	var err error
	var saftHeader saftXml
	var saftXmlBytes []byte

	c.IsXml = strings.HasSuffix(c.source, ".xml");

	saftXmlBytes, err = ioutil.ReadFile(c.source)

	if err != nil {
		fmt.Printf("nie udało się odczytać zawartości pliku JPK: %v", err)
	} else {
		if err = xml.Unmarshal(saftXmlBytes, &saftHeader); err != nil {
			fmt.Printf("nie udało się sparsować nagłówka JPK do struktury: %v", err)
		} else {
			metadataTemplateVars.Metadata.FormCode = saftHeader.Header.FormCode.Value
			metadataTemplateVars.Metadata.SystemCode = saftHeader.Header.FormCode.SystemCode
			metadataTemplateVars.Metadata.SchemaVersion = saftHeader.Header.FormCode.SchemaVersion
		}
	}

        if metadataTemplateVars.Metadata.SchemaVersion == "" {
		metadataTemplateVars.Metadata.SchemaVersion = "1-2E"
	}
        if metadataTemplateVars.Metadata.SystemCode == "" {
		metadataTemplateVars.Metadata.SystemCode = "JPK_V7M (1)"
        }

	return c.IsXml
}