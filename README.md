# open-match-demo
## Introduction

This tutorial aims to give you the tools to create your OpenMatch and use it with Arbitrium (Edgegap's solution). Like a wise man once said: "Keep it stupid simple." In this tutorial, we will not go in-depth about performances. This code will not be ready for production purposes. However, you should be able to create your production-ready Open Match!

## Objectives

At the end of this tutorial, you will have built 3 components (Director, Front End, Match function) that will run on a local Open Match environment. You will be able to create, get and delete tickets using HTTP requests. Your Matchmaker will match players together and automatically deploy servers on Arbitrium.

[Follow Tutorial](https://docs.edgegap.com/docs/matchmaker/component-tutorial)

### Table of Contents

<!-- TABLE OF CONTENTS -->
* [About The Project](#about-the-project)
* [Getting Started](#getting-started)
    * [Prerequisites](#prerequisites)
    * [Installation](#installation)
    * [How to use it](#how-to-use-it)
* [Acknowledgements](#acknowledgements)



<!-- ABOUT THE PROJECT -->
## About The Project

This is the source code associated with [this](https://docs.edgegap.com/docs/openmatch/tutorials/basic/open-match-tutorial-basics-introduction) tutorial.



<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

* [Install](https://golang.org/dl/) Golang on your computer
* Install [Docker](https://docs.docker.com/docker-for-windows/install/)
* Having a Kubernetes ecosystem available _([Docker Desktop includes a standalone Kubernetes server and client](https://docs.docker.com/docker-for-windows/kubernetes/))_


### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/edgegap/open-match.git
   ```
2. Go to the tutorial's folder
   ```sh
   cd [PATH_TO_PROJECT]/tutorials/basic-components/go
   ```
3. In ***./director/main.go*** change the variables to put your values.
    ```go
    // Game server data
    gameServerPort = "<APP_PORT>"      // String | E.G. 25565
    appName        = "<APP_NAME>"      // E.G. MySuperGame
    appVersion     = "<APP_VERSION>"   // E.G. V1
    // You MUST have a forward slash (/) at the end of your URL
    arbitriumAPI   = "<ARBITRIUM_API>" // E.G. https://api.edgegap.com/
    // You MUST NOT have prefix "token" in your API token value
    // token 08230a25-0fdb-4f56-917b-0a58ec35cbaf INVALID
    // 08230a25-0fdb-4f56-917b-0a58ec35cbaf VALID
    apiToken       = "<API_TOKEN>"     // E.G. 08230a25-0fdb-4f56-917b-0a58ec35cbaf
    ```
4. Install using the appropriate script

   **Windows (Powershell)**
    ```sh
    install.ps1
    ```

   **Windows (CMD)**
    ```sh
    install.bat
    ```

5. You can delete it using the appropriate script

   **Windows (Powershell)**
    ```ps1
    delete.ps1
    ```

   **Windows (CMD)**
    ```ps1
    delete.bat
    ```

### How to use it

You can learn how to use your matchmaker [here](https://docs.edgegap.com/docs/openmatch/tutorials/basic/open-match-tutorial-basics-how-to-use-it)


## Acknowledgements
* [Open Match](https://openmatch.dev/site/docs/guides/)