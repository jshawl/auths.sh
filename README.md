# auths.sh

## Local Setup

Update `~/.ssh/config` to include:

```
Host localhost
    UserKnownHostsFile /dev/null
```

and then start the server

```
HOST_PRIVATE_KEY="$(cat .ssh/host_private_key)" go run main.go

# or

docker build . -t authssh:latest
docker run -p 4202:4202 -p 8080:8080 -e "HOST_PRIVATE_KEY=$(cat .ssh/host_private_key)" authssh:latest
```
