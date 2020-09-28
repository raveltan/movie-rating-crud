echo "Creating deployment executable for Mac,Linux and windows ...\n"

echo "Building for linux ..."

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./dist/server.linux .

echo "Building for windows ..."

GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o ./dist/server.exe .

echo "Building for Mac OSx ..."

GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o ./dist/server.darwin .

echo "Building frontend ..."

cd ./frontend
yarn
yarn build

cd ..

cp -r ./frontend/dist ./dist/public 
cp ./README.md ./dist/README.md
tar -cvzf dist.tar.gz dist
rm -r ./dist
echo "Done :)"