# contributions

[![Go Report Card](https://goreportcard.com/badge/github.com/oke-py/contributions)](https://goreportcard.com/report/github.com/oke-py/contributions)
[![Coverage Status](https://coveralls.io/repos/github/oke-py/contributions/badge.svg?branch=master)](https://coveralls.io/github/oke-py/contributions?branch=master)

count up GitHub contributions by repositories

## Usage

### build

```
git clone https://github.com/oke-py/contributions.git
cd contributions
make build
```

### export GITHUB_TOKEN

Generate a new token from https://github.com/settings/tokens and export it.

```
export GITHUB_TOKEN=<YOUR GITHUB TOKEN>
```

### run

```
./bin/contribution -u <GITHUB ACCOUNT> -m 2020/1
```
