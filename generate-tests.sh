#!/bin/sh

SRCDIR=$(basename $1)
FILES="call header inline inputs message results" # except function.tmpl
ORIGINALDIR="../original" # relative to $SRCDIR

for FILE in $FILES; do
  ln -sf ${ORIGINALDIR}/${FILE}.tmpl ${SRCDIR}/${FILE}.tmpl
done

go run github.com/cweill/gotests/gotests@latest -w -all -parallel -template_dir ./${SRCDIR} ./example.go
mv example_test.go example_${SRCDIR}_test.go
