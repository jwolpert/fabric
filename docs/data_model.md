# Transaction Data Model

[wip]

The current implementation of Hyperledger Fabric's transaction data model
utilizes read-write sets, coupled with the added protection of multi-version
concurrency control to counter against threats such as double-spend.

There are three main features to the model - MVCC pre-image, post-image, validation:

1. Transactions arrive in the form of a proposal to the endorsing peers.  The
endorsing peers simulate the transaction using a set of inputs (i.e. the read set).  
The read set contains unique keys and the committed versions of these keys,
which exist in the stateDB.  In other words, the read set or "pre-image" is a
representation of the current world state for those specified keys.
[__Note__: the peers can only read the value for the committed state of a key.  
They cannot read the results of a write operation during the
simulation/endorsement phase.  This reliance on committed versioning is the
foundation for data and transactional integrity.]

2. The ensuing result of the transaction simulation is the write set or the
"post-image." It contains the keys and their updated values.
(No version specified here).  The peers encode the contents (R/W set), sign the
message, and deliver the payload to an SDK.  All endorsing peers for a specified
chaincode will perform this operation.  The SDK aggregates the responses, and
the tx is passed along to the ordering service if the endorsement policy is fulfilled.

3. Transactions are then validated prior to being committed to the ledger.  A
validation system chaincode will ensure that the committed versions of the keys
are the same as they were during the simulation phase.  If the current versions
match the MVCC pre-image, then the transaction is committed and the world state
is updated (i.e. a block is appended and the stateDB is updated with the keys'
new versions).

The validation phase prevents the same key from being modified more than once
in the same block.  Let's look at a very simplistic example to see how the MVCC pre-image
& post-image approach is implemented.  Assume that the version of a key `k` is
represented in the world state (stateDB) by a tuple `(k,ver,val)` where `ver` is
the latest version of the key `k` with a value of `val`.

Now, consider a set of five transactions `T1, T2, T3, T4, T5`, all simulated on the same
snapshot of the world state.  The following snippet shows the current snapshot
of the world state and the read/write operations performed by each transaction:

```
World state: (k1,1,v1), (k2,1,v2), (k3,1,v3), (k4,1,v4), (k5,1,v5)
T1 -> Write(k1, v1'), Write(k2, v2')
T2 -> Read(k1), Write(k3, v3')
T3 -> Write(k2, v2'')
T4 -> Write(k2, v2'''), read(k2)
T5 -> Write(k6, v6'), read(k5)
```
Assume these transactions are ordered in the sequence of T1,..,T5. They could be in the
same block or in different blocks.  The important takeaway is that they were all
simulated against the same world state.

1. `T1` passes the validation because it does not perform any read. Further, the
tuple of keys `k1` and `k2` in the world state are updated to `(k1,2,v1'), (k2,2,v2')`

2. `T2` fails the validation because it reads a key `k1` which is modified by a preceding transaction `T1`

3. `T3` passes the validation because it does not perform a read. Further the
tuple of the key `k2` in the world state are updated to `(k2,3,v2'')`

4. `T4` fails the validation because it reads a key `k2` which is modified by a preceding transaction `T1`

5. `T5` passes the validation because it reads a key `k5` which is not modified
by any of the preceding transactions

