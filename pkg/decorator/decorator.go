package decorator

import (
	"fmt"
)

type Text interface {
	Format(s interface{}) string
}

type Plain struct {
}

func (p Plain) Format(s interface{}) string {
	return fmt.Sprintf("%s", s)
}

type Bold struct {
	Text Text
}

func (b Bold) Format(s interface{}) string {
	return "**" + b.Text.Format(s) + "**"
}

type Italic struct {
	Text Text
}

func (b Italic) Format(s interface{}) string {
	return "~~" + b.Text.Format(s) + "~~"
}
