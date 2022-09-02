---
order: 2
---

# Installation

## Install `go`

::: tip
**Go 1.18** is recommended for building and installing the Spartan-Cosmos software.
:::

Install `go` by following the [official docs](https://go.dev/doc/install).

Remember to set your `$GOPATH`, `$GOBIN`, and `$PATH` environment variables, for example:

```bash
mkdir -p $HOME/go/bin
echo "export GOPATH=$HOME/go" >> ~/.bashrc
source ~/.bashrc
echo "export GOBIN=$GOPATH/bin" >> ~/.bashrc
source ~/.bashrc
echo "export PATH=$PATH:$GOBIN" >> ~/.bashrc
source ~/.bashrc
```

Verify that `go` has been installed successfully.

```bash
go version
```

## Install `spartan`

After setting up `go` correctly, you should be able to compile and run `spartan`.

Make sure that your server can access Google because our project depends on some libraries provided by Google. (If you don't have access to google.com, you can also try to add a proxy: `export GOPROXY=https://goproxy.io`) 

```bash
git clone https://github.com/bianjieai/spartan-cosmos.git
cd spartan-cosmos
git checkout <version>
make install
```

Have you set up the Environment Variables correctly, you will get no error during `spartan` installation.

Now check your `spartan` version.

```bash
spartan version
```
