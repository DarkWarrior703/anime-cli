<div style="text-align:center">
<h1>anime-cli</h1>
</div>
![Example](./photos/example.png)

## A simple app for fetching data about animes

This app is written in Go, so the only dependency is the Go compiler. 
This project is heavily inspired by [Genzyy's anime-cli](https://github.com/genzyy/anime-cli), but it's not a copypaste.

## Installing
Simple as 
```
go install
```

This will place the binary where $GOBIN indicates to.
By default,
```
$GOBIN=$GOPATH/bin 
```

If you wanna install it in a specific directory, change the GOBIN
```
GOBIN=$PATH go install
```

## Usage

```
anime-cli <args>
```
where args is the name of anime

It will display a table containing info about the animes that contain the args.

## To-do

- Turn the anime fetcher into a module
- Make a module for manga