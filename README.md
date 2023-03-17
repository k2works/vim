# My Vim Environment
## 概要

### 目的

### 前提

| ソフトウェア | バージョン | 備考 |
| :----------- | :--------- | :--- |
| vim          |     |      |
| neovim       |     |      |
| nodejs       | 16.3.0    |      |

## 構成

- [構築](#構築)
- [配置](#配置)
- [運用](#運用)
- [開発](#開発)

## 詳細

### Quick Start

#### Windows

neovim

```
cp init.vim  $env:LOCALAPPDATA\nvim
```
#### Linux

neovim

```
ln -s init.vim  $HOME/.config/nvim
```

vim

```
ln -s .vimrc $HOME/.vimrc
```

### 構築


[coc.vimで'[coc.nvim] build/index.js not found, please install dependencies and compile coc.nvim by: yarn install'と言われた](https://qiita.com/Taichi-yzrh/items/5868e618c82e328c89f6)
```vim
: call coc#util#install()
```

[Copilot.vim](https://github.com/github/copilot.vim)

[Easiest way to update Neovim on Ubuntu](https://medium.com/@leonardormlins/easiest-way-to-update-neovim-on-ubuntu-a283c66d5322)
```
sudo add-apt-repository ppa:neovim-ppa/stable
sudo apt-get update
sudo apt-get install neovim
```
[Ubuntu: Vimの最新バージョンインストール](https://qiita.com/Fell/items/8619385da8e09a59c787)
```
sudo add-apt-repository ppa:jonathonf/vim
sudo apt update
sudo apt install vim
```

[NeoVim: vimspector unavailable: Requires Vim compiled with +python3](https://stackoverflow.com/questions/74036547/neovim-vimspector-unavailable-requires-vim-compiled-with-python3)
```
pip3 install neovim
```

**[⬆ back to top](#構成)**

### 配置

**[⬆ back to top](#構成)**

### 運用

**[⬆ back to top](#構成)**

### 開発

**[⬆ back to top](#構成)**

## 参照
- [Conventional Commits 1.0.0](https://www.conventionalcommits.org/ja/v1.0.0/)
