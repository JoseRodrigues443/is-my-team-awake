#!/bin/bash

echo "Starting demo of all the commands of the command is-my-team-awake"

echo "\n\n__ HELP __ \n\n"
go run . -h

echo "\n\n__ ADD __ \n\n"
go run . add -name "San Jonh Doe" --location "Europe/London"

echo "\n\n__ LIST __ \n\n"
go run . list

echo "\n\n__ CHECK __ \n\n"
go run . check --name "San Jonh Doe"

