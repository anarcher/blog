+++
title="Releasing with bumpversion, govvv and drone"
tags=["tool","deployment"]
draft=false
date="2017-12-05"
+++

One of pleasures about coding is using good tools. Recently I use [bumpversion][], [govvv][] and [drone][] for version releasing.  

## bumpversion

bumpversion is automation for semantic versioning. Most of my projects have simple config file like below. (.bumpversion.cfg)

```
[bumpversion]
commit = True
tag = True
current_version = 0.8.4
parse = (?P<major>\d+)\.(?P<minor>\d+)\.(?P<patch>\d+)
serialize = 
	{major}.{minor}.{patch}
[bumpversion:file:VERSION]
```

Before release,I just run like below:

```
bumpversion minor # or major or patch
```

This update minor version number in VERSION file. (In this case, `0.8,4` -> `0.9.0`)
And I configure `commit=True` and `tag=True` in .bumpversion.cfg,The [bumpversion][] makes a commit and tag. 

```
commit 243ccf6932d97a3d100de5a2a7b102e1b2dc8cb2 (tag: v0.0.2)
Author: anarcher
Date:   Tue Dec 5 21:09:21 2017 +0900

    Bump version: 0.0.1 → 0.0.2

commit 9e8ff3f08a1039cc10d93749aa805e939cf28f73 (tag: v0.0.1)
Author: anarcher
Date:   Tue Dec 5 21:08:17 2017 +0900

    Bump version: 0.0.0 → 0.0.1
```

## govvv 

[govvv][] is go binary versioning tool that wraps the go build command. Go provides access to setting variables by compile time linker flags. [govvv][] provides easy versioning for go versioning even I don't know that `--ldflag` of `go build`. 

```    
govvv install -pkg github.com/pkg/version github.com/cmd/
```

I just created version/version.go 

```
package version

var (
	Version, GitCommit, GitState, BuildDate string
)
```

Because I make new `VERSION` file with [bumpversion][], [govvv][] adds a Version variable from `VERSION` file to my version package.
I like to seperate version variable with main.go. version variable can be used in another ways. (e.g. metric labels, api version call..)

```
$cmd -v
Cmd version 0.8.4 git=2639d4e build=2017-12-05T15:57:13Z
```


## drone 

[drone][] is an open source continuous delivery tool. Interesting features of [drone][] are that all tasks in drone is docker containers and drone's yaml configure file is a superset  of docker-compose configure file.  

```
clone:
    git:
        image: plugins/git
        network_mode: host
        depth: 1
pipeline:
    test:
        image: golang:1.7.1
        commands:
            - make deps
            - make test
            - make test-ci
    docker:
        image: plugins/docker
        dockerfile: Dockerfile
        repo: app 
        force_tag: true
        tags: ${DRONE_TAG}
        when:
            event: tag
    slack:
        image: plugins/slack
        webhook: https://hooks.slack.com/
        when:
            status: [ success, failure ]
``` 

If repository __tag__ hooks is enable in drone settings (default disable),git commit tag will sync up with docker tag. 
This is simply how to use the tools for releasing and versioning.

[bumpversion]: https://github.com/peritus/bumpversion 
[govvv]: https://github.com/ahmetb/govvv
[drone]: https://github.com/drone/drone
