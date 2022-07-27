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

import "strconv"

func ConvertToBool(value interface{}) bool {
	if value == nil {
		return false
	}

	switch v := value.(type) {
	case int, int8, int16, int32, int64:
		return v != 0

	case uint, uint8, uint16, uint32, uint64:
		return v != 0

	case float32, float64:
		return v != 0

	case string:
		s, _ := strconv.ParseBool(v)
		return s

	case bool:
		return v
	}

	// any non-nil struct is considered true
	return true
}
