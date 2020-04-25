package template_method

type Archiver interface {
	Pack(s string) string
}

type ArchiverTemplate struct {
	algorithm func(s string) string
}

func (r ArchiverTemplate) Pack(s string) string {
	str := "packed start\n"
	str += r.algorithm(s) + "\n"
	str += "packed end\n"

	return str
}
