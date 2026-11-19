// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metricdata_test

import (
	"fmt"

	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

func ExampleExtrema() {
	var unset metricdata.Extrema[int64]
	zero := metricdata.NewExtrema[int64](0)

	fmt.Println(unset.Value())
	fmt.Println(zero.Value())

	// Output:
	// 0 false
	// 0 true
}
