# About 

Simple Project to create passwords via a cli tool in go

# Installation

to install the application use:
```
go installl
```
## To use

```
pgen -l 100 -s
```

The command line arguments are

`-l` used to specify the length of the password [Optional] 
`-s` Use Special characters [Optional]

## To remove

uninstall using Linuc command

```shell
rm $(which pgen)
hash -r
```

The first command removes the executable from where it was copied.
The second command clears Shell Cache

# 3rd Party libraries

uses `https://github.com/urfave/cli` is a "declarative, simple, fast, and fun package for building command line tools in Go"

Set up the project using

```shell
go get github.com/urfave/cli/v2@latest
```
