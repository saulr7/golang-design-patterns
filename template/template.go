package template

import "strings"

type MessageRetriever interface {
	Message() string
}
type Template interface {
	first() string
	third() string
	ExecuteAlgorithm(MessageRetriever) string
}

type TemplateImpl struct{}

func (t *TemplateImpl) first() string {
	return "hello"
}

func (t *TemplateImpl) third() string {
	return "t"
}

func (t *TemplateImpl) ExecuteAlgorithm(m MessageRetriever) string {
	return strings.Join([]string{t.first(), m.Message(), t.third()}, " ")
}

type TetsSturct struct {
	Template
}

func (m TetsSturct) Message() string {
	return "world"
}
