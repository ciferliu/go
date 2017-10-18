package currency

//ccyNtry represents the CcyNtry element in ISO 4217 XML
type ccyNtry struct {
	CtryNm     string `xml:"CtryNm"`
	CcyNm      string `xml:"CcyNm"`
	Ccy        string `xml:"Ccy"`
	CcyNbr     string `xml:"CcyNbr"`
	CcyMnrUnts string `xml:"CcyMnrUnts"`
}

//ccyTbl represents the CcyTbl element int ISO 4217 XML
type ccyTbl struct {
	CcyNtrys []ccyNtry `xml:"CcyNtry"`
}

// iso4217Xml represents the root node of ISO 4217 XML
type iso4217Xml struct {
	CcyTbl ccyTbl `xml:"CcyTbl"`
}
