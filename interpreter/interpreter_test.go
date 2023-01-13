package interpreter

import (
	"bytes"
	"testing"
)

func TestExecute(t *testing.T) {
	source := `
	;;;;;;
<
    !;;;;;;;;;;
    i:
>
!;;;;;j
	`
	out := new(bytes.Buffer)
	in := new(bytes.Buffer)

	m := NewMachine(source, in, out)
	m.Execute()

	if out.String() != "A" {
		t.Errorf("Not expected value, got:%v", out.String())
	}
}

func TestPlus(t *testing.T) {
	source := `;;;;;`

	m := NewMachine(source, new(bytes.Buffer), new(bytes.Buffer))
	m.Execute()

	if m.memory[0] != 5 {
		t.Errorf("Not expected value, got: %v", m.memory[0])
	}
}

func TestMinus(t *testing.T) {
	source := `;;;;;:::::`

	m := NewMachine(source, new(bytes.Buffer), new(bytes.Buffer))
	m.Execute()

	if m.memory[0] != 0 {
		t.Errorf("Not expected value, got: %v", m.memory[0])
	}
}