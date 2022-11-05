
# Tiny Container
Simple Container implementation in go using linux namespaces and cgroups.
##  How to use
**Get the package**
```sh 
 go get github.com/Eslam-Nawara/tinycontainer
```
## Usage and Installation:
**Install the tinycontainer**
```sh 
go install github.com/Eslam-Nawara/tinycontainer/cmd/container@latest
```
**Install the container requirements**
- Download and the run script `install.sh` in your work space. 
```sh
wget https://github.com/Eslam-Nawara/tinycontainer/raw/main/install.sh
 ```
 ```sh
chmod +x install.sh
```
 ```sh
./install.sh
```
- Run the container
```sh
 sudo container run <Cmd> <args>
```
