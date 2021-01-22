# contributions

[![Go Report Card](https://goreportcard.com/badge/github.com/oke-py/contributions)](https://goreportcard.com/report/github.com/oke-py/contributions)
[![Coverage Status](https://coveralls.io/repos/github/oke-py/contributions/badge.svg?branch=main)](https://coveralls.io/github/oke-py/contributions?branch=main)

count up GitHub contributions by repositories

## Usage

Generate a token from https://github.com/settings/tokens.

```
docker run -e GITHUB_TOKEN=<YOUR GITHUB TOKEN> okepy/contribution -u <YOUR GITHUB ACCOUNT> -m 2020/1
```
