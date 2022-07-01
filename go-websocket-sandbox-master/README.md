# websocket-sandbox / Clean Architecture 


## TODO

- [ ] Add Parse JWT 
- [ ] Add CI setting 
- [ ] Add Dokcer setting
- [ ] Add error tracker (SENTRY Go)
- [ ] Implement notification response (REST)
- [ ] Implement Cloud Server Setting

## Requirements

### Go

```bash
1.11.* 
```

### Insrall goenv and change Go version.

```bash
brew install goenv
goenv install 1.11.6
goenv global 1.11.6
goenv rehash
```

## Server Config

copy `config.yaml.example` to `config.yaml` and edit params.

## Commands

### Test

```bash
make test
```

### Build

```bash
make build
```

### Serve

```bash
make serve
```
