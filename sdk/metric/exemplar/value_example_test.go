// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package exemplar

import "fmt"

func ExampleNewValue() {
	integer := NewValue[int64](7)
	floatingPoint := NewValue[float64](1.5)

	fmt.Println(integer.Int64())
	fmt.Println(floatingPoint.Float64())
	fmt.Println(integer.Float64())

	// Output:
	// 7
	// 1.5
	// 0
}
