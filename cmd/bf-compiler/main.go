package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "strings"

  "github.com/ehehalt/brainfuck/virtualmachine"
)

func main() {
  var fileName string
  if len(os.Args) > 1 {
    fileName = os.Args[1]
  } else {
    fmt.Fprintf(os.Stderr, "error: please call with filename\n")
    os.Exit(-1)
  }

  code, err := ioutil.ReadFile(fileName)
  if err != nil {
    fmt.Fprintf(os.Stderr, "error: %s\n", err)
    os.Exit(-1)
  }

  compiler := virtualmachine.NewCompiler(string(code))
  instructions := compiler.Compile()

  fmt.Println(PreCode())

  lvl := 1
  for _, ins := range instructions {
    switch ins.Type {
    case virtualmachine.Left:
      fmt.Println(strings.Repeat(" ", lvl*2), "machine.left(", ins.Argument, ")")
    case virtualmachine.Right:
      fmt.Println(strings.Repeat(" ", lvl*2), "machine.right(", ins.Argument, ")")
    case virtualmachine.Plus:
      fmt.Println(strings.Repeat(" ", lvl*2), "machine.inc(", ins.Argument, ")")
    case virtualmachine.Minus:
      fmt.Println(strings.Repeat(" ", lvl*2), "machine.dec(", ins.Argument, ")")
    case virtualmachine.JumpIfZero:
      fmt.Println(strings.Repeat(" ", lvl*2), "for machine.get() != 0 {")
    case virtualmachine.JumpIfNotZero:
      fmt.Println(strings.Repeat(" ", (lvl-1)*2), "}")
    case virtualmachine.PutChar:
      fmt.Println(strings.Repeat(" ", lvl*2), "machine.out()")
    default:
      fmt.Println(strings.Repeat(" ", lvl*2), ins)
    }
    switch ins.Type {
    case virtualmachine.JumpIfZero:
      lvl++
    case virtualmachine.JumpIfNotZero:
      lvl--
    }
  }

  fmt.Println(PostCode())
}

// PreCode returns the code with the machine
func PreCode() string {
  return `
package main

import (
    "fmt"
    "io"
    "log"
)

type Machine struct {
    memory [30000]int
    dp     int

    input  io.Reader
    output io.Writer

    readBuf []byte
}

func (m *Machine) inc(cnt int) {
    m.memory[m.dp] += cnt
}

func (m *Machine) dec(cnt int) {
    m.memory[m.dp] -= cnt
}

func (m *Machine) right(cnt int) {
    m.dp += cnt
}

func (m *Machine) left(cnt int) {
    m.dp -= cnt
}

// get gibt den aktuellen Speicherwert zurück
func (m *Machine) get() int {
    return m.memory[m.dp]
}

func (m *Machine) out() {
    s := string(m.get())
    fmt.Print(s)
}

func NewMachine() *Machine {
    return &Machine{
        readBuf: make([]byte, 1),
    }
}

func main() {
    log.Println("Machine!")
    machine := NewMachine()
`
}

// PostCode returns the code after the generated one
func PostCode() string {
  return `}
`
}
