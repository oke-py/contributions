# contributions

[![Go Report Card](https://goreportcard.com/badge/github.com/oke-py/contributions)](https://goreportcard.com/report/github.com/oke-py/contributions)
[![Coverage Status](https://coveralls.io/repos/github/oke-py/contributions/badge.svg?branch=main)](https://coveralls.io/github/oke-py/contributions?branch=main)

count up GitHub contributions by repositories

## Usage

Generate a token from https://github.com/settings/tokens.

When using a Fine-grained token, you only need to grant "Events: Read-only" permission in Account Permissions. No repository permissions are required.

```
docker run -e GITHUB_TOKEN=<YOUR GITHUB TOKEN> okepy/contribution -u <YOUR GITHUB ACCOUNT> -m 2020/1
```

## Sample

```
docker run -e GITHUB_TOKEN=<MY GITHUB TOKEN> okepy/contribution -u oke-py -m 2021/2
| Repository                     | Commits | Issues  | PRs     | Reviews |
| ------------------------------ | ------: | ------: | ------: | ------: |
| aquasecurity/kube-bench        |       1 |       0 |       1 |       0 |
| kubernetes/enhancements        |       1 |       0 |       1 |       0 |
| kubernetes/website             |       8 |       5 |       8 |       2 |
| oke-py/issue-creator           |       1 |       0 |       0 |       0 |
| open-policy-agent/conftest     |       0 |       2 |       0 |       0 |
```
