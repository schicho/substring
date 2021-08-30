package substring

import (
	"testing"
)

var helloWorlds = []struct {
	s   string // full string
	h   string // only Hello
	w   string // only World
	i_h int    // EndIndex to get Hello (e.g. 5 for 0-indexed)
	i_w int    // StartIndex to get World (e.g. 6 for 0-indexed)
	i_l int    // Index to get the full `Hello World` (e.g. 11 for 0-indexed)
}{
	{s: "Hello World", h: "Hello", w: "World", i_h: 5, i_w: 6, i_l: 11},
	{s: "H€ll0_Wör1d", h: "H€ll0", w: "Wör1d", i_h: 5, i_w: 6, i_l: 11},
	{s: "Привет мир", h: "Привет", w: "мир", i_h: 6, i_w: 7, i_l: 10},
	{s: "ഹലോ വേൾഡ്", h: "ഹലോ", w: "വേൾഡ്", i_h: 4, i_w: 5, i_l: 10},
}

func TestSubstring(t *testing.T) {
	for _, hw := range helloWorlds {
		var substring string
		var wanted string

		substring = Substring(hw.s, 0, hw.i_h)
		wanted = hw.h
		if wanted != substring {
			t.Fatalf("Wanted %v, but got %v", wanted, substring)
		}

		substring = Substring(hw.s, hw.i_w, hw.i_l)
		wanted = hw.w
		if wanted != substring {
			t.Fatalf("Wanted %v, but got %v", wanted, substring)
		}
	}
}

func TestSubstringStart(t *testing.T) {
	for _, hw := range helloWorlds {
		substring := SubstringStart(hw.s, hw.i_w)
		wanted := hw.w
		if wanted != substring {
			t.Fatalf("Wanted %v, but got %v", wanted, substring)
		}
	}
}

func TestSubstringEnd(t *testing.T) {
	for _, hw := range helloWorlds {
		substring := SubstringEnd(hw.s, hw.i_h)
		wanted := hw.h
		if wanted != substring {
			t.Fatalf("Wanted %v, but got %v", wanted, substring)
		}
	}
}

func TestStartIndexNegative(t *testing.T) {
	for _, hw := range helloWorlds {
		substring := SubstringStart(hw.s, -1)
		wanted := hw.s
		if wanted != substring {
			t.Fatalf("Wanted %v, but got %v", wanted, substring)
		}
	}
}

func TestStartIndexLarger(t *testing.T) {
	for _, hw := range helloWorlds {
		substring := SubstringStart(hw.s, 100)
		wanted := emptyString
		if wanted != substring {
			t.Fatalf("Wanted %v, but got %v", wanted, substring)
		}
	}
}

func TestEndIndexNegative(t *testing.T) {
	for _, hw := range helloWorlds {
		substring := SubstringEnd(hw.s, -1)
		wanted := emptyString
		if wanted != substring {
			t.Fatalf("Wanted %v, but got %v", wanted, substring)
		}
	}
}

func TestEndIndexLarger(t *testing.T) {
	for _, hw := range helloWorlds {
		substring := SubstringEnd(hw.s, 100)
		wanted := hw.s
		if wanted != substring {
			t.Fatalf("Wanted %v, but got %v", wanted, substring)
		}
	}
}

func TestStartIndexLargerEndIndex(t *testing.T) {
	for _, hw := range helloWorlds {
		substring := Substring(hw.s, hw.i_h, 0)
		wanted := emptyString
		if wanted != substring {
			t.Fatalf("Wanted %v, but got %v", wanted, substring)
		}
	}
}

func TestStartIndexEqualsEndIndex(t *testing.T) {
	for _, hw := range helloWorlds {
		substring := Substring(hw.s, 0, 0)
		wanted := emptyString
		if wanted != substring {
			t.Fatalf("Wanted %v, but got %v", wanted, substring)
		}
	}
}

func TestBothIndicesNegative(t *testing.T) {
	for _, hw := range helloWorlds {
		substring := Substring(hw.s, -5, -2)
		wanted := emptyString
		if wanted != substring {
			t.Fatalf("Wanted %v, but got %v", wanted, substring)
		}
	}
}

func TestSubstringSingleRune(t *testing.T) {
	s := "Hallo Welt"
	var substring string
	var wanted string
	substring = Substring(s, 0, 1)
	wanted = "H"
	if wanted != substring {
		t.Fatalf("Wanted %v, but got %v", wanted, substring)
	}
	substring = Substring(s, 6, 7)
	wanted = "W"
	if wanted != substring {
		t.Fatalf("Wanted %v, but got %v", wanted, substring)
	}
}

func TestSubstringMinus1To0(t *testing.T) {
	s := "Hallo Welt"
	substring := Substring(s, -1, 0)
	wanted := emptyString
	if wanted != substring {
		t.Fatalf("Wanted %v, but got %v", wanted, substring)
	}
}
