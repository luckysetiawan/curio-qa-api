# Curio QA API

## About

Curio QA is a QA app that allows its users to send curios to all of its users. Curio is something considered novel, rare, or bizarre ([Merriam-Webster](https://www.merriam-webster.com/dictionary/curio)). Something rare must be unique, just like each question each person has for another person. With this app, its users will be able to send and receive their questions in the form of curios. The receiver could answer the curio on any platform they like and then mark the curio as answered.

## Setup

This guide will assume setup in Windows OS with VS Code, WSL2, and Docker Desktop. For another development environment, please adjust accordingly.

### Install VS Code

Download the installation files [here](https://code.visualstudio.com/Download).

### Install WSL2

Open PowerShell and install WSL2
```
wsl --install -d Ubuntu-22.04
```

### Install Docker Desktop

Download the installation files [here](https://www.docker.com/products/docker-desktop/).

### Go Setup for Local Development

A local Go installation is needed for [the VS Code Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go) to work. If you don't need to use the extension, feel free to skip this step. 

Details about local Go installation can be found [here](docs/go_setup.md).

## Run Curio QA API

### Run With Docker

#### Run automatically with docker-compose (recommended)

Run all
```
docker-compose -f docker-compose.yml up -d
```
Shut down all
```
docker-compose down
```

#### Run manually without docker-compose

To run manually please use [this guide](/docs/docker_setup.md).
