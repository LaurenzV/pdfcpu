package validate

import (
	"testing"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func testXRefTable(version model.Version) *model.XRefTable {
	conf := model.NewDefaultConfiguration()
	conf.ValidationMode = model.ValidationStrict
	return &model.XRefTable{
		HeaderVersion:     &version,
		ValidationMode:    conf.ValidationMode,
		ValidateLinks:     conf.ValidateLinks,
		Conf:              conf,
		Table:             map[int]*model.XRefTableEntry{},
		Names:             map[string]*model.Node{},
		NameRefs:          map[string]model.NameMap{},
		KeywordList:       types.StringSet{},
		Properties:        map[string]string{},
		LinearizationObjs: types.IntSet{},
		PageAnnots:        map[int]model.PgAnnots{},
		PageThumbs:        map[int]types.IndirectRef{},
		Signatures:        map[int]map[int]model.Signature{},
		Stats:             model.NewPDFStats(),
		URIs:              map[int]map[string]string{},
		UsedGIDs:          map[string]map[uint16]bool{},
		FillFonts:         map[string]types.IndirectRef{},
	}
}

func TestValidateFileSpecDictAFRelationshipPDF17(t *testing.T) {
	xRefTable := testXRefTable(model.V17)
	d := types.Dict{
		"Type":           types.Name("Filespec"),
		"F":              types.StringLiteral("attachment.txt"),
		"UF":             types.StringLiteral("attachment.txt"),
		"AFRelationship": types.Name("Data"),
	}

	if err := validateFileSpecDict(xRefTable, d, 0); err != nil {
		t.Fatalf("validateFileSpecDict: %v", err)
	}
}

func TestValidateAssociatedFilesPDF20(t *testing.T) {
	xRefTable := testXRefTable(model.V20)
	rootDict := types.Dict{
		"AF": types.Array{
			types.Dict{
				"Type":           types.Name("Filespec"),
				"F":              types.StringLiteral("attachment.txt"),
				"UF":             types.StringLiteral("attachment.txt"),
				"AFRelationship": types.Name("Data"),
			},
		},
	}

	if err := validateAF(xRefTable, rootDict, OPTIONAL, model.V20); err != nil {
		t.Fatalf("validateAF: %v", err)
	}
}
