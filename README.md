# Greeting

## Steps to Run

    make help

### run database

```sh
make roachup
```

### create database

```sh
make roachcli
CREATE DATABASE greeting;
CREATE USER account_user;
GRANT ALL ON DATABASE greeting TO account_user;
```

### run project

```sh
make run
or
make drun
```

### TODO

 - [ ] add more test  
 - [ ] use context
 - [ ] graceful sigterm
