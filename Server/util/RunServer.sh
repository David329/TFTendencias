#!/bin/bash

## para agregar permisos
# Agregar permisos al script chmod +x ./run.sh

#creamos la estructura
cd ..
cd ..

if [ ! -d bin ]; then
  echo mkdir bin
fi

#compilamos el servidor
cd ./Server/src
go build -o "../../bin/Server" Server.go

#volvemos a bin
cd ..
cd ..
cd bin

#corremos el servidor
read -p "Quieres levantar el Servidor?(+si)" par
if test "$par" = "si"
then
     ./Server
else
     echo "Compilado en Bin Server, pero no ejecutado"	
fi
