# protovision

_Protovision, I have you now_

![screenshot](./screenshot.png)

___
## release

(Current executable file `protovision` build for `amd64` only)

## dev

```bash
go mod tidy
```
___
## build



```bash
sudo apt install portaudio19-dev
sudo apt install libmpg123-dev


go build -o protovision
```
___
## run
```bash
./protovision
```


## arguments

    Usage of ./prototype
    -bd int
            Specify baud rate. (default 20)
    -load
            Load conversation from files
    -var int
            Specify variability. (default 3)
