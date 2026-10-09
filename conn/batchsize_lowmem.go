//go:build uniray_lowmem

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2025 WireGuard LLC. All Rights Reserved.
 */

package conn

// uniray_lowmem: smaller batches, for a small memory limit (see
// device/queueconstants_lowmem.go). Each reader holds a batch of buffers
// before any packet comes, so the batch must stay well below the buffer
// pools' bound; it must also stay a multiple of udpSegmentMaxDatagrams (64),
// the most datagrams the kernel coalesces into one read (UDP GRO).
const (
	IdealBatchSize = 64 // maximum number of packets handled per read and write
)
