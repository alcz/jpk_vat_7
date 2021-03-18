package converter

import (
	"strings"

//	"github.com/toudi/jpk_vat_7/common"
)


func (c *Converter) checkXml() bool {
	c.IsXml = strings.HasSuffix(c.source, ".xml");
	// TODO: odczyt wersji formularza i schemy z XML
	metadataTemplateVars.Metadata.SchemaVersion = "1-2E"
	metadataTemplateVars.Metadata.SystemCode = "JPK_V7M (1)"
	return c.IsXml
}