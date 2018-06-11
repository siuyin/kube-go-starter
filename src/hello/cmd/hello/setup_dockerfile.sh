#!/bin/sh
sed -e "s/_BIN_/${BIN}/g" Dockerfile.orig > Dockerfile
