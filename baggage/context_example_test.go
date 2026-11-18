// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package baggage_test

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/baggage"
)

func ExampleContextWithoutBaggage() {
	b, err := baggage.Parse("tenant=acme")
	if err != nil {
		panic(err)
	}
	original := baggage.ContextWithBaggage(context.Background(), b)
	derived := baggage.ContextWithoutBaggage(original)
	fmt.Println(baggage.FromContext(original).Len())
	fmt.Println(baggage.FromContext(derived).Len())
	// Output:
	// 1
	// 0
}
