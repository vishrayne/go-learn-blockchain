# go-learn-blockchain
[![baby-gopher](https://raw.githubusercontent.com/drnic/babygopher-site/gh-pages/images/babygopher-logo-small.png)](http://www.babygopher.org)

Attempt to learn how blockchain works by following [this](https://jeiwan.cc/posts/building-blockchain-in-go-part-1) guide.

# Run
1. Install dep, if not, and run
```
$ dep ensure
```
2. `$ go build github.com/vishrayne/go-learn-blockchain/cmd/blockchain-cli`
3. `$ go install github.com/vishrayne/go-learn-blockchain/cmd/blockchain-cli`
4. Run! 
```
$ blockchain-cli help

    Usage:
        addblock -data <BLOCK_DATA> - add the given block to the blockchain
        printchain - print all the blocks of the chain

$ blockchain-cli addblock -data "First transaction data"

    Mining the block containing "Genesis block"
    00000090785688fad3411423c11e084d03fbfba497d39f824a3adbb3dc2aaa95

    Mining the block containing "First transaction data"
    0000007056aa6db982634a0d4cea379a7741d06c3575472eb858a0e9df900855

    Success!

$ blockchain-cli addblock -data "second transaction data"

    Mining the block containing "second transaction data"
    00000071bce7b25ce51cb13c3f744050803acef28ffc6bbc76b1fb273042b86e

    Success!

$ blockchain-cli printchain

    Prev block hash: 0000007056aa6db982634a0d4cea379a7741d06c3575472eb858a0e9df900855
    Data: second transaction data
    Hash: 00000071bce7b25ce51cb13c3f744050803acef28ffc6bbc76b1fb273042b86e
    PoW: true

    Prev block hash: 00000090785688fad3411423c11e084d03fbfba497d39f824a3adbb3dc2aaa95
    Data: First transaction data
    Hash: 0000007056aa6db982634a0d4cea379a7741d06c3575472eb858a0e9df900855
    PoW: true

    Prev block hash: 
    Data: Genesis block
    Hash: 00000090785688fad3411423c11e084d03fbfba497d39f824a3adbb3dc2aaa95
    PoW: true
```

## Progress/TODO
- [x] Basic prototype
- [x] Proof of work
- [x] Persistence
- [X] CLI
- [ ] Transactions (Part I)
- [ ] Addresses
- [ ] Transactions (Part II)
- [ ] Network


