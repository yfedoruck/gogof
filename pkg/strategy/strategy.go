package strategy

type Archiver interface {
	Archive(s string) string
}

type Rar struct {
}

func (r *Rar) Archive(s string) string {
	return "/rar/ " + s + " /rar/"
}

type Zip struct {
}

func (r *Zip) Archive(s string) string {
	return "~zip~ " + s + " ~zip~"
}

func Strategy(s string) Archiver {
	var a Archiver
	switch s {
	case "windows":
		return &Rar{}
	case "linux":
		return &Zip{}
	}
	return a
}
