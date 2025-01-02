#!/bin/bash

# Detecta o diretório do script atual, independentemente de onde estiver
SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)

# Gera a versão com base no timestamp e número do build do TeamCity
BUILD_NUMBER="${BUILD_NUMBER:-SNAPSHOT}"
VERSION="$(date +%Y.%m).${BUILD_NUMBER}"

# Exporta a versão para o ambiente
export VERSION

# Define as imagens com a versão e com a tag 'latest'
IMAGE_NAME="scheffer/bifrost-fr"
IMAGE_VERSION_TAG="${IMAGE_NAME}:${VERSION}"
IMAGE_LATEST_TAG="${IMAGE_NAME}:latest"

# Faz o build da imagem Docker com as duas tags
docker build -f "./deploy/Dockerfile" --build-arg VERSION="$VERSION" -t "$IMAGE_VERSION_TAG" -t "$IMAGE_LATEST_TAG" .

# Faz o deploy usando docker-compose na raiz do projeto
VERSION=$VERSION docker compose -f "./deploy/docker-compose.yml" up -d --build --force-recreate

# Exibe mensagem de sucesso
echo "Deploy realizado com sucesso! Versão: $VERSION"