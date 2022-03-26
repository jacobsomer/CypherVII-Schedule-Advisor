# CypherVII-Schedule-Advisor
Degree Works with a more intuitive UX/UI

## Installation
### Prerequisites
1. A recent installation of Go. Here is the [download for Windows](https://go.dev/dl/go1.18.windows-amd64.msi). Here is the [download for macOS](https://go.dev/dl/go1.18.darwin-amd64.pkg)
2. A recent installation of node and npm. Here is the [download for Windows](https://nodejs.org/dist/v16.14.2/node-v16.14.2-x64.msi). Here is the [download for macOS](https://nodejs.org/dist/v16.14.2/node-v16.14.2.pkg)
3. [Git](https://git-scm.com/downloads) if you don't already have it installed.

Once downloaded successfully, your command promp (Terminal if using mac) should look something like this when you run the following commands:

![Windows Command Prompt After Download](https://user-images.githubusercontent.com/60264650/160229555-526685ae-86ee-4446-bb74-f8cc49488638.png)

## Project Setup
Open a terminal(macOS) or command prompt(Windows). `cd` to the folder you wish to work out of. Then run the following command:
### Go Setup

```shell
mkdir cypherVIIWorkspace
cd cypherVIIWorkspace
git clone https://github.com/jacobsomer/CypherVII-Schedule-Advisor.git
cd CypherVII-Schedule-Advisor
git checkout -b myLocalBranch
```
Once you have done this. Feel free to open up a code editor such as VSCode in the CypherVII-Schedule-Advisor folder.
### Backend Setup
From the root of the project, run the following commands to start the go server.
```shell
cd backend
go run main.go
```
Here is what my console looked like:
![Screenshot (41)](https://user-images.githubusercontent.com/60264650/160234335-216b9f0c-8620-4bf7-bdd2-9ca5806c706b.png)
### Frontend Setup
From the root of the project, run the following commands to start the react server.
```shell
cd frontend
npm i
npm start
```

### Speed Tip!
If you are using VScode, I recommend installing this [extension](https://marketplace.visualstudio.com/items?itemName=golang.Go) to make debugging go a lot easier. 
