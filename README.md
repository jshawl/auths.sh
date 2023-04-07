# auths.sh

## Local Setup

Update `~/.ssh/config` to include:

```
Host localhost
    UserKnownHostsFile /dev/null
```

and then start the server

```
go run main.go
```