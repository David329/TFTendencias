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
cd ./WebServer2/src
go build -o "../../bin/WebServer2" ./WebServer.go

#corremos el Servidor Web
./../../bin/WebServer2

##NO CORRER DESDE LA MISMA RUTA LOS BINARIOS, X EL ERROR MAGICO!
