#!/bin/bash

## para agregar permisos
# Agregar permisos al script chmod +x ./run.sh


#Regresamos a la raiz para crear la carpeta bin
cd ./../../

#creamos la estructura

if [ ! -d bin ]; then
  echo mkdir bin
fi

#compilamos el Servidor Web
cd ./WebServer/src
go build -o "../../bin/WebServer" ./WebServer.go

#corremos el Servidor Web
read -p "Quieres levantar el Servidor Web?(+si)" par
if test "$par" = "si"
then
     ./../../bin/WebServer
else
     echo "Compilado en Bin WebServer, pero no ejecutado"	
fi

##NO CORRER DESDE LA MISMA RUTA LOS BINARIOS, X EL ERROR MAGICO!
