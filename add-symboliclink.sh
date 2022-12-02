#!/bin/sh

SRCDIR=$1
FILES="call header inline inputs message results" # except functions.tmpl
ORIGINALDIR="../original" # relative to $SRCDIR

for FILE in $FILES; do
  ln -sf ${ORIGINALDIR}/${FILE}.tmpl ${SRCDIR}/${FILE}.tmpl
done
