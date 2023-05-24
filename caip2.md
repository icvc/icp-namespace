---
namespace-identifier: icp-caip2
title: ICP Namespace - Chains
author: Quint Daenen (@q-uint)
discussions-to: TODO
status: Draft
type: Standard
created: 2023-05-17
requires: CAIP-2
---

# CAIP-2

*For context, see the [CAIP-2][] specification.*

## Rationale

The CAIP-2 namespace for the Internet Computer Protocol (ICP) provides a standardized system for identifying assets
within the network. It promotes interoperability, efficient querying, and clarity in asset identification, benefiting
developers and users working with decentralized applications and assets on the Internet Computer.

## Syntax

The ICP *namespace* is `icp` and the *reference* is the first 32 characters of hexadecimal representation of the SHA-256
hash of the root key of the ICP instance.

### NameSpace

```text
icp
```

### Reference

```text
737ba355e855bd4b61279056603e0550
```

### ChainId

```text
icp:737ba355e855bd4b61279056603e0550
```

## Resolution Method

You can resolve the root key of the ICP by querying the status endpoint of the ICP API.
The response is CBOR encoded and contains the root key.

```shell
curl "https://icp0.io/api/v2/status" --output - | hexdump -C
```

Once the root key is obtained, the reference can be calculated by taking the first 32 characters of hexadecimal
representation of the SHA-256 hash of the root key.

```shell
rootKey=$(echo "308182301d060d2b0601040182dc7c0503010201060c2b0601040182dc7c05030201036100814c0e6ec71fab583b08bd81373c255c3c371b2e84863c98a4f1e08b74235d14fb5d9c0cd546d9685f913a0c0b2cc5341583bf4b4392e467db96d65b9bb4cb717112f8472e0d5a4d14505ffd7484b01291091c5f87b98883463f98091a0baaae" | xxd -r -p)
hash=$(echo -n "$rootKey" | sha256sum | cut -d ' ' -f 1 | cut -c 1-32)
echo "icp:$hash"
```

The `chain_id` is a case-sensitive string in the form

```
chain_id:    namespace + ":" + reference
namespace:   [-a-z0-9]{3,8}
reference:   [-_a-zA-Z0-9]{1,32}
```

## Examples

### JS

```js
const {createHash} = require('crypto');

const rootKey = Buffer.from(
    '308182301d060d2b0601040182dc7c0503010201060c2b0601040182dc7c05030201036100814' +
    'c0e6ec71fab583b08bd81373c255c3c371b2e84863c98a4f1e08b74235d14fb5d9c0cd546d968' +
    '5f913a0c0b2cc5341583bf4b4392e467db96d65b9bb4cb717112f8472e0d5a4d14505ffd7484' +
    'b01291091c5f87b98883463f98091a0baaae', "hex"
);

(async () => {
    const hash = await createHash('sha256').update(rootKey).digest('hex');
    console.log("icp:" + hash.slice(0, 32))
})()
```

### Golang

```go
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"

	"github.com/aviate-labs/agent-go"
)

var icp0, _ = url.Parse("https://icp0.io/")

func main() {
	c := agent.NewClient(agent.ClientConfig{Host: icp0})
	status, _ := c.Status()
	hash := sha256.Sum256(status.RootKey)
	hashHexString := hex.EncodeToString(hash[:])
	fmt.Printf("icp:%s", hashHexString[:32])
}
```

## Test Cases

```text
# icp0.io (root_key = "308182301d060d2b0601040182dc7c0503010201060c2b0601040182dc7c05030201036100814c0e6ec71fab583b08bd81373c255c3c371b2e84863c98a4f1e08b74235d14fb5d9c0cd546d9685f913a0c0b2cc5341583bf4b4392e467db96d65b9bb4cb717112f8472e0d5a4d14505ffd7484b01291091c5f87b98883463f98091a0baaae")
icp:737ba355e855bd4b61279056603e0550
```

## Considerations

### Evolving Network Structure

While there is currently one "mainnet" running the Internet Computer, it is important to acknowledge that the network's
structure may change in the future. To accommodate potential network variations, the CAIP-2 spec recommends using the
hash of the root key as the reference, rather than solely relying on the term "mainnet." This approach allows for
flexibility and adaptation as the Internet Computer continues to evolve and potentially introduces new network
configurations.

## References

- [CAIP-2](https://github.com/ChainAgnostic/CAIPs/blob/master/CAIPs/caip-2.md)
- [ICP API Status](https://internetcomputer.org/docs/current/references/ic-interface-spec#api-status)
- [ICP Namespace - Reference Implementations](https://github.com/icvc/icp-namespace)

## Rights

Copyright and related rights waived via CC0.
