# Variables

export CLICOLOR='1'
export COMPLETION_WAITING_DOTS='true'
export EDITOR='vim'
export GOPATH="$HOME/Developer/go"
export HISTCONTROL='ignoreboth'
export HISTFILESIZE="$HISTSIZE"
export HISTIGNORE="&:bg:clear:exit:fg:history:jobs:pwd:* --help:* --version"
export HISTSIZE='32768'
export LANG='en_US.UTF-8'
export LC_ALL='en_US.UTF-8'
export NODE_REPL_HISTORY_SIZE='32768'
export PATH="$PATH:$GOPATH/bin:/Library/TeX/texbin"
export PYTHONIOENCODING='UTF-8'
export ZSH="$HOME/.oh-my-zsh"
export ZSH_THEME='robbyrussell'

# Zsh

fpath=(/usr/local/share/zsh-completions $fpath)
plugins=(colored-man-pages copydir copyfile docker gitfast git-extras github golang history-substring-search httpie jsontools node npm osx pip python rails ruby ssh-agent sudo terminalapp thefuck tmux urltools vagrant wd web-search z)

source "$ZSH/oh-my-zsh.sh"
autoload -U zmv

eval "$(thefuck --alias)"

source /usr/local/share/zsh-navigation-tools/zsh-navigation-tools.plugin.zsh
source ~/.fzf.zsh

type _zsh_autosuggest_start &>/dev/null || source /usr/local/share/zsh-autosuggestions/zsh-autosuggestions.zsh
type _zsh_highlight &>/dev/null || source /usr/local/share/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh

zstyle -e ':completion:*:(ssh|scp|sftp|rsh|rsync):hosts' hosts 'reply=(${=${${(f)"$(cat {/etc/ssh_,~/.ssh/known_}hosts(|2)(N) /dev/null)"}%%[# ]*}//,/ })'
zstyle :omz:plugins:ssh-agent agent-forwarding on

# Aliases

alias ..='cd ..'
alias ...='cd ../..'
alias ....='cd ../../..'
alias .....='cd ../../../..'
alias ......='cd ../../../../..'

alias akg='ack --go --ignore-dir pb --ignore-dir vendor'
alias akj='ack --ignore-dir migrations --ignore-dir node_modules --js'
alias akr='ack --ignore-dir coverage --ignore-dir log --ignore-dir vendor --ruby'

alias br='brew'
alias brc='br cleanup'
alias brk='br cask'
alias brkc='brk cleanup'
alias brki='brk install'
alias brkf='brk info'
alias brkr='brk reinstall'
alias brks='brk search'
alias brku='brk uninstall'
alias brd='br doctor'
alias bri='br install'
alias brl='br list'
alias brn='br info'
alias brr='br rm'
alias brs='br search'
alias bru='brud && brug'
alias bruc='bru && brc && brkc'
alias brud='br update'
alias brug='br upgrade'

alias cd..='cd ..'
alias cd...='cd ../..'
alias cd....='cd ../../..'
alias cd.....='cd ../../../..'
alias cd......='cd ../../../../..'

alias cl='clear'

alias cdd="cd $HOME/Developer"
alias cdg="cd $HOME/Developer/go"
alias cds="cd $HOME/Developer/go/src"
alias cdw="cd $HOME/Developer/go/src/github.com/willfaught"

alias cp='cp -i'

alias dad='date +"%d-%m-%Y"'
alias dat='date +"%T"'

alias diff='colordiff'

alias dsflush='dscacheutil -flushcache'

alias dnsrestart='sudo killall -HUP mDNSResponder'

alias dr='docker'
alias drc='docker-compose'
alias drm='docker-machine'

alias gia='git add'
alias gib='git branch'
alias gibd='git branch -d'
alias gibdf='git branch -D'
alias gic='git checkout'
alias gicb='git checkout -b'
alias gicl='git clean -dfx'
alias gico='git commit'
alias gicoa='git commit -a'
alias gicoaa='git commit -a --amend'
alias gicoaf='git commit -a --fixup'
alias gicoam='git commit -a -m'
alias gicom='git commit -m'
alias gicor='git commit --amend --no-edit --reset-author'
alias gicm='git checkout master'
alias gicp='git cherry-pick'
alias gicpa='git cherry-pick --abort'
alias gicpc='git cherry-pick --continue'
alias gid='git diff'
alias gidc='git diff --cached'
alias gidcw='git diff --cached --word-diff=color'
alias gidw='git diff --word-diff=color'
alias gif='git fetch'
alias gifm='git fetch me'
alias gifo='git fetch origin'
alias gifu='git fetch upstream'
alias gil='git log'
alias gild='git log --decorate=short'
alias gilo='git log --decorate=short --oneline'
alias gilog='git log --decorate=short --oneline --graph'
alias gim='git merge'
alias gima='git merge --abort'
alias gimc='git merge --continue'
alias gip='git pull'
alias gipp='git pull && git push'
alias gipu='git push'
alias gipum='git push me'
alias gipumf='git push me --force'
alias gipuo='git push origin'
alias gipuof='git push origin --force'
alias gipuu='git push upstream'
alias gipuuf='git push upstream --force'
alias girb='git rebase'
alias girba='git rebase --abort'
alias girbc='git rebase --continue'
alias girbi='git rebase -i'
alias girbim='git rebase -i master'
alias girbm='git rebase master'
alias girbmm='git rebase me/master'
alias girbom='git rebase origin/master'
alias girbum='git rebase upstream/master'
alias girs='git reset'
alias girsh='git reset --hard'
alias girshm='git reset --hard master'
alias girshmm='git reset --hard me/master'
alias girshom='git reset --hard origin/master'
alias girshum='git reset --hard upstream/master'
alias girm='git remote'
alias girma='git remote add'
alias girmam='git remote add me'
alias girmao='git remote add origin'
alias girmau='git remote add upstream'
alias girmrm='git remote remove'
alias girmrn='git remote rename'
alias girms='git remote set-url'
alias girmsp='git remote set-url --push'
alias girmspm='git remote set-url --push me'
alias girmspo='git remote set-url --push origin'
alias girmspu='git remote set-url --push upstream'
alias girmsm='git remote set-url me'
alias girmso='git remote set-url origin'
alias girmsu='git remote set-url upstream'
alias girmv='git remote -v'
alias gis='git status'
alias gish='git show'
alias gishw='git show --word-diff=color'

