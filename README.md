<div align="center">
<h1>infinite</h1>
<span>  Component library for developing interactive CLI(tui,terminal) programs. </span>
<br>
<a href="https://goreportcard.com/report/github.com/yunginnanet/infinite"><img src="https://goreportcard.com/badge/github.com/yunginnanet/infinite" alt="go report card"></a>
</div>
<img src="https://user-images.githubusercontent.com/65269574/184916069-076a0f6a-70bd-49e1-b7d7-0d2e7fc5c6bb.gif" alt="demo">


## Features

- Provides a range of out-of-the-box components 
    - autocomplete
    - progress-bar group
    - multi/single select
    - spinner
    - confirm(input/selection)
    - text input
- Cross-platform
- Customizable, you can replace some options or methods in the component with your own implementation

- Composable, you can combine components together to build your own interactive program
    - `autocomplete` + `input` To achieve input reception into `selection` to choose options.
    - `selection` instantiation via `input` for option filtering.

## Install

```shell
go get github.com/yunginnanet/infinite
```
