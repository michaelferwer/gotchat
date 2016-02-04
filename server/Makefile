NO_COLOR=\033[0m
OK_COLOR=\033[1;32m

all: env dependencies build run

env: ; @echo "$(OK_COLOR)--SETTING ENVIRONMENT VARIABLES$(NO_COLOR)" ; export GB_PROJECT_DIR=`pwd`

dependencies: ; @echo "$(OK_COLOR)--DEPENDENCIES$(NO_COLOR)" ; gb vendor restore

build : ; echo "$(OK_COLOR)--BUILD$(NO_COLOR)" ;  gb build

run: ; echo "$(OK_COLOR)--RUN$(NO_COLOR)" ; ./bin/gotchat