alias gob='go build ./...'
alias gof='go fmt ./...'
alias gog='go generate ./...'
alias goi='go install ./...'
alias gol='gometalinter --deadline 30s --enable-all --disable lll --disable test --disable testify --tests --vendor'
alias golf='gometalinter --deadline 30s --enable-all --disable lll --disable test --disable testify --fast --tests --vendor'
alias gom='goimports -d -w .'
alias gor='go run *.go'
alias got='go test ./...'
alias gotr='go test -run'

alias hi='history'

alias jo='jobs'

alias lsa='ls -A'
alias lsal='ls -Ahl'
alias lsl='ls -hl'

alias ln='ln -i'

alias macappinfo='codesign -dvv'
alias macappstartup='lsal /Library/LaunchAgents /Library/LaunchDaemons /Library/StartupItems'
alias macappverify='codesign -vv'
alias macpkgexpand='pkgutil --expand'
alias macpkgverify='pkgutil --check-signature'
alias macprintgatekeeper='spctl --status'

alias mk='make'

alias mkdir='mkdir -p'

alias mv='mv -i'

alias prfilescommand='sudo lsof -c'
alias prfilesestablished='lsof -P -i4 | grep ESTABLISHED'
alias prfileslisten='lsof -P -i4 | grep LISTEN'
alias prfilespid='sudo lsof -p'
alias prpath='echo -e ${PATH//:/\\n}'
alias prpidsdir='sudo lsof +d'

alias up1='cd ..'
alias up2='cd ../..'
alias up3='cd ../../..'
alias up4='cd ../../../..'
alias up5='cd ../../../../..'

alias vi='vim'
alias vid='vi Dockerfile'
alias vic='vi docker-compose.yml'
alias viz='vi ~/.zshrc && source ~/.zshrc'
alias vizl='vi ~/.zshrc.local; test -e ~/.zshrc.local && source ~/.zshrc.local; true'

alias web='python -m SimpleHTTPServer 8080'

alias zmv='noglob zmv -W'

# Functions

banner() { toilet -w "$(tput cols)" "$@" }

finish() { eval "$@" && finished "$@" }

notify() {
  local command=''
  if test -n "$1"; then command="$command with title \"$1\""; fi
  if test -n "$2"; then command="$command subtitle \"$2\""; fi
  if test -n "$3"; then command="$command sound name \"$3\""; fi
  command="display notification \"${@:4}\"$command"
  osascript -e "$command"
}

notify-alert() { notify 'Alert' "$1" 'Ping.aiff' "${@:2}" }

notify-beginning() { note 'Beginning' "$@" }

notify-ending() { note 'Ending' "$@" }

notify-failed() { warning 'Failed' "$@" }

notify-finished() { alert 'Finished' "$@" }

notify-note() { notify 'Note' "$1" 'Pop.aiff' "${@:2}" }

notify-starting() { note 'Starting' "$@" }

notify-stopping() { note 'Stopping' "$@" }

notify-succeeded() { alert 'Succeeded' "$@" }

notify-warning() { notify 'Warning' "$1" "Basso.aiff" "${@:2}" }

page() { "$@" |& less -F }

port() { lsof -ni ":$1" | grep LISTEN }

up() { while test "$(pwd)" != "/" && test "$(basename "$(pwd)")" != "$1"; do cd ..; done }

# Local

test -f ~/.zshrc.local && source ~/.zshrc.local

# Tools

# Debug: top, nettop, lsof, man -k snoop, opensnoop, rwsnoop, iosnoop, dtrace, fs_usage, tccutil, tcpdump
# Disk: diskutil, df, du
# Format: textutil, plutil, iconutil, hexdump, xxd
# Compression: zip, unzip, gzip, gunzip, bzip2, bunzip2, tar, gzcat, bzcat
# Cryptography: shasum
