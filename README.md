# oracle-go
**Golang NodeDAO Oracle daemon**

Made by:  [HashKing](https://www.hashking.fi)

![HashKing](./docs/images/HashKingLogo.svg)

Oracle daemon for NodeDAO decentralized staking service. Collects and reports Beacon Chain states to the NodeDAO contract running on Ethereum execution layer.

For oracle contracts, see: [NodeDAO/NodeDAO-Protocol](https://github.com/NodeDAO/NodeDAO-Protocol)



## Config File

See the complete configuration file：[conf/config-default.yaml](conf/config-default.yaml)

You are advised to modify the configuration：

```yaml
log:
  level:
    server: info
  # console or json
  format: console

eth:
  network: mainnet
  # ETH execution layer connection address
  elAddr:
  # ETH consensus layer connection address
  clAddr:
  # Oracle member's private key（Make enough ETH to support the GAS of transactions; without ‘0x’ prefix）
  privateKey:
```

### Configuration details

| Attribute        | Required | Default | Description                                                                                          |
|:-----------------|----------|---------|------------------------------------------------------------------------------------------------------|
| log.level.server | false    | info    | log level                                                                                            |
| log.format       | false    | console | log format: console or json                                                                                         |
| eth.network      | false    | mainnet | network                                                                                              |
| eth.clAddr       | true     |         | ETH consensus layer connection address                                                               |
| eth.elAddr       | true     |         | ETH execution layer connection address                                                               |
| eth.privateKey   | true     |         | Oracle member's private key（Make enough ETH to support the GAS of transactions; without ‘0x’ prefix） |

### Note:

- Ensure connectivity between the ETH execution layer and the ETH consensus layer nodes for which the `elAddr` and `clAddr` parameters are set.
- Ensure that privateKey's private key is authorized by Oracle member(privateKey without ‘0x’ prefix).



## Setup

### Quick Commands
```shell
./cli-init.sh
./oracle-go oracle withdraw --config <YOUR_CONFIG_FILE>
```

### Docker

[NodeDAO oracle-go docker images](https://hub.docker.com/r/nodedao/oracle-go)

```sh
docker pull nodedao/oracle-go:latest
docker run --name nodedao-oracle -v `pwd`/config-dev.yaml:/config/config-dev.yaml -d --restart=always nodedao/oracle-go:latest --config /config/config-dev.yaml
```



# License

2022 kinghash

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, version 3 of the License, or any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the [GNU General Public License](https://github.com/NodeDAO/oracle-go/blob/main/LICENSE) along with this program. If not, see https://www.gnu.org/licenses/.
