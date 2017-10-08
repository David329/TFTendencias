#!/bin/bash

## para agregar permisos
# Agregar permisos al script chmod +x ./run.sh

#creamos la estructura
cd ..

if [ ! -d bin ]; then
  echo mkdir bin
fi

#compilamos el servidor
cd ./src
go build -o "../bin/Server" Server.go

//volvemos a bin
cd ..
cd bin

#corremos el servidor
read -p "Quieres levantar el servidor?(+si)" par
if test "$par" = "si"
then
     ./Server
else
     echo "Compilado en Bin, pero no ejecutado"	
fi