# goclip
simple xsel replacement
allows using */+ buffer in neovim on headless servers

## install
```
go get github.com/dbldots/goclip
sudo ln -s $GOPATH/bin/goclip /usr/local/bin/xsel
```

### usage

* open vim (or nvim)
* select text (using C-v or C-V)
* hit `"*y`
* open another vim
* hit `"*p`

the 'copied' text is written into `~/.goclip`
