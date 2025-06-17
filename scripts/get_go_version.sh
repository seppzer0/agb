#!/bin/bash

echo v$(cat $(dirname $(realpath "$0"))/../agb/go.mod | head -n 1)
