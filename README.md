# oracle-go

Made by:  [kinghash](https://www.kinghash.com/)

![kinghash](./docs/images/kingHashLogo.PNG)

**NodeDAO Oracle V2 daemon**

Oracle daemon for NodeDAO decentralized staking service. Collects and reports Beacon Chain states (the number of visible validators and their summarized balances) to the NodeDAO dApp contract running on Ethereum execution layer.

For oracle contracts, see: [NodeDAO/NodeDAO-Protocol](https://github.com/NodeDAO/NodeDAO-Protocol)

## Working Mechanism
TODO

## Config File

config-default.json

```yaml
server:
  name: oracle-go
  version: 0.0.1-beta

cli:
  name: oracle-go

log:
  level:
    server: debug

eth:
  network: goerli
  elAddr:
  clAddr:
  privateKey:
```

## exec oracle
```shell
./oracle-go oracle withdraw
```

## Setup
### Quick Commands
TODO

### Docker

[NodeDAO oracle-go docker images](https://hub.docker.com/r/nodedao/oracle-go)

```sh
docker pull nodedao/oracle-go
docker run -v `pwd`/config-dev.yaml:/config/config-dev.yaml -d  nodedao/oracle-go:latest --config /config/config-dev.yaml
```

# License

2022 kinghash

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, version 3 of the License, or any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the [GNU General Public License](https://github.com/NodeDAO/oracle-go/blob/main/LICENSE) along with this program. If not, see https://www.gnu.org/licenses/.
