#!/bin/bash
# [-o] prevent errors in a pipeline from being masked.
# [-u] fail if found references to any variable which was not previously defined
set -uo pipefail


GHCDCD=$(dirname $BASH_SOURCE[0])/../../src/bash/ghcicd.sh
echo "################################################################################"
echo "test source: $GHCDCD"
echo "################################################################################"
source $GHCDCD
