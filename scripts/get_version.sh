#!/bin/bash

cat $(dirname $(realpath "$0"))/../agb/version.txt | head -n 1
