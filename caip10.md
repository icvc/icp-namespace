---
namespace-identifier: icp-caip10
title: ICP Namespace - Chains
author: Quint Daenen (@q-uint)
discussions-to: TODO
status: Draft
type: Standard
created: 2023-05-24
requires: [ CAIP-2, CAIP-10 ]
---

# CAIP-10

*For context, see the [CAIP-10][] specification.*

## Rationale

In the world of blockchain technology, the need to identify and differentiate various entities within a decentralized
network is crucial. CAIP-10 addresses this challenge by providing a standardized method to identify an account in any
blockchain specified by the CAIP-2 blockchain ID.

On the Internet Computer the concept of identification takes the form of principals. Principals serve as generic
identifiers for users, smart contracts known as canisters, and potentially other entities in the future. Principals play
a similar role to addresses on Ethereum Virtual Machine (EVM) chains, where they can refer to contracts and users ids.

In the context of the Internet Computer, principals are considered opaque binary blobs, typically ranging in length from
0 to 29 bytes. Importantly, there intentionally exists no mechanism to distinguish between canister IDs and user IDs.
This design choice allows for flexibility and future extensibility, enabling the Internet Computer to accommodate new
concepts and entities as they emerge.

## Syntax

There are different forms of principals, but this is out of scope for this specification. The specification will only
focus on the textual representation of a principal.

The textual representation of a `principal p` is `Text(Base32(CRC32(p) · p))` where

- `CRC32` is a four byte check sequence, calculated as defined by ISO 3309 and stored as big-endian.

- `Base32` is the Base32 encoding as defined in [RFC 4648](https://tools.ietf.org/html/rfc4648#section-6), with no
  padding character added.

- The middle dot (`.`) denotes concatenation.

- `Text` takes an ASCII string and inserts the separator `-` (dash) every 5 characters. The last group may contain
  less than 5 characters. A separator never appears at the beginning or end.

The textual representation is conventionally printed with lower case letters, but parsed case-insensitively.

Because the maximum size of a principal is 29 bytes, the textual representation will be no longer than 63 characters (10
times 5 plus 3 characters with 10 separators in between them).

```shell
function textual_encode() {
  ( echo "$1" | xxd -r -p | /usr/bin/crc32 /dev/stdin; echo -n "$1" ) |
  xxd -r -p | base32 | tr A-Z a-z |
  tr -d = | fold -w5 | paste -sd'-' -
}

function textual_decode() {
  echo -n "$1" | tr -d - | tr a-z A-Z |
  fold -w 8 | xargs -n1 printf '%-8s' | tr ' ' = |
  base32 -d | xxd -p | tr -d '\n' | cut -b9- | tr a-z A-Z
}
```

The `account_id` is a case-sensitive string in the form

```
account_id:        chain_id + ":" + account_address
chain_id:          [-a-z0-9]{3,8}:[-_a-zA-Z0-9]{1,32} (See [CAIP-2][])
account_address:   [-.%a-zA-Z0-9]{1,128}
```

## Test Cases

```
# Zero Principal
icp:737ba355e855bd4b61279056603e0550:aaaaa-aa

# Anonymous Principal
icp:737ba355e855bd4b61279056603e0550:2vxsx-fae

# User Principal
icp:737ba355e855bd4b61279056603e0550:g27xm-fnyhk-uu73a-njpqd-hec7y-syhwe-bd45b-qm6yc-xikg5-cylqt-iae

# Canister Principal (ICP Ledger)
icp:737ba355e855bd4b61279056603e0550:ryjl3-tyaaa-aaaaa-aaaba-cai
```

## Considerations

### Account Ids

Within the ledger, an account is identified by its address, which is derived from the principal ID and the sub-account
identifier. For better understanding, you can compare principal identifiers to the hash of a user's public key in
systems like Bitcoin or Ethereum. To authenticate and perform operations on the principal's account in the ledger
canister, the corresponding secret key is used to sign messages. Additionally, canisters themselves can possess accounts
within the ledger canister, and in such cases, the address is derived from the principal of the canister.

It is possible for an account owner to have authority over more than one account. In such cases, each account
corresponds to a pair consisting of an account owner and a sub-account. The sub-account is an optional bitstring that
serves the purpose of distinguishing between different sub-accounts belonging to the same owner.

## References

- [CAIP-2](https://github.com/ChainAgnostic/CAIPs/blob/master/CAIPs/caip-2.md)
- [CAIP-10](https://github.com/ChainAgnostic/CAIPs/blob/master/CAIPs/caip-10.md)
- [Principal IDs](https://internetcomputer.org/docs/current/references/ic-interface-spec#principal)

## Rights

Copyright and related rights waived via CC0.
