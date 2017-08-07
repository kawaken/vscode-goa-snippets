#!/bin/bash

: ${GITHUB_TOKEN?}
: ${GITHUB_USER?}

github-release info -r vscode-goa-snippets
