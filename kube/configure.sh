#!/bin/bash
CFG_PROG='../vendor/github.com/siuyin/cfg-apply/cfg-apply'

if ! [[ -x $CFG_PROG ]]; then
        echo "please make ${CFG_PROG}"
        echo "  cd `dirname ${CFG_PROG}`"
        echo "  go build"
        exit
fi

$CFG_PROG local.config.yaml vols.template.yaml > vols.local.yaml

$CFG_PROG local.config.yaml infra.template.yaml > infra.local.yaml

#$CFG_PROG local.config.yaml deploy.template.yaml > deploy.local.yaml
