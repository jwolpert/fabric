
## Channels

A **channel** is a private line of communication between two or more specific
peers, their ledgers, chaincode applications and an ordering service. Each
transaction on a Hyperledger Fabric network is executed on a specific channel,
where each party must be authenticated and authorized to transact on that
channel. Upon joining the network, a peer is assigned an identity certificate
and a derivative PKI ID by the membership services certificate authority (CA),
which authenticate each peer to its channel peers and services.

To create a new channel, the client SDK calls configuration system chaincode
with reference to the gossip metadata attributes on each target peer. This
request creates a **genesis block**, which stores the configuration information
for the channel on the ledger. When adding a new member to an existing channel,
either the genesis block, or if applicable, a more recent reconfiguration block,
is shared with the new peer.

The election of a channel **leader** determines which peer can communicate with
the ordering service on behalf of the channel. If no leader is identified, an
algorithm is used to identify the leader. The consensus service orders
transactions and transmits them, in a block, to the channel leader, which then
distributes the block across the channel.

Although any one peer (member) can belong to multiple channels, and therefore
maintain multiple ledgers, no data can pass from one channel to another. This
separation of data by channel is defined and implemented by configuration
chaincode, the identity membership service and a **gossip** data dissemination
protocol. The dissemination of data, which includes information on transactions,
ledger state and channel membership, is restricted to peers with verifiable
membership for the channel. This isolation of peers and data, by channel,
allows network members that require private and confidential transactions to
coexist with business competitors and other restricted members on the same
blockchain network.
