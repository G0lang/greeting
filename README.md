# Greeting

    please see the first branch .

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
CREATE USER greeting;
GRANT ALL ON DATABASE greeting TO greeting;
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
