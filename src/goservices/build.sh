VER=$(cat seq.txt)
VER=$((VER+1))
echo $VER>seq.txt

docker build -t alemoracr/goconsole:${VER} -f Dockerfile-Console .
docker build -t alemoracr/gofiber:${VER} -t alemoracr/gofiber:latest .
