# Introduction

This is my third attempt.

I just want made an easy to use, easy to maintenance.

So far, I am very satisfied with the project.

# Features

- wire - injects dependencies
- ent - database orm
- gin - router
- config - koanf
- logger - zap
- and so on...
    - mail
    - ...

# Usage

I also made a cli to init project from template and generate new route file.

So, first install the cli tool:

```bash
go install github.com/Xwudao/neter/cmd/neter@latest
```

Then, you can use the cli to init project from template:

```bash
neter init <project-name> [-m <mod-name>]
```

Then, run `neter run` to start the project.

# Database

In this template, I use `ent` to connect to database.

And, in newer version of ent, it supported versioned-migration.

So, in `neter-cli` I also added a command to migrate database.

```bash
# also you need config the database info in config.yaml that locate in the project root directory
# generate migration file
neter run migrate name=<migration-filename>

# up
neter run migrate up [all=true]

# down
neter run migrate down [all=true]
```



