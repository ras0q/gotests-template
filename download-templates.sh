#!/bin/sh

FILES="call function header inline inputs message results"
BASEURL=" https://raw.githubusercontent.com/cweill/gotests/develop/internal/render/templates"
ORIGINALDIR="./original"

mkdir -p $ORIGINALDIR
for FILE in $FILES; do
  curl ${BASEURL}/${FILE}.tmpl -o ${ORIGINALDIR}/${FILE}.tmpl -s
done
