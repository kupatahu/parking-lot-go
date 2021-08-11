package prompt

import (
	"bufio"
	"fmt"
	"strings"
)

type Prompt struct {
	reader *bufio.Reader
}

func New(r *bufio.Reader) *Prompt {
	return &Prompt{reader: r}
}

func (p *Prompt) GetInput(q string) (string, error) {
	fmt.Print(q)
	s, err := p.reader.ReadString('\n')

	s = strings.TrimSpace(s)

	return s, err
}