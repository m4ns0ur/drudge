# drudge

[![Documentation](https://godoc.org/github.com/willfaught/drudge?status.svg)](https://godoc.org/github.com/willfaught/drudge)
[![Build](https://travis-ci.org/willfaught/drudge.svg?branch=master)](https://travis-ci.org/willfaught/drudge)
[![Report](https://goreportcard.com/badge/github.com/willfaught/drudge)](https://goreportcard.com/report/github.com/willfaught/drudge)
[![Test Coverage](https://coveralls.io/repos/github/willfaught/drudge/badge.svg?branch=master)](https://coveralls.io/github/willfaught/drudge?branch=master)

Drudge takes the drudgery out of setting up and tearing down your computer user accounts. You define a set of tasks and then invoke them by name. A task is a Go function that returns an error. A task name has a verb and an object, like "install foo". A task can dependend on other tasks or the environment. A task runs after all its dependencies and their dependencies and so on.

My personal tasks are checked in, like tasks for installing program and operating system configuration, installing Homebrew, Go, and Node packages, generating SSH keys, and so on. If you don't want them, or want to modify them, or want to add your own, then fork this repository and go nuts. It was designed to be a minimal and flexible framework for doing whatever drudgery you need.

By default, it prints a period every time a task makes progress, which is when either the task finishes or it reports the progress itself. To print nothing but errors, use the q flag. To print the log, use the v flag. To print the log and not make changes, use the d flag.

Install Drudge with

```sh
go get -u github.com/willfaught/drudge
```

Get help with

```sh
drudge -h
```

Tasks are run like

```sh
drudge verb
```

which runs all tasks with that verb, or

```sh
drudge verb object
```

which runs the task "verb object", or

```sh
drudge verb object otherobject
```

which runs the tasks "verb object" and "verb otherobject".
