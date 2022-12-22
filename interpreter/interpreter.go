package interpreter

import (
	"hthcrwzy/anySemicolon/tokens"
	"io"
	"log"
)

const data_size int = 300000

type Machine struct {
	code string
	ip   int

	memory [data_size]int
	dp     int

	input io.Reader
	output io.Writer

	buf []byte
}

func NewMachine(code string, in io.Reader, out io.Writer) *Machine  {
	return &Machine{
		code: code,
		input: in,
		output: out,
		buf: make([]byte, 1),
	}
}

func (m *Machine) Execute() {
	for m.ip < len(m.code) {
		token := m.code[m.ip]

		switch token {
		case tokens.SEMICOLON:
			m.memory[m.dp]++
		case tokens.COLON:
			m.memory[m.dp]--
		case tokens.BANG:
			m.dp++
		case tokens.I:
			m.dp--
		case tokens.LEFT_ANGLE_BRACKET:
			if m.memory[m.dp] == 0 {
				depth := 1
				for depth != 0 {
					m.ip++
					switch m.code[m.ip] {
					case tokens.LEFT_ANGLE_BRACKET:
						depth++ // 更に深いループ
					case tokens.RIGHT_ANGEL_BRACKET:
						depth--
					}
				}
			}
		case tokens.RIGHT_ANGEL_BRACKET:
			if m.memory[m.dp] != 0 {
				depth := 1
				for depth != 0 {
					m.ip--
					switch m.code[m.ip] {
					case tokens.RIGHT_ANGEL_BRACKET:
						depth++
					case tokens.LEFT_ANGLE_BRACKET:
						depth--
					}
				}
			}
		case tokens.J:
			m.putChar()
		case tokens.QUESTION:
			m.readCher()

		default:
			// PASS
		}

		m.ip++
	}
}

func (m *Machine) readCher() {
	n, err := m.input.Read(m.buf) // バッファにInputの内容を書き込む
	if err != nil {
		log.Fatalf("%v", err)
	}
	if n != 1 {
		log.Fatalf("Wrong num bytes read")
	}

	m.memory[m.dp] = int(m.buf[0]) // TODO: 入力に変更
}

func (m *Machine) putChar() {
	m.buf[0] = byte(m.memory[m.dp])

	n, err := m.output.Write(m.buf) // バッファの内容をOutputに書き出す。
	if err != nil {
		log.Fatalf("%v", err)
	}
	if n != 1 {
		log.Fatalf("Wrong num bytes written")
	}
}
