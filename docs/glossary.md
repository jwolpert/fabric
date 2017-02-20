# Hyperledger Fabric Glossary

*Note: This glossary is structured to prioritize new terms and features specific
to architecture.  It makes the assumption that one already possesses a
working familiarity with the basic tenets of blockchain.*

## Blockchain Network
A blockchain network consists of, at minimum, one peer
(responsible for endorsing and committing transactions) leveraging an ordering
service, and a membership services component (certificate authority) that
distributes and revokes cryptographic certificates representative of user
identities and permissions.

## Permissioned Network
A blockchain network where any entity (node) is
required to maintain a member identity on the network.  End users must be
authorized and authenticated in order to use the network. Permissioned networks can be restrictive or permissive, depending on the needs of the network. The Internet itself is a permissioned network, with ICANN maintaining a registry of domains and ownership, but it is highly permissive, not requiring extensive identity proofing and de-duplication. In a highly restrictive permissioned network, the most essential task of permissioning is ensuring that each party on the network is not only the party the members of that network believe it to be but that it is the ONLY such party. Hence, it is generally desirable for parties to maintain validated identities and enrich them over time to enhance trust among the networks they participate in.  

## Peer
Peer is a component that executes, and maintains a ledger of, transactions.  
There are two roles for a peer – endorser and committer.  The architecture has
been designed such that a peer is always a committer, but not necessarily always
an endorser.  Peers play no role in the ordering of transactions.

## Member
A Member is a participant (such as a company or organization) that
operates components - Peers, Orderers, and applications - in the blockchain
network.  A member is identified by its CA certificate (i.e. a unique enrollment).  
A Member’s peer will be leveraged by end users in order to perform transaction
operations on specific channels.

## Transaction
Refers to an operation in which an authorized end user
performs read/write operations against the ledger.  There are three unique types
of transactions - deploy, invoke, and query.

