#!/bin/bash

## para agregar permisos
# Agregar permisos al script chmod +x ./run.sh

#creamos la estructura
cd ..
cd ..

if [ ! -d bin ]; then
  echo mkdir bin
fi

#compilamos el cliente
cd ./Client/src
go build -o "../../bin/Client" Client.go

#volvemos a bin
cd ..
cd ..
cd bin

#corremos el cliente
read -p "Quieres levantar el Cliente?(+si)" par
if test "$par" = "si"
then
     ./Client
else
     echo "Compilado en Bin Client, pero no ejecutado"	
fi
