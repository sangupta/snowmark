/**
 * snowmark - HTML templates for Go.
 *
 * MIT License.
 * Copyright (c) 2022, Sandeep Gupta.
 * https://github.com/sangupta/snowmark
 *
 * Use of this source code is governed by a MIT style license
 * that can be found in LICENSE file in the code repository:
 */

package snowmark

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluatorBuilder(t *testing.T) {
	builder := strings.Builder{}
	ev := &Evaluator{
		builder:   &builder,
		processor: nil,
	}

	assert.Equal(t, "", ev.GetEvaluation())

	ev.WriteByte(65)
	assert.Equal(t, "A", ev.GetEvaluation())

	ev.WriteString(". ")
	assert.Equal(t, "A. ", ev.GetEvaluation())

	ev.WriteString("hello world")
	assert.Equal(t, "A. hello world", ev.GetEvaluation())

	ev.WriteRune('!')
	assert.Equal(t, "A. hello world!", ev.GetEvaluation())

	ev.Write([]byte(" you are awesome!"))
	assert.Equal(t, "A. hello world! you are awesome!", ev.GetEvaluation())
}
