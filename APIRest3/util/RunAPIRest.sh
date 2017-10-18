cd ./../
if [ ! -d bin ]; then
  echo mkdir bin
fi

cd ./APIRest3/src/

go build -o "../../bin/APIRest3" ./APIRest.go

./../../bin/APIRest3
