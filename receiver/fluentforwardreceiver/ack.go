// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package fluentforwardreceiver // import "github.com/GlancingMind/opentelemetry-collector-contrib/receiver/fluentforwardreceiver"

import "github.com/tinylib/msgp/msgp"

type AckResponse struct {
	Ack string `msg:"ack"`
}

func (z AckResponse) EncodeMsg(en *msgp.Writer) error {
	// map header, size 1
	// write "ack"
	err := en.Append(0x81, 0xa3, 0x61, 0x63, 0x6b)
	if err != nil {
		return err
	}

	err = en.WriteString(z.Ack)
	if err != nil {
		return msgp.WrapError(err, "Ack")
	}
	return nil
}
