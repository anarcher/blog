<!-- Title: Using docker-machine -->
<!-- Tags: docker -->

Docker machine is a command tool created by docker team to manage docker servers. It automatically create hosts and install docker engine on them,configure the docker client to talk. 

If you install the docker machine tool,you can use it like below:

    $docker-machine ls
    NAME    ACTIVE   DRIVER       STATE     URL                        SWARM
    host1            generic      Running   tcp://192.168.99.100:2376   host1 (master)
    host2            generic      Running   tcp://192.168.99.101:2376   host1
    host3            generic      Running   tcp://192.168.99.102:2376   host1
    demo0            virtualbox   Stopped                              
    demo1            virtualbox   Stopped                              
    demo2            virtualbox   Stopped                              
    demo3            virtualbox   Stopped                              
    infra            virtualbox   Stopped      

    $docker-machine ssh host1
    ...


    $docker $(docker-machine config host1) ps 
    CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS               NAMES
    ea13099c15d3        swarm:latest        "/swarm join --addr 1"   7 weeks ago         Up 4 days           2375/tcp            swarm-agent

    ...


It's useful to use docker command with docker machine command for to connect docker engine on multi hosts.

`docker-machine config` is just printed like below.

    $docker-machine config host1
    --tlsverify --tlscacert="/home/anarcher/.docker/machine/machines/host1/ca.pem" 
    --tlscert="/home/anarcher/.docker/machine/machines/host1/cert.pem" 
    --tlskey="/home/anarcher/.docker/machine/machines/host1/key.pem" -H=tcp://192.168.99.100:2376

As you know,The bash command substitution can make the output of docker-machine command to replace the command itself. It's the one of funny things of unix shell too. :-)

Docker machine keeps the configuraion files that is an information to connect on docker hosts to ~/.docker/machine/. If you want to share it the others,you just to copy it and share it. 

Docker machine is based on ssh authentication to setup docker host nodes. So you can log in hosts to use `docker-machine ssh [name]`.

Docker machine has many drivers that represent the virtual environments.The virtual environments are local linux,AWS,Digital ocean and many more.  Also You can add a existed Docker host to use driver `none` or If the host hasn't a docker engine,you can use the driver `generic` to install docker engine. 

Docker machine has many options. So I used to use `grep` command with Unix pipe for searching the right command. 

    $docker-machine create -h | grep digi
       --digitalocean-access-token              Digital Ocean access token [$DIGITALOCEAN_ACCESS_TOKEN]
       --digitalocean-backups                   enable backups for droplet [$DIGITALOCEAN_BACKUPS]
       --digitalocean-image "ubuntu-14-04-x64"  Digital Ocean Image [$DIGITALOCEAN_IMAGE]
       --digitalocean-ipv6                      enable ipv6 for droplet [$DIGITALOCEAN_IPV6]
       --digitalocean-private-networking        enable private networking for droplet [$DIGITALOCEAN_PRIVATE_NETWORKING]
       --digitalocean-region "nyc3"             Digital Ocean region [$DIGITALOCEAN_REGION]
       --digitalocean-size "512mb"              Digital Ocean size [$DIGITALOCEAN_SIZE]



Basically It's convenience tool looks like other docker tools. But unlike `ansible` or `chef`, It's only care Docker engine. If you want to host configuration, You should use the other tools with docker engine. 

