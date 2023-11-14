# Ionix
is an api that has an authentication service with two cruds, one for vaccination and one for drugs.

# Prerequisites
Please create the `.env` file take in account that you can find an example in the root path `.env-sample`

# Run
You can find a file called Makefile to run it you will need to have "make" insatalled on your computer.

- you can be able to installed with the command below with 
linux:
`sudo apt install make`
macOS:
`brew install make`

__1.)__ Run with docker

`make run-docker`

once you no longer need the container please run the following command

`make down-docker`

Once the project is up I suggest you go to the documentation built in swagger and make the request with the user
interface.

`http://localhost:8080/swagger/index.html`