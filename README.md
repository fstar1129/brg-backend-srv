# bridge-backend-service

The service is charge of *slashing*, *penalty* and *rewarding* and *relayers registration* functionality of cross-chain swap.

There is only *relayer registeration* functionality in current version.

I recommend to compare this project with relayer-smart-contrats.

# TODO

- Update `config.json` with your configuration.
- Update file `./service/unregister-relayer.go` by analogy with the file `./service/register-relayer.go`.
- Update all files with **!!! TODO !!!**:
    - `./service/relayer.go`
    - `./service/workers/eth-compatible/worker.go` method *Unregister()* also check if *Unregister()* contains in worker interface `./service/workers/worker.go`
    - add in storage new types of events. Go to `./service/storage/types.go` and do the same as with `REGISTER`.
    - update `./service/storage/tx.go` method *ConfirmTx()*