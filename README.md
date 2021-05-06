# trimr

> trimr - Trim those unused branches with ease.

## Table of Contents

- [trimr](#trimr)
  - [Table of Contents](#table-of-contents)
  - [What is `trimr`](#what-is-trimr)
  - [Commands](#commands)
    - [`trimr`](#trimr-1)
    - [`trimr config`](#trimr-config)
  - [Install](#install)
  - [TODO](#todo)

## What is `trimr`

Do you work in a high velocity environment where you are constantly creating new branches on your fork for all your feature contributions? Do you find you have an abundance of merged local branches you no longer use? The `trimr` application aims to make clean-up of those stale branches simple.

## Commands

### `trimr`

To run the trimr program, simply give it the path to your repository:

```bash
trimr -p <repository_dir>
```

Trimr will then give you a prompt asking you if you want to delete a specific branch. To avoid the prompts you can supply the `--no-confirm` flag.

### `trimr config`

Trimr ignores branches listed in the trimr configuration file under `branches.protected`. The config command allows you to add or remove branches from the protected branches list.

You can add a branch to the protected branches list as follows:

```bash
trimr config protected-branch add -n <branch_name>
```

You can do something similar to remove branches:

```bash
trimr config protected-branch remove -n <branch_name>
```

## Install

You can download pre-built binaries from the [latest release](https://github.com/FuzzyStatic/trimr/releases) and copy them to your preferred binary directory.

## TODO

- Ability to trim remote branches
- Per repo configuration
