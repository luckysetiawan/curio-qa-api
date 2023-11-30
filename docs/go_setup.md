# Curio QA API

## Go Setup for Local Development

This guide will assume Go installation for local development in Windows OS with WSL2, for another development environment, please adjust accordingly.

#### Download Go

You can find a detailed guide and updated download link [here](https://go.dev/doc/install).
```
wget "https://go.dev/dl/go1.21.4.linux-amd64.tar.gz"
```

#### Remove any previous Go installation and extract the downloaded Go

```
sudo rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.4.linux-amd64.tar.gz
```

#### Add Go bin to PATH

```
export PATH=$PATH:/usr/local/go/bin
```
You might not want to type the above command every time WSL2 restarts as exported variables are only kept for a session. In this case, you'll want to add the export command to `~/.bashrc`.
```
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
```
After that you can restart the shell or reload the `~/.bashrc`
```
source ~/.bashrc
```
#### Additional info

If you want to remove the Go installation you can simply do the following steps
1. Delete the go diretory (/usr/local/go)
1. Remove the Go bin directory from PATH
