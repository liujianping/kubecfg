kubecfg
===
kubecfg, more convenient tool for kube config

<p align="center"><img src="/image/kubecfg.gif?raw=true"/></p>

## Install

### Shell Install support Linux & MacOS

````bash
# binary will be $(go env GOPATH)/bin/ts
$: curl -sfL https://raw.githubusercontent.com/liujianping/kubecfg/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# In alpine linux (as it does not come with curl by default)
$: wget -O - -q https://raw.githubusercontent.com/liujianping/kubecfg/master/install.sh | sh -s 

```` 

### Brew Install only MacOS
````bash
$: brew tap liujianping/tap && brew install kubecfg
````

### Source Install
````bash
$: git clone git@github.com:liujianping/kubecfg.git
$: cd ts
$: go install 
````

## Quick Start

````bash
$: kubecfg -h
kubecfg, more convenient tool for kube config

Usage:
  kubecfg [command]

Available Commands:
  cat         view ~/.kube/config
  cd          change context
  help        Help about any command
  ls          list all contexts
  merge       merge a new kube config
  ncd         change current context namespace
  rename      rename context
  rm          remove context

Flags:
  -h, --help      help for kubecfg
      --version   version for kubecfg
````
