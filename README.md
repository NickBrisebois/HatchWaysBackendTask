#Blog Posts API
## Project for HatchWays job application

### Install dependencies
Use go mod (requires Go 1.13)

```
go mod vendor
```

### Building
Use makefile to build

```
make
```

### Running
Call program with path to config.toml

```
./build/hwApi --config "./build/config.toml"
```