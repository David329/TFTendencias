#!/bin/bash

## para agregar permisos
# Agregar permisos al script chmod +x ./run.sh


#Regresamos a la raiz para crear la carpeta bin
#cd ./../../

#creamos la estructura

if [ ! -d bin ]; then
  echo mkdir bin
fi

#compilamos el Servidor Web
cd ./WebServer1/src
go build -o "../../bin/WebServer1" ./WebServer.go

#corremos el Servidor Web
./../../bin/WebServer1

##NO CORRER DESDE LA MISMA RUTA LOS BINARIOS, X EL ERROR MAGICO!
