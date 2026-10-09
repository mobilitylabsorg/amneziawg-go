//go:build uniray_lowmem && !android && !ios && !windows

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2025 WireGuard LLC. All Rights Reserved.
 */

package device

import "github.com/amnezia-vpn/amneziawg-go/v3/conn"

// uniray_lowmem: for a small memory limit, such as UniRay's MikroTik
// container (64 MB). Each packet in flight takes a buffer of MaxSegmentSize
// (64 KiB), and the default pools are unbounded: a fast download fills the
// 1024-packet queues, 64 MiB each. Bounded pools make a reader wait for a
// free buffer instead, as on Android and iOS. The two UDP readers (IPv4,
// IPv6) and the TUN reader hold a batch each up front (3 * 64 with
// conn.IdealBatchSize at 64), which leaves half the bound to packets on
// their way. MaxSegmentSize stays as it is: UDP GRO reads whole coalesced
// datagrams into these buffers.
const (
	QueueStagedSize            = conn.IdealBatchSize
	QueueOutboundSize          = 1024
	QueueInboundSize           = 1024
	QueueHandshakeSize         = 1024
	MaxSegmentSize             = (1 << 16) - 1 // largest possible UDP datagram
	PreallocatedBuffersPerPool = 384           // 24 MiB of buffers at most
)
