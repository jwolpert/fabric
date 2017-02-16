
## Channels

A Hyperledger Fabric **channel** is a private line of communication between two
or more specific network members, for the purpose of conducting private and
confidential transactions. A channel is defined by an anchor peer for each
member, the shared ledger, chaincode applications and the ordering service
nodes. Each transaction on the network is executed on a channel, where each
party must be authenticated and authorized to transact on that channel. Upon
joining a channel, a member is assigned an identity certificate and a derivative
PKI ID by the membership services provider (MSP), which authenticate each peer
to its channel peers and services.

To create a new channel, the client SDK calls configuration system chaincode
and references metadata attributes on each **anchor peer**. This request creates
a **genesis block** for the channel ledger, which stores configuration
information about the channel policies, members and anchor peers. When adding a
new member to an existing channel, either this genesis block, or if applicable,
a more recent reconfiguration block, is shared with the new anchor peer.

The election of a **leading peer** for each member on a channel determines which
peer communicates with the ordering service on behalf of the member. If no
leader is identified, an algorithm is used to identify the leader. The consensus
service orders transactions and delivers them, in a block, to each leading peer,
which then distributes the block to its member peers, and across the channel,
using the **gossip** protocol.

Although any one anchor peer can belong to multiple channels, and therefore
maintain multiple ledgers, no ledger data can pass from one channel to another.
This separation of ledgers, by channel, is defined and implemented by
configuration chaincode, the identity membership service and the **gossip** data
dissemination protocol. The dissemination of data, which includes information on
transactions, ledger state and channel membership, is restricted to peers with
verifiable membership on the channel. This isolation of peers and ledger data,
by channel, allows network members that require private and confidential
transactions to coexist with business competitors and other restricted members,
on the same blockchain network.
