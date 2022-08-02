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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModel(t *testing.T) {
	model := NewModel()

	assert.Equal(t, 0, model.Size())
	assert.True(t, model.IsEmpty())
	assert.NotNil(t, model.GetMap())

	model.Put("hello", "world")
	assert.Equal(t, "world", model.GetString("hello", ""))
	assert.Equal(t, "", model.GetString("hello-not-exists", ""))
	assert.Equal(t, 1, model.Size())
	assert.False(t, model.IsEmpty())

	assert.False(t, model.PutIfNotExists("hello", "world"))
	model.Remove("hello")
	assert.False(t, model.Replace("hello", "world2"))
	assert.True(t, model.PutIfNotExists("hello", "world"))
	assert.True(t, model.Replace("hello", "world2"))
	assert.Equal(t, "world2", model.GetString("hello", ""))
	model.Remove("hello")
	assert.True(t, model.PutIfNotExists("hello", "world"))
	assert.True(t, model.Replace("hello", "world2"))
	assert.Equal(t, 1, model.Size())
	assert.False(t, model.IsEmpty())
	model.Clear()
	assert.Equal(t, 0, model.Size())
	assert.True(t, model.IsEmpty())

	model.Clear()

	// get tests
	value, exists := model.Get("hello")
	assert.False(t, exists)
	assert.Nil(t, value)

	// exists
	model.Put("hello", 2)
	assert.True(t, model.GetBool("hello", false))
	assert.False(t, model.GetBool("hello-not-exists", false))

	assert.Equal(t, uint64(2), model.GetUInt64("hello", 0))
	assert.Equal(t, uint64(0), model.GetUInt64("hello-no-exists", 0))

	assert.Equal(t, int64(2), model.GetInt64("hello", 0))
	assert.Equal(t, int64(0), model.GetInt64("hello-no-exists", 0))

	assert.Equal(t, float64(2), model.GetFloat64("hello", 0))
	assert.Equal(t, float64(0), model.GetFloat64("hello-no-exists", 0))
}
