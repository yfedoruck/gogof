package abstract_factory

type Os interface {
	GetArchiver() Archiver
}

type Windows struct {}

func (w Windows) GetArchiver() Archiver {
	return Rar{}
}

type Linux struct {}

func (l Linux) GetArchiver() Archiver {
	return Zip{}
}

type Archiver interface {
	Archive(s string) string
}

type Zip struct {
}

func (z Zip) Archive(s string) string {
	return "[zip++" + s + "++zip]"
}

type Rar struct {
}

func (r Rar) Archive(s string) string {
	return "[rar++" + s + "++rar]"
}
