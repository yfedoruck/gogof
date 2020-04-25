package template_method

type Zip struct {
	ArchiverTemplate
}

func NewZip() Archiver {
	a := Zip{}
	a.algorithm = func(s string) string {
		return "== zip == " + s + " == zip =="
	}
	return a
}
