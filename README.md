# jsgql

Extract GraphQL mutation and query names from files.

## Install

```bash
go install github.com/0xdead4f/jsgql@latest
```

## Usage

1. From stdin:

```bash
cat example.js | jsgql
```

2. From file:

```bash
jsgql example.js
```

## Examples

**Output from stdin:**

```bash
$ cat example.js | jsgql
mutation createUser($input: UserInput)
query getUser($id: ID!)
```

**Output from file:**

```bash
$ jsgql example.js
{"name":"mutation createUser($input: UserInput){ createUser()}","file":"example.js"}
{"name":"query getUser($id: ID!){ getUser()}","file":"example.js"}
```
