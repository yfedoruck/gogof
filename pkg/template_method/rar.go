package template_method

type Rar struct {
	ArchiverTemplate
}

func NewRar() Archiver {
	a := Rar{}
	a.algorithm = func(s string) string {
		return "** rar ** " + s + " ** rar **"
	}
	return a
}
