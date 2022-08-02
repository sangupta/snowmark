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

import "github.com/sangupta/berry"

//
// Represents a model of values
// that are used to merge with a page or a fragment.
//
type Model struct {
	_map map[string]interface{}
}

//
// Create new model instance.
//
func NewModel() *Model {
	return &Model{
		_map: make(map[string]interface{}),
	}
}

//
// Return the size of the model.
//
func (model *Model) Size() int {
	return len(model._map)
}

//
// Check if model is empty or not. Returns `true` if
// there are no keys in model, `false` otherwise.
//
func (model *Model) IsEmpty() bool {
	return model.Size() == 0
}

//
// Clear model and remove all keys
//
func (model *Model) Clear() {
	model._map = make(map[string]interface{})
}

//
// Return the map associated with this model.
//
func (model *Model) GetMap() map[string]interface{} {
	return model._map
}

//
// Get the value from the model.
//
func (model *Model) Get(key string) (interface{}, bool) {
	value, exists := model._map[key]
	return value, exists
}

//
// Get a `string` value for a key from the model, or the default value if
// the key does not exist, or is not a `string`.
//
func (model *Model) GetString(key string, defaultValue string) string {
	value, exists := model._map[key]
	if !exists {
		return defaultValue
	}

	return berry.ConvertToString(value)
}

//
// Get a `bool` value for a key from the model, or the default value if
// the key does not exist, or is not a `bool`.
//
func (model *Model) GetBool(key string, defaultValue bool) bool {
	value, exists := model._map[key]
	if !exists {
		return defaultValue
	}

	b, _ := berry.ConvertToBool(value)
	return b
}

//
// Get a `uint64` value for a key from the model, or the default value if
// the key does not exist, or is not a `uint64`.
//
func (model *Model) GetUInt64(key string, defaultValue uint64) uint64 {
	value, exists := model._map[key]
	if !exists {
		return defaultValue
	}

	b, _ := berry.ConvertToUint64(value, defaultValue)
	return b
}

//
// Get a `int64` value for a key from the model, or the default value if
// the key does not exist, or is not a `int64`.
//
func (model *Model) GetInt64(key string, defaultValue int64) int64 {
	value, exists := model._map[key]
	if !exists {
		return defaultValue
	}

	b, _ := berry.ConvertToInt64(value, defaultValue)
	return b
}

//
// Get a `float64` value for a key from the model, or the default value if
// the key does not exist, or is not a `float64`.
//
func (model *Model) GetFloat64(key string, defaultValue float64) float64 {
	value, exists := model._map[key]
	if !exists {
		return defaultValue
	}

	b, _ := berry.ConvertToFloat64(value, defaultValue)
	return b
}

//
// Put the value in model against given key.
//
func (model *Model) Put(key string, value interface{}) {
	model._map[key] = value
}

//
// Put the value in model if the key does not exist. Returns
// `true` if value was added to model, `false` otherwise.
//
func (model *Model) PutIfNotExists(key string, value interface{}) bool {
	_, exists := model._map[key]
	if exists {
		return false
	}

	model._map[key] = value
	return true
}

//
// Replace the current key with new value. If no key exists
// the function returns `false`. Returns `true` if value was
// replaced.
//
func (model *Model) Replace(key string, value interface{}) bool {
	_, exists := model._map[key]
	if !exists {
		return false
	}

	model._map[key] = value
	return true
}

//
// Remove the key from the model, if it exists.
//
func (model *Model) Remove(key string) {
	delete(model._map, key)
}
