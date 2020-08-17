# conntest

cli tool for testing network connectivity based on a config file

## Install

```bash
git clone https://github.com/rmorris1218/conntest.git && cd conntest
go build -o conntest
./conntest --help
```

## Usage

To use the tool, you need a config file with the following format

*./conntest-config.json*
```json
{
    "tests": [
        {
            "endpoint": "1.1.1.1",
            "proto_ports": [
                {"port": 53, "protocol": "tcp"}
            ]
        }
    ]
}
```

Then to run the tool

```
$ ./conntest --config-file ./conntest-config.json
successfully reached all ports on 1.1.1.1
```

If all endpoints are reachable, there will be an exit code of 0. If there are any errors, there will be an exit code of 1.

