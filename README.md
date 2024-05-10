# ketos <img alt="ketos" src="https://github.com/Doer-org/ketos/assets/77223796/0bb34886-2efb-4577-b934-23f1eb9d6129" width="40px"/> 

<img alt="ketos" src="https://images.unsplash.com/photo-1568430328012-21ed450453ea?q=80&w=1000&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxleHBsb3JlLWZlZWR8Mnx8fGVufDB8fHx8fA%3D%3D" width="600px"/> 
Ketos is a tool designed to swiftly share your local development environment with people around the world. It transforms your local setup into a Docker image, making it accessible for anyone to pull and use. This tool streamlines the process of sharing your work environment, ensuring that collaborators or users can replicate your setup with ease, fostering a more collaborative and efficient development process. Ideal for developers looking to simplify the sharing of their environments, Ketos bridges the gap between local development and global collaboration.

## Motivation
When you want to share your web apps with others, you can either upload the docker image on docker hub or source code on github. With ketos, it will be easier to share errors in your development environment with others, or to share the specifications of your development environment with others.
![image](https://github.com/Doer-org/ketos/assets/88176012/17d19ee6-12aa-431b-bf46-616d05fd9fb0)


## Install
```
brew install buildpacks/tap/pack
go install github.com/Doer-org/ketos@latest
```
## Usage
### push
```
ketos push -d "./examples/go" -l "go" -f "Dockerfile" -D true -p 8090:8090 -s "http://localhost:8000/"
directory:  ./examples/go
Creating image with Dockerfile
Path:  ./examples/go
Dockerfile:  Dockerfile
Sending tar to server...
{"id":"6ce1da4c-b7be-4c02-b262-88cd6350e025","port":"8090:8090"}

            __ __ ________________  _____
           / //_// ____/_  __/ __ \/ ___/
          / ,<  / __/   / / / / / /\__ \
         / /| |/ /___  / / / /_/ /___/ /
        /_/ |_/_____/ /_/  \____//____/

🐳🐳🐳 Share this command 🐳🐳🐳

ketos pull -i 6ce1da4c-b7be-4c02-b262-88cd6350e025
```

### pull
```
ketos pull -i 6ce1da4c-b7be-4c02-b262-88cd6350e025 -s "http://localhost:8000"
http://localhost:8000/info/6ce1da4c-b7be-4c02-b262-88cd6350e025
{"id":"6ce1da4c-b7be-4c02-b262-88cd6350e025","port":"8090:8090"}
https://ketos.doer-app.com/info/6ce1da4c-b7be-4c02-b262-88cd6350e025
Container ketos-tmp-container has been created with ID: ce98e6ba0f24a82ac2ec09ee34f90a151682c5d2707c0fe9b4bdea958e0a6975
Container ketos-tmp-container has been started

            __ __ ________________  _____
           / //_// ____/_  __/ __ \/ ___/
          / ,<  / __/   / / / / / /\__ \
         / /| |/ /___  / / / /_/ /___/ /
        /_/ |_/_____/ /_/  \____//____/
```

## LICENSE
Code licensed under 
[the MIT License](https://github.com/Doer-org/ketos/blob/main/LICENSE).
