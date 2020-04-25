package decorator

import "testing"

func TestPlain_Format(t *testing.T) {
	want := "Hello, world."
	if got := (Plain{}).Format(want); got != want {
		t.Errorf("Plain.Format() = %q, want %q", got, want)
	}
}

func TestBold_Format(t *testing.T) {
	txt := "Hello, world"
	want := "**Hello, world**"
	if got := (Bold{Plain{}}).Format(txt); got != want {
		t.Errorf("Bold.Format() = %q, want %q", got, want)
	}
}

func TestItalic_Format(t *testing.T) {
	txt := "Hello, world"
	want := "~~Hello, world~~"
	if got := (Italic{Plain{}}).Format(txt); got != want {
		t.Errorf("Italic.Format() = %q, want %q", got, want)
	}
}

func TestBoldItalic(t *testing.T) {
	txt := "Hello, world"
	want := "**~~Hello, world~~**"
	if got := (Bold{Italic{Plain{}}}).Format(txt); got != want {
		t.Errorf("BoldItalic.Format() = %q, want %q", got, want)
	}
}
