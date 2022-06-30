VER=$(cat seq.txt)
VER=$((VER+1))
echo $VER>seq.txt

docker build -t alemoracr/goconsole:${VER} -f Dockerfile-Console .
docker build -t alemoracr/gofiber:${VER} -t alemoracr/gofiber:latest .
docker build -t alemoracr/cmdvendors:${VER} -f Dockerfile-cmdVendors .
docker build -t alemoracr/cmdproducts:${VER} -f Dockerfile-cmdVendors .
docker build -t alemoracr/cmdwarehouses:${VER} -f Dockerfile-cmdVendors .
docker build -t alemoracr/evtinventory:${VER} -f Dockerfile-cmdVendors .