# oracle-go
**Golang NodeDAO Oracle daemon**

Made by:  [kinghash](https://www.kinghash.com/)

![kinghash](./docs/images/kingHashLogo.PNG)

Oracle daemon for NodeDAO decentralized staking service. Collects and reports Beacon Chain states to the NodeDAO contract running on Ethereum execution layer.

For oracle contracts, see: [NodeDAO/NodeDAO-Protocol](https://github.com/NodeDAO/NodeDAO-Protocol)



## Config File

See the complete configuration file：[conf/config-default.yaml](conf/config-default.yaml)

You are advised to modify the configuration：

```yaml
log:
  level:
    server: debug

eth:
  network: mainnet
  elAddr:
  clAddr:
  privateKey:
```

### Configuration details

| Attribute        | Required | Default | Description                                                  |
| ---------------- | -------- | ------- | ------------------------------------------------------------ |
| log.level.server | false    | info    | log level                                                    |
| network          | false    | mainnet | network                                                      |
| clAddr           | true     |         | ETH consensus layer connection address                       |
| elAddr           | true     |         | ETH execution layer connection address                       |
| privateKey       | true     |         | Oracle member's private key（Make enough ETH to support the GAS of transactions） |

### Note:

- Ensure connectivity between the ETH execution layer and the ETH consensus layer nodes for which the `executionLayerAddr` and `beaconNodeAddr` parameters are set.
- Ensure that privateKey's private key is authorized by Oracle member.



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
docker run -v `pwd`/config-dev.yaml:/config/config-dev.yaml -d  nodedao/oracle-go:latest --config /config/config-dev.yaml
```



# License

2022 kinghash

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, version 3 of the License, or any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the [GNU General Public License](https://github.com/NodeDAO/oracle-go/blob/main/LICENSE) along with this program. If not, see https://www.gnu.org/licenses/.
