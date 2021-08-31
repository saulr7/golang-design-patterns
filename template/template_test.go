package template

import (
	"strings"
	"testing"
)

func TestTemplate_ExecuteAlgorithm(t *testing.T) {

	t.Run("Using interfaces", func(t *testing.T) {
		s := &TemplateImpl{}
		s1 := &TestSturct{}

		res := s.ExecuteAlgorithm(s1)
		expected := "world"

		if !strings.Contains(res, expected) {
			t.Errorf("Expected string '%s' wasn't found on returned string\n", expected)
		}
	})
}
