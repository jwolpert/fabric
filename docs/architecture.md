Hyperledger Fabric is a unique implementation of distributed ledger technology
(DLT) that ensures data integrity and consistency, while delivering
accountability, transparency and efficiencies unmatched by other blockchain
or DLT technology.

The Hyperledger Fabric implements a specific type of permissioned blockchain
network over which members may track, exchange and interact with digitized
assets using transactions that are governed by smart contracts - what we call
chaincode - in a secure and robust manner, while enabling participants in the
network to interact in a manner that assures that their transactions and data
can be restricted to an identified subset of network participants - something
we call a channel.

The blockchain network supports the ability for members to establish a set
of shared ledgers that contain the source of truth about those digitized
assets, and recorded transactions, that is replicated in a secure manner only
to the set of nodes participating in that channel.

The Hyperledger Fabric architecture is comprised of the following components:
peer nodes, ordering nodes, certificate authority nodes and the clients, that
are likely leveraging one of the language-specific Fabric SDKs.

The Hyperledger Fabric peer node maintains the ledger/state, manages the
execution of chaincode (smart contracts) and in certain configurations, is
responsible for simulating execution of and endorsing transactions. Such nodes
are called endorsing peers or simply endorsers.

The orderers consent on the order of transactions in a block to be committed to
the ledger. In common blockchain architectures (including Hyperledger Fabric
v0.6 and earlier) the roles played by the peer and orderer nodes were unified
(cf. validating peer in Hyperledger Fabric v0.6).

Two or more participants may create and join a channel and begin to interact.
Initially, the members in a channel agree on the terms of the chaincode (smart
contract) that will govern the transaction(s). When consensus is reached on the 
proposal to deploy a given chaincode, it is committed to the ledger.

Once the chaincode is deployed to the peer nodes in the channel, network
participants with the right privileges can propose transactions on the channel
by using one of the language-specific client SDKs to invoke functions on the
deployed chaincode.

The proposed transactions are sent to endorsers that evaluate and apply the
endorsement policy for the proposed transaction and return the result to the
client that initiated the proposal.

The client application ensures that the results from the endorsers is consistent
and signed by the appropriate endorsers, according to the endorsement policy.
The application then sends the transaction, the result and endorsements,  to the
ordering nodes.  

Ordering nodes order the transactions, result and endorsements received from
the clients into a block which is then sent to the peer nodes to be committed
to the ledger.  The peers than execute the transactions, validates that the
each result is consistent with the expected result, and commits the block to
the ledger.  

Some key capabilities of Hyperledger Fabric include:

- Allows for complex query for applications that need ability to handle complex
data structures.

- Enables decisions to remain with the stakeholders with n-lateral transactions
(endorsement model, multichain)- need to turn this into something more readable.

- Implements a permissioned network, also known as a consortia network, where
all members are known to each other.

- Incorporates a modular approach to various capabilities, enabling network
designers to plug in their preferred implementations for various capabilities
such as consensus (ordering), identity management and encryption.

- Ability to have multiple channels that are isolated from one another that
allows for multi-lateral transactions amongst select peer nodes that assures
high degrees of privacy and confidentiality required by competing businesses
and highly regulated industries on a common network.

- Network scability and performance are achieved through separation of
chaincode execution from transaction ordering, which limits the required levels
of trust and verification across nodes for optimization.
