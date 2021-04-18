#!/bin/bash

# Eh gerado o env com a shell script
if [ ! -f ".env" ]; then
    cp .env.example .env
fi

npm install

# A funcao abaixo faz com que recompile o codigo sempre que uma alteracao for feita
npm run start:dev