## End User
An end user (User for short) is someone who interacts with the blockchain
through a set of published APIs (i.e. the hfc SDK).  You can have an admin user
who can grant permissions to other users, a client
user, who can drive deploy, invoke, query transactions on various channels.  In the case
of self-executing transactions, the application itself can also be thought of
as the end user. [Let's talk about this.]

## Ordering Service
A centralized or decentralized service that orders
transactions in a block.  You can select different implementations of the
"ordering" function - e.g "solo" for simplicity and testing, Kafka for crash
fault tolerance, or sBFT/PBFT for byzantine fault tolerance. You can also
develop your own protocol and offer it as a service.

## Consensus
A broader term overarching the entire transactional flow, which
serves to generate an agreement on the ordering of transactions and to confirm the correctness of
the set of transactions constituting a block.

## Orderer or Node
One of the runtime nodes that form the ordering service.  A
collection of ordering service nodes (OSNs) will order transactions into blocks
according to the network's chosen ordering implementation.  In the case of
"solo", only one OSN is required.  Transactions are "broadcast" to orderers, and
then "delivered" as blocks to the appropriate channel.

## Endorser
A specific function performed by Peers. A Peer performing the Endorsement function is responsible for
simulating transactions, and in turn preventing unstable or non-deterministic
transactions from passing through the network.  A transaction is sent to an
endorser in the form of a transaction proposal.  All Peers that perform Endorsement must also perform Committing (i.e. they write to the ledger — See ## Committer).

## Committer
A specific function performed by Peers. A Peer performing the Committer function appends the
validated transactions to the ledger [let's talk about whether and how to introduce the notion of channels here].  A peer can act as both
an endorser and committer, but in more regulated circumstances might only serve
as a committer. All Peers that perform Endorsement must also perform Committing (i.e. they write to the ledger — See ## Committer), but not all Peers that Commit Transactions to the Ledger must also perform the Endorsement function. [Should explain this one with examples.]

## Bootstrap
The initial setup of a network during which policies, system chaincode, and cryptographic materials (certs) are disseminated amongst Peers, and Orderers.  The bootstrap of the Ordering Network must precede the
bootstrap of the Peer Network, because a Peer network is contingent upon the presence
of an Ordering Service.  A network need only be “bootstrapped” once.

## Block
A batch of ordered transactions, potentially containing ones of an
invalid nature, that is delivered to the Peers for validation and Committal.

## System Chain 
Contains a configuration block defining the network at a system level.  
The system chain lives within the Ordering Service and — similar to a
channel — has an initial configuration containing: root certificates for participating organizations and ordering service nodes, policies [more on this?], listening address for OSN [link to OSN definition], and configuration details.  Any change to the overall network (e.g. a new org joining or a new OSN being added) will result
in a new configuration block being added to the system chain.  

The System Chain can be thought of as the common binding for a channel or group
of channels.  For instance, a collection of financial institutions may form a
consortium (represented through the system chain), and then proceed to create
channels relative to their various business agendas.  

## Channel
A Channel is formed as an offshoot of the system chain [need more clarity here]; and best thought of
as a “topic” for Peers to subscribe to, or rather, a subset of a
broader blockchain network [subledger?]. A Peer may subscribe on various channels and can
only access the transactions on the subscribed channels.  Each channel maintains
a unique ledger, thus accommodating confidentiality and execution of both bilateral and multilateral
agreements.

## Multi-channel
The fabric will allow for multiple channels with a
designated ledger per channel.  This capability allows for multilateral contracts
where only the restricted participants on the channel will submit, endorse,
order, or commit transactions on that channel.  As such, a single peer can
maintain multiple ledgers without compromising privacy and confidentiality. [Needs more here...probably an illustration.]

## Configuration Block
Contains the configuration data defining members and
policies for a system chain or channel(s).  Any changes to the channel(s) or
overall network (e.g. a new member successfully joining) will result in a new
configuration block being appended to the appropriate chain [chain or channel or just ledger?].  This block will
contain the contents of the genesis block, plus the delta.  The policy to alter
or edit a channel-level configuration block is defined through the Configuration
System Chaincode (CSCC).  

## Genesis Block
The configuration block that initializes a blockchain
network or channel, and also serves as the first block on a chain.  

## Ledger
An append-only transaction log managed by peers.  Ledger keeps the
log of ordered transaction batches.  There are two denotations for ledger; peer [not loving the notion of "peer ledger"]
and validated.  The peer ledger contains all batched transactions coming out of
the ordering service, some of which may in fact be invalid.  The validated ledger
will contain fully endorsed and validated transaction blocks.  In other words,
transactions in the validated ledger have passed the entire gamut of "consensus"
- i.e. they have been endorsed, ordered, and validated.

## Dynamic membership
The fabric will allow for Peers to
come and go [what does come and go mean?] based on membership, and the blockchain network will continue to
operate. Dynamic membership is critical when businesses grow and members need to
be added or removed for various reasons.

## Query/Non-Key Value Query
[should mention that couchDB is one data module...and you can plug in your own]  Using couchDB 2.0 you now have the capability
to leverage an API to perform more complex queries against combinations of
variables, including time ranges, transaction types, users, etc.  This feature
allows for auditors and regulators to aggregate and mine large chunks of data.

## Gossip Protocol
A communication protocol used among Peers in a channel, to
maintain their network and to elect Leaders. All
communications with the Ordering Service go through Gossip. [Isn't gossip a well known system we can point to docs for outside hyperledger?]  Gossip allows for data dissemination,
therein providing support for scalability due to the fact that not all peers
are required to execute transactions and communicate with the ordering service.

## System Chaincode
System Chaincode (SCC) is a chaincode built with the peer
and run in the same process as the peer.  SCC is responsible for broader
configurations of fabric behavior, such as timing and naming services.

## Lifecycle System Chaincode
Lifecycle System Chaincode (LSCC) is a system chaincode that handles deployment,
upgrade and termination transactions for user chaincodes.

## Configuration System Chaincode
Configuration System Chaincode (CSCC) is a “management” system chaincode that
handles configuration requests to alter an aspect of a channel
(e.g. add a new member).  The CSCC will interrogate the channel’s policies to
determine if a new configuration block can be created.

## Endorsement System Chaincode
Endorsement System Chaincode (ESCC) is a system chaincode that handles the
endorsement policy for specific pieces of chaincode deployed on a network.
It also defines the necessary parameters (percentage or combination of signatures
from endorsing peers) for a transaction proposal to receive a successful
proposal response (i.e. endorsement).  Deployments and invocations of user
chaincodes both require a corresponding ESCC, which is defined at the time of
the deployment transaction proposal for the user chaincode.

## Validation System Chaincode
Validation System Chaincode (VSCC) Handles the validation policy for
specific pieces of chaincode deployed on a network.  Deployments and invocations
of user chaincodes both require a corresponding VSCC, which is defined at the
time of the deployment transaction proposal for the user chaincode.  VSCC
validates the specified level of "endorsement" (i.e. endorsement policy) in
order to prevent malicious or faulty behavior from the client.

## Policy
There are policies for endorsement, validation, block committal,
chaincode management and network/channel management.  Policies are defined
through system chaincodes, and contain the requisite specifications for a
network action to succeed.  For example, an endorsement policy may require that
100% of endorsers achieve the same result upon transaction simulation.

## Endorsement policy
A blockchain network must establish rules that govern
the endorsement of proposed, simulated transactions. This endorsement
policy could require that a transaction be endorsed by a minimum number of
endorsing peers, a minimum percentage of endorsing peers, or by all endorsing
peers that are assigned to a specific chaincode application [we have not defined the notion of chaincode application]. Policies can be
curated based on the application and the desired level of resilience against
misbehavior (deliberate or not) by the endorsing peers.  A distinct endorsement
policy for deploy transactions, which install new chaincode, is also required. [Need patterns for some of these to get people going.]

## Proposal
A transaction request sent from a client or admin user to one or
more peers in a network; examples include deploy, invoke, query, or
configuration request.

## Deploy
Refers to the function through which chaincode applications are
deployed on `chain` [channel?].  A deploy is first sent from the client SDK or CLI to a
Lifecycle System Chaincode in the form of a proposal.

## Invoke
Used to call chaincode functions [agreements?].  Invocations are captured as
transaction proposals, which then pass through a modular flow of endorsement,
ordering, validation, committal.  The structure of invoke is a function and an
array of arguments.

## Membership Services
Membership Services manages user identities on a
permissioned blockchain network; this function is implemented through the
`fabric-ca` component.  `fabric-ca` is comprised of a client and server, and
handles the distribution and revocation of enrollment materials (certificates),
which serve to identify and authenticate [Members and] users on a network.

The in-line `MembershipSrvc` code (MSP) runs on the peers themselves, and is used by
the peer when authenticating transaction processing results, and by the client
to verify/authenticate transactions.  Membership Services provides a distinction
of roles by combining elements of Public Key Infrastructure (PKI) and
decentralization (consensus). By contrast, [permissionless] networks do not
provide member-specific authority or a distinction of roles.

A permissioned blockchain requires entities to register for long-term identity
credentials (Enrollment Certificates), which can be distinguished according to
entity type. For users, an Enrollment Certificate authorizes the Transaction
Certificate Authority (TCA) to issue pseudonymous credentials; these
certificates authorize transactions submitted by the user. Transaction
certificates persist on the blockchain, and enable authorized auditors to
associate, and identify the transacting parties for otherwise un-linkable
transactions. [This is a big idea and should get an illustration of its own.]

## Membership Service Provider
The Membership Service Provider (MSP) refers to an abstract component of the
system that provides (anonymous) credentials to clients, and peers for them to
participate in a Hyperledger/fabric network. Clients use these credentials to
authenticate their transactions, and peers use these credentials to authenticate
transaction processing results (endorsements). While strongly connected to the
transaction processing components of the systems, this interface aims to have
membership services components defined, in such a way that alternate
implementations of this can be smoothly plugged in without modifying the core of
transaction processing components of the system.

## Initialize
A chaincode method to define the assets and parameters in a
piece of chaincode prior to issuing deploys and invocations.  As the name
implies, this function should be used to do any initialization to the chaincode,
such as configure the initial state of a key/value pair on the ledger.

## appshim [this needs clarification]
An application client used by ordering service nodes to
process "broadcast" messages arriving from clients or peers.  This shim allows
the ordering service to perform membership-related functionality checks.  In
other words, is a peer or client properly authorized to perform the requested
function (e.g. upgrade chaincode or reconfigure channel settings).

## osshim
An ordering service client used by the application to process
ordering service messages (i.e. "deliver" messages) that are advertised
within a channel.  

## Hyperledger Fabric Client SDK
Provides a powerful set of APIs and contains “methods” or “calls”
that expose the capabilities and functionality in the Hyperledger Fabric
code base.  For example, `addMember`, `removeMember`.  The Fabric SDK
comes in multiple [languages] - Node.js, Java, and Python, and others will follow -
thus, allowing developers to write application code in any of those
programming languages.  

## Chaincode
Embedded logic that encodes the rules for specific types of
network transactions. Developers write chaincode applications, which are then
deployed onto a chain by an appropriately authorized member. End users then
invoke chaincode through a client-side application that interfaces with a
network peer. Chaincode runs network transactions, which if validated, are
appended to the shared ledger and modify world state. [Need to make clear that "chaincode" is just logic written in a specific programming language and deployed essentially as a "stored procedure" on a ledger...either a system ledger or a channel ledger.]
