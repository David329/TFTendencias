#!/bin/bash

## para agregar permisos
# Agregar permisos al script chmod +x ./run.sh

#Levantamos mongo
sudo service mongod start

#Regresamos a la raiz para crear la carpeta bin
cd ./../../

#creamos la estructura

if [ ! -d bin ]; then
  echo mkdir bin
fi

#compilamos el servidor
cd ./Server/src
go build -o "../../bin/Server" ./Server.go

#corremos el servidor
read -p "Quieres levantar el Servidor?(+si)" par
if test "$par" = "si"
then
     ./../../bin/Server
else
     echo "Compilado en Bin Server, pero no ejecutado"	
fi

##NO CORRER DESDE LA MISMA RUTA LOS BINARIOS, X EL ERROR MAGICO!
