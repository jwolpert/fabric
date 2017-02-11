
## Gossip-based data dissemination protocol

Hyperledger Fabric optimizes blockchain network performance, security and
scalability by dividing workload across transaction execution (endorsing and
committing) peers and transaction ordering nodes. This decoupling of network
operations requires a secure, reliable and scalable data dissemination
protocol to ensure data integrity and consistency. To meet these requirements,
the fabric implements a **gossip data dissemination protocol**.

### Gossip protocol

The fabric gossip protocol directs each peer on a **channel** to periodically
send its data--blocks of node, channel and network information--to other peers
on the channel. Because gossip messaging is continuous and redundant, each peer
on a channel is continuously receiving current and consistent data, from
multiple peers. Any data that is transmitted from a Byzantine node, which is
data that is not validated by the security layer, is thereby identified and
filtered out of the network. Newly added peers, and peers that have gone
offline for a period of time, are also synced with the channel using the gossip
method.

The gossip-based data dissemination protocol performs three functions on a
fabric network:

1. Manages peer discovery and channel membership, by continually identifying
peers that are available, and dropping those that are not.
2. Disseminates ledger data across all peers on a channel.
3. Syncs ledger state across all peers on a channels. Any peer with data that
is out of sync with the rest of the channel identifies the missing blocks and
syncs itself by copying the correct data.

Gossip-based broadcasting operates by a single peer receiving a message from
another peer, the channel leader or the ordering service. The recipient peer
then forwards the message to randomly-selected peers (channel membership is
maintained on each peer) on the channel. This cycle repeats, with the result of
membership, ledger and state information for the channel being continually kept
current and in sync.

### Gossip messaging

Each gossip message contains the **public key infrastructure (PKI)** ID and
signature of the sender. An "alive" message sends the peer's certificate to the
channel, allowing the peer to replicate and gossip with other peers on the
channel. Online peers continually broadcast "alive" messages to indicate their
availability. Gossipping peers maintain channel membership by collecting these
alive messages; if no peer receives an alive message from a specific peer, this
"dead" peer is eventually purged from channel membership. Because "alive"
messages are cryptographically signed, peers can never receive them from a
malicious node, which lacks a signing key authorized by a root certificate
authority (CA).

In addition to the automatic forwarding of received messages, a state
reconciliation process synchronizes **world state** across peers on each
channel. Each peer continually pulls state information, in ledger blocks, from
other peers on the channel, in order to reconcile its own state if
discrepancies are identified. Because fixed connectivity is not required to 
maintain gossip-based data dissemination, the process reliably provides data
consistency and integrity to the shared ledger, including tolerance for node
crashes.

Because channels are segregated, peers on one channel cannot message or share
information on any other channel. Though any one peer can belong to multiple
channels, partitioned messaging prohibits any
cross-channel communication.

**Notes:**
1. Point-to-point messages are handled by the peer TLS layer, and do not
require signatures. Peers are authenticated by their certificates, which are
assigned by a CA. Although TLS certs are also used, it is the peer certificates
that are authenticated, which are not TLS certificates (identities). Ledger
blocks are signed by the ordering service, and are neither point-to-point nor
signed by the peers.
2. Authentication is governed by the membership service provider for the peer.
When the peer connects to the channel for the first time, the TLS session binds
with fabric membership identity. This essentially authenticates each peer
to the connecting peer, with respect to membership in the network and channel.
3. Ordering service blocks contain the latest configuration block sequence
number for the channel. Peers check whether a received block was sent by a peer
with channel membership, by comparing its configuration block sequence number
to the sender's number. If they do not match, then the peer does not forward
the block, pending further verification of the sender's channel membership.
