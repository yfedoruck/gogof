package builder

import "fmt"

type Pizza struct {
	Filling   string
	Border    string
	Seasoning string
}

func (p Pizza) String() string {
	return fmt.Sprintf("build Pizza from: %s, %s, %s", p.Filling, p.Border, p.Seasoning)
}
