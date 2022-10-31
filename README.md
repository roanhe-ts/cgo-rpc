# cgo thrift
### Prerequest
* thrift devel is installed in your system

### Build
run `go mod tidy` before make binary.
```sh
# build wrapper library
mkdir build 
cd build
cmake ..
make

# build binary processed by cgo
make bookstore_go
```