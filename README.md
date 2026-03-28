# mkexp

Create date-prefixed experiment directories from the command line.

```
mkx my new project        # creates ~/.src/20260328_my-new-project and cds in
mkx --git --readme        # with git repo and README
mkx                       # random name, e.g. ~/.src/20260328_silent-river
```

## Install

```bash
go install github.com/dreikanter/mkexp@latest
```

## Setup

Add to `~/.zshrc`:

```zsh
eval "$(mkexp init)"
```

Set a default base directory in `~/.zshenv`:

```zsh
export MKEXP_PATH=$HOME/.src
```

## Usage

```
mkexp new [--readme] [--git] [text...]   create experiment directory
mkexp init                               print shell integration snippet
mkexp --version                          print version
```

| Flag       | Description                        |
|------------|------------------------------------|
| `--readme` | Create a `README.md` in the new directory |
| `--git`    | Run `git init` in the new directory |

## Environment

| Variable     | Default | Description                        |
|--------------|---------|-------------------------------------|
| `MKEXP_PATH` | `./`    | Base directory for new experiments |
