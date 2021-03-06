[![Build Status](https://travis-ci.org/HeavyHorst/easyKV.svg?branch=master)](https://travis-ci.org/HeavyHorst/easyKV) [![Go Report Card](https://goreportcard.com/badge/github.com/HeavyHorst/easyKV)](https://goreportcard.com/report/github.com/HeavyHorst/easyKV) [![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/HeavyHorst/easyKV/master/LICENCE)
[![](https://godoc.org/github.com/HeavyHorst/easyKV?status.svg)](http://godoc.org/github.com/HeavyHorst/easyKV) [![codecov](https://codecov.io/gh/HeavyHorst/easyKV/branch/master/graph/badge.svg)](https://codecov.io/gh/HeavyHorst/easyKV)


# easyKV
easyKV is based on the backends of [confd](https://github.com/kelseyhightower/confd).
easyKV provides a very simple interface to work with some key-value stores.
The goal of easyKV is to abstract these 2 common operations for multiple backends:

  - recursively query the kv-store for key-value pairs.
  - watch a key-prefix for changes.

## Interface
A **storage backend** in `easyKV` should implement (fully or partially) this interface:
```go
type ReadWatcher interface {
	GetValues(keys []string) (map[string]string, error)
	WatchPrefix(ctx context.Context, prefix string, opts ...WatchOption) (uint64, error)
	Close()
}
```

## Compatibility matrix

| Calls                 |   Consul   | Etcdv2 | Etcdv3  |  env  | file |   redis |  vault  |  zookeeper |
|-----------------------|:----------:|:------:|:-------:|:-----:|:----:|:-------:|:-------:|:----------:|
| GetValues             |     X      |   X    |      X  |    X  |  X   |     X   |   X     |     X      |
| WatchPrefix           |     X      |   X    |      X  |       |  X   |         |         |     X      |
| Close                 |     X      |   X    |      X  |    X  |  X   |     X   |   X     |     X      |
