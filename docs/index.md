# Incubation Notice

This project is a Hyperledger project in _Incubation_. It was proposed to the
community and documented [here](https://goo.gl/RYQZ5N). Information on what
_Incubation_ entails can be found in the [Hyperledger Project Lifecycle
document](https://goo.gl/4edNRc).

[![Build Status](https://jenkins.hyperledger.org/buildStatus/icon?job=fabric-merge-x86_64)](https://jenkins.hyperledger.org/view/fabric/job/fabric-merge-x86_64/)
[![Go Report Card](https://goreportcard.com/badge/github.com/hyperledger/fabric)](https://goreportcard.com/report/github.com/hyperledger/fabric)
[![GoDoc](https://godoc.org/github.com/hyperledger/fabric?status.svg)](https://godoc.org/github.com/hyperledger/fabric)
[![Documentation Status](https://readthedocs.org/projects/hyperledger-fabric/badge/?version=latest)](http://hyperledger-fabric.readthedocs.io/en/latest/?badge=latest)

# Hyperledger Fabric

The Fabric is an implementation of blockchain technology, leveraging familiar
and proven technologies. It is a modular architecture allowing pluggable
implementations of various function. It features powerful container technology
to host any mainstream language for smart contracts development.

## Releases

The Fabric releases are documented [here](releases.md). We
released our second release under the governance of the Hyperledger Project -
v0.6-preview in October, 2016.

If you are seeking a stable version of the Hyperledger Fabric on which to
develop applications or explore the technology, we **strongly recommend** that
you use the v0.6 [Starter Kit](http://hyperledger-fabric.readthedocs.io/en/v0.6/starter/fabric-starter-kit/)
while the v1.0 architectural refactoring is in development.

If you'd like a taste of what the Fabric v1.0 architecture has in store, we
invite you to read the preview [here](abstract_v1.md). Finally, if you are
adventurous, we invite you to explore the current state of development with
the caveat that it is not yet completely stable.

## Contributing to the project

We welcome contributions to the Hyperledger Project in many forms. There's
always plenty to do! Full details of how to contribute to this project are
documented in the [Fabric developer's guide](#fabric-developer-guide) below.

## Maintainers

The project's [maintainers](MAINTAINERS.md) are responsible for reviewing and
merging all patches submitted for review and they guide the over-all technical
direction of the project within the guidelines established by the Hyperledger
Project's Technical Steering Committee (TSC).

## Communication <a name="communication"></a>

We use [Hyperledger Slack](https://slack.hyperledger.org/) for communication and
Google Hangouts&trade; for screen sharing between developers. Our development
planning and prioritization is done in [JIRA](https://jira.hyperledger.org),
and we take longer running discussions/decisions to the
[mailing list](http://lists.hyperledger.org/mailman/listinfo/hyperledger-fabric).

## Still Have Questions?
We try to maintain a comprehensive set of documentation (see below) for various
audiences. However, we realize that often there are questions that remain
unanswered. For any technical questions relating to the Hyperledger Fabric
project not answered in this documentation, please use
[StackOverflow](http://stackoverflow.com/questions/tagged/hyperledger). If you
need help finding things, please don't hesitate to send a note to the
[mailing list](http://lists.hyperledger.org/mailman/listinfo/hyperledger-fabric),
or ask on [Slack]((https://slack.hyperledger.org/)).

# Hyperledger Fabric Documentation

The Hyperledger Fabric is an implementation of blockchain technology, that has
been collaboratively developed under the Linux Foundation's
[Hyperledger Project](http://hyperledger.org). It leverages familiar and
proven technologies, and offers a modular architecture
that allows pluggable implementations of various function including membership
services, consensus, and smart contracts (Chaincode) execution. It features
powerful container technology to host any mainstream language for smart
contracts development.


## Read all about it

If you are new to the project, you might want to familiarize yourself with some
of the basics before diving into either developing applications using the
Hyperledger Fabric, or collaborating with us on our journey to continuously
extend and improve its capabilities.

- [Canonical use cases](biz/usecases.md)
- [Glossary](glossary.md): to understand the terminology that we use throughout
the Fabric project's documentation.
- [Fabric FAQs](https://github.com/hyperledger/fabric/tree/master/docs/FAQ)


**Note:** if you are looking for instructions to operate the fabric for a POC,
we recommend that you use the more stable v0.6 [Starter Kit](http://hyperledger-fabric.readthedocs.io/en/v0.6/starter/fabric-starter-kit/)

# License <a name="license"></a>
The Hyperledger Project uses the [Apache License Version 2.0](LICENSE) software
license.
