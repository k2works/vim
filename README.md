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

**[⬆ back to top](#構成)**

### 配置

**[⬆ back to top](#構成)**

### 運用

**[⬆ back to top](#構成)**

### 開発

**[⬆ back to top](#構成)**

## 参照
- [Conventional Commits 1.0.0](https://www.conventionalcommits.org/ja/v1.0.0